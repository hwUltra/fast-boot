package globalkey

import (
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"strconv"
	"time"
)

var CopierProtoOptions = copier.Option{
	IgnoreEmpty: true,
	DeepCopy:    true,
	Converters: []copier.TypeConverter{
		{
			SrcType: time.Time{},
			DstType: copier.String,
			Fn: func(src interface{}) (interface{}, error) {
				s, ok := src.(time.Time)
				if !ok {
					return nil, errors.New("src type :time.Time not matching")
				}
				return s.Format(DateTimeFormatStandardTime), nil
			},
		},
		{
			SrcType: copier.String,
			DstType: copier.Int,
			Fn: func(src interface{}) (interface{}, error) {
				s, ok := src.(string)

				if !ok {
					return nil, errors.New("src type not matching")
				}

				return strconv.Atoi(s)
			},
		},
	},
}
