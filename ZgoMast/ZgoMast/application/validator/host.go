package validator

import (
	. "ZgoMast/application/Model"
	. "ZgoMast/application/utils"
	"errors"
	"github.com/go-playground/validator/v10"
)

/**
添加主机类别的验证器
*/

func HostCategoryValidator(hostCategory *HostCategory) error {
	validate, trans := GenValidate()
	err := validate.Struct(hostCategory)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return errors.New(err.Translate(trans))
		}
	}
	return nil
}
