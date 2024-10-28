package common

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/chuckpreslar/inflect"
	"github.com/serenize/snaker"
)

type TabField struct {
	Typ          string
	Name         string
	Comment      string
	Length       int64
	DefaultValue string
}

type Tab struct {
	TableName string
	Name      string
	Comment   string
	Fields    []TabField
}

type SchemaApi struct {
	Syntax      string
	ServiceName string
	Tables      []Tab
	Enums       EnumCollection
	pconf       *ProtoConfig
}

func (conf ProtoConfig) GenerateApi() (*SchemaApi, error) {

	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", conf.User, conf.Password, conf.Host, conf.Port, conf.Schema)
	db, err := sql.Open(conf.DbType, connStr)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	s := &SchemaApi{
		pconf: &conf,
	}

	dbs, err := dbSchema(db)
	if nil != err {
		return nil, err
	}

	s.Syntax = "v1"
	s.ServiceName = conf.ServiceName
	table := strings.Join(conf.Tables, ",")
	cols, err := dbColumns(db, dbs, table)
	if nil != err {
		return nil, err
	}

	typMap := map[string]*Tab{}
	ignoreMap := map[string]bool{}
	for _, ig := range conf.IgnoreTables {
		ignoreMap[ig] = true
	}

	for _, col := range cols {
		if _, ok := ignoreMap[col.TableName]; ok {
			continue
		}

		name := snaker.SnakeToCamel(col.TableName)
		name = inflect.Singularize(name)

		msg, ok := typMap[name]
		if !ok {
			typMap[name] = &Tab{Name: name, Comment: col.TableComment, TableName: col.TableName}
			msg = typMap[name]
		}

		err := parseApiColumn(s, msg, col)
		if nil != err {
			return nil, err
		}

	}

	for _, t := range typMap {
		s.Tables = append(s.Tables, *t)
	}

	return s, nil
}

func parseApiColumn(s *SchemaApi, msg *Tab, col Column) error {
	typ := strings.ToLower(col.DataType)
	var fieldType string

	switch typ {
	case "char", "varchar", "text", "longtext", "mediumtext", "tinytext":
		fieldType = "string"
	case "enum", "set":
		enumPage := regexp.MustCompile(`[enum|set]\((.+?)\)`).FindStringSubmatch(col.ColumnType)
		enums := strings.FieldsFunc(enumPage[1], func(c rune) bool {
			cs := string(c)
			return cs == "," || cs == "'"
		})

		enumName := inflect.Singularize(snaker.SnakeToCamel(col.TableName)) + snaker.SnakeToCamel(col.ColumnName)
		enum, err := newEnumFromStrings(enumName, col.ColumnComment, enums)
		if nil != err {
			return err
		}

		s.Enums = append(s.Enums, enum)

		fieldType = enumName
	case "blob", "mediumblob", "longblob", "varbinary", "binary":
		fieldType = "bytes"
	case "date", "time", "datetime", "timestamp":

		fieldType = "string"
	case "bool":
		fieldType = "bool"
	case "tinyint", "smallint", "int", "mediumint", "bigint":
		fieldType = "int64"
	case "float", "decimal", "double":
		fieldType = "float64"
	case "json":
		fieldType = "string"
	}

	if fieldType == "" {
		return fmt.Errorf("no compatible protobuf type found for `%s`. column: `%s`.`%s`", col.DataType, col.TableName, col.ColumnName)
	}

	field := TabField{
		Typ:          fieldType,
		Name:         col.ColumnName,
		Comment:      col.ColumnComment,
		Length:       col.CharacterMaximumLength.Int64,
		DefaultValue: col.DefaultValue.String,
	}

	msg.Fields = append(msg.Fields, field)

	return nil
}

func (s *SchemaApi) String() string {
	buf := new(bytes.Buffer)
	buf.WriteString(fmt.Sprintf("syntax = \"%s\"\n", s.Syntax))
	buf.WriteString("\n")
	buf.WriteString("info( \n")
	buf.WriteString("\t title: \" " + s.ServiceName + " \"\n")
	buf.WriteString("\t desc: \"API文件\"\n")
	buf.WriteString("\t author: \"kyle\"\n")
	buf.WriteString("\t email: \"kyle@imguo.com\"\n")
	buf.WriteString("\t version: \"v1\"\n")
	buf.WriteString(")\n\n")

	buf.WriteString("// ------------------------------------ \n")
	buf.WriteString("// Types\n")
	buf.WriteString("// ------------------------------------ \n\n")

	buf.WriteString("type (\n")
	buf.WriteString("   IDReq {\n")
	buf.WriteString("      ID int64  `path:\"id\"`\n")
	buf.WriteString("   }\n\n")
	for _, tab := range s.Tables {
		buf.WriteString("   //--------------------------------" + tab.Comment + "--------------------------------")
		buf.WriteString("\n")
		tab.genDefault(buf, s)
		buf.WriteString("\n")
		tab.genGetAll(buf, s)
		tab.genGetInfo(buf)
		buf.WriteString("\n")
		tab.genAdd(buf, s)
		buf.WriteString("\n")
		tab.genUpdate(buf, s)
		buf.WriteString("\n")

	}
	buf.WriteString(")\n\n\n")

	buf.WriteString("// ------------------------------------ \n")
	buf.WriteString("// Services\n")
	buf.WriteString("// ------------------------------------ \n\n")
	temp := `
@server(
	prefix: v1/%s
	group: %s
	%s
	%s
)`
	jwt := ""
	if s.pconf.Jwt {
		jwt = "jwt:Auth"
	}
	middleware := ""
	if len(s.pconf.Middleware) > 0 {
		middleware = "middleware:" + strings.Join(s.pconf.Middleware, ",")
	}

	for _, tab := range s.Tables {
		buf.WriteString("//--------------------------------" + tab.Comment + "--------------------------------")
		buf.WriteString(fmt.Sprintf(temp, s.ServiceName, tab.TableName, jwt, middleware))
		buf.WriteString("\n")
		buf.WriteString("service " + s.ServiceName + "-api{")
		buf.WriteString("\n")

		buf.WriteString("   @handler Page \n")
		buf.WriteString(fmt.Sprintf("   get /%s (%sPageReq) returns(%sPageResp)\n\n", tab.TableName, tab.Name, tab.Name))

		buf.WriteString("   @handler Get \n")
		buf.WriteString(fmt.Sprintf("   get /%s/:id (IDReq) returns(%sInfoResp)\n\n", tab.TableName, tab.Name))

		buf.WriteString("   @handler Add \n")
		buf.WriteString(fmt.Sprintf("   post /%s (%sAddReq) \n\n", tab.TableName, tab.Name))

		buf.WriteString("   @handler Update \n")
		buf.WriteString(fmt.Sprintf("   put /%s/:id (%sUpdateReq) \n\n", tab.TableName, tab.Name))

		buf.WriteString("   @handler Del \n")
		buf.WriteString(fmt.Sprintf("   delete /%s/:id (IDReq) \n\n", tab.TableName))
		buf.WriteString("}\n\n")
	}

	return buf.String()
}

func (tab Tab) genDefault(buf *bytes.Buffer, s *SchemaApi) {
	buf.WriteString("   " + tab.Name + " {\n")
	for _, field := range tab.Fields {
		if isInSlice(s.pconf.IgnoreColumns, field.Name) {
			continue
		}
		name := From(field.Name).ToCamel()
		comment := ""
		tag := fmt.Sprintf("`json:\"%s\"`", field.Name)
		if field.Comment != "" {
			comment = "// " + field.Comment
		}
		buf.WriteString(fmt.Sprintf("      %s  %s %s  %s \n", name, field.Typ, tag, comment))
	}

	buf.WriteString("   }\n\n")
}

func (tab Tab) genGetAll(buf *bytes.Buffer, s *SchemaApi) {
	buf.WriteString("" + tab.Name + "PageReq {\n")
	for _, field := range tab.Fields {

		if isInSlice(s.pconf.IgnoreColumns, field.Name) {
			continue
		}

		name := From(field.Name).ToCamel()
		comment := ""
		tag := fmt.Sprintf("`form:\"%s,optional,omitempty\"`", field.Name)
		if field.Comment != "" {
			comment = "// " + field.Comment
		}
		buf.WriteString(fmt.Sprintf("      %s  %s %s  %s \n", name, field.Typ, tag, comment))

	}
	buf.WriteString("      Page  uint `form:\"page,optional,default=1\"`\n")
	buf.WriteString("      Limit uint `form:\"limit,optional,default=10\"`\n")
	buf.WriteString("   }\n\n")
	buf.WriteString("   " + tab.Name + "PageResp {\n")
	buf.WriteString(fmt.Sprintf("      Items   []%s `form:\"items\"`\n", tab.Name))
	buf.WriteString("      Total uint  `form:\"total\"`\n")
	buf.WriteString("   }\n\n")
}

func (tab Tab) genGetInfo(buf *bytes.Buffer) {

	buf.WriteString("   " + tab.Name + "InfoResp {\n")
	buf.WriteString(fmt.Sprintf("      %s \n", tab.Name))

	buf.WriteString("   }\n\n")
}

func (tab Tab) genAdd(buf *bytes.Buffer, s *SchemaApi) {
	buf.WriteString("   " + tab.Name + "AddReq {\n")
	for _, field := range tab.Fields {

		if isInSlice(s.pconf.IgnoreColumns, field.Name) || isInSlice(s.pconf.OnlySearch, field.Name) {
			continue
		}

		name := From(field.Name).ToCamel()
		comment := ""
		validate := ""
		if field.Typ == "string" {
			if field.Length > 0 {
				validate = fmt.Sprintf(" validate:\"min=2,max=%v\"", field.Length)
			} else {
				validate = " validate:\"min=2\""
			}
		}
		fm := field.Name
		if len(field.DefaultValue) > 0 {
			fm = fm + ",default=" + field.DefaultValue
		}
		tag := fmt.Sprintf("`json:\"%s\"%s`", fm, validate)
		if field.Comment != "" {
			comment = "// " + field.Comment
		}
		buf.WriteString(fmt.Sprintf("      %s  %s %s  %s \n", name, field.Typ, tag, comment))

	}

	buf.WriteString("   }\n\n")

}

func (tab Tab) genUpdate(buf *bytes.Buffer, s *SchemaApi) {
	buf.WriteString("" + tab.Name + "UpdateReq {\n")
	buf.WriteString("      ID int64  `path:\"id\"`\n")

	for _, field := range tab.Fields {
		if isInSlice(s.pconf.IgnoreColumns, field.Name) || isInSlice(s.pconf.OnlySearch, field.Name) {
			continue
		}

		name := From(field.Name).ToCamel()
		comment := ""
		validate := ""
		if field.Typ == "string" {
			if field.Length > 0 {
				validate = fmt.Sprintf(" validate:\"min=2,max=%v\"", field.Length)
			} else {
				validate = " validate:\"min=2\""
			}
		}
		fm := field.Name
		if len(field.DefaultValue) > 0 {
			fm = fm + ",default=" + field.DefaultValue
		}
		tag := fmt.Sprintf("`json:\"%s\"%s`", fm, validate)
		if field.Comment != "" {
			comment = "// " + field.Comment
		}
		buf.WriteString(fmt.Sprintf("      %s  %s %s  %s \n", name, field.Typ, tag, comment))

	}

	buf.WriteString("   }\n\n")

}
