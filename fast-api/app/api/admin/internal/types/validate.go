package types

import "github.com/go-playground/validator/v10"

var validate = validator.New()

func (r *BrandDelReq) Validate() error {
	return validate.Struct(r)
}
