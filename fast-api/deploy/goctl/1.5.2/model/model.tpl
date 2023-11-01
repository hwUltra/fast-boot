package {{.pkg}}
{{if .withCache}}
import (
	"context"
	{{if .time}}"time"{{end}}

	"fast-boot/common/xerr"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/Masterminds/squirrel"
)
{{else}}
import (
	"context"
	{{if .time}}"time"{{end}}
	"fast-boot/common/xerr"
		"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/Masterminds/squirrel"
	
)

{{end}}
var _ {{.upperStartCamelObject}}Model = (*custom{{.upperStartCamelObject}}Model)(nil)

type (
	// {{.upperStartCamelObject}}Model is an interface to be customized, add more methods here,
	// and implement the added methods in custom{{.upperStartCamelObject}}Model.
	{{.upperStartCamelObject}}Model interface {
		{{.lowerStartCamelObject}}Model
		Trans(ctx context.Context,fn func(context context.Context,session sqlx.Session) error) error
		RowBuilder() squirrel.SelectBuilder
		CountBuilder(field string) squirrel.SelectBuilder
		SumBuilder(field string) squirrel.SelectBuilder
		DeleteSoft(ctx context.Context,session sqlx.Session,{{if .containsIndexCache}}newData{{else}}data{{end}} *{{.upperStartCamelObject}}) error
}

	custom{{.upperStartCamelObject}}Model struct {
		*default{{.upperStartCamelObject}}Model
	}
)

// New{{.upperStartCamelObject}}Model returns a model for the database table.
func New{{.upperStartCamelObject}}Model(conn sqlx.SqlConn{{if .withCache}}, c cache.CacheConf, opts ...cache.Option{{end}}) {{.upperStartCamelObject}}Model {
	return &custom{{.upperStartCamelObject}}Model{
		default{{.upperStartCamelObject}}Model: new{{.upperStartCamelObject}}Model(conn{{if .withCache}}, c, opts...{{end}}),
	}
}

// export logic
func (m *default{{.upperStartCamelObject}}Model) Trans(ctx context.Context,fn func(ctx context.Context,session sqlx.Session) error) error {
	{{if .withCache}}
	return m.TransactCtx(ctx,func(ctx context.Context,session sqlx.Session) error {
		return  fn(ctx,session)
	})
	{{else}}
	return m.conn.TransactCtx(ctx,func(ctx context.Context,session sqlx.Session) error {
		return  fn(ctx,session)
	})
	{{end}}
}

func (m *default{{.upperStartCamelObject}}Model) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select({{.lowerStartCamelObject}}Rows).From(m.table)
}

func (m *default{{.upperStartCamelObject}}Model) CountBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("COUNT("+field+")").From(m.table)
}

func (m *default{{.upperStartCamelObject}}Model) SumBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("IFNULL(SUM("+field+"),0)").From(m.table)
}

func (m *default{{.upperStartCamelObject}}Model) DeleteSoft(ctx context.Context,session sqlx.Session,{{if .containsIndexCache}}newData{{else}}data{{end}} *{{.upperStartCamelObject}}) error {
	
    {{if .containsTime}}
		data.DeleteTime = time.Now()
		if err:= m.Update(ctx,session, data);err!= nil{
			return errors.Wrapf(xerr.NewErrMsg("删除数据失败"),"{{.upperStartCamelObject}}Model delete err : %+v",err)
		}
   {{else}}
		if err:= m.Delete(ctx, data.Id);err!= nil{
			return errors.Wrapf(xerr.NewErrMsg("删除数据失败"),"{{.upperStartCamelObject}}Model delete err : %+v",err)
		}
	{{end}}
	return nil

}
