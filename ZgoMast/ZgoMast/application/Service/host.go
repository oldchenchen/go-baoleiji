package Service

import (
	. "ZgoMast/application/Model"
	"ZgoMast/application/validator"
	"github.com/gin-gonic/gin"
)

/**
添加主机类别
*/

func CreateHostCategory(ctx *gin.Context) (HostCategoryInstance, error) {
	hostCategory := HostCategory{}
	var instance HostCategoryInstance
	var err error
	if err = ctx.ShouldBindJSON(&hostCategory); err != nil {
		return instance, err
	}
	if err = validator.HostCategoryValidator(&hostCategory); err != nil {
		return instance, err
	}
	err = hostCategory.Insert()
	instance = HostCategoryInstance{
		ID:   hostCategory.ID,
		Name: hostCategory.Name,
	}
	return instance, err
}
