package Api

import (
	"ZgoMast/application/Constants"
	"ZgoMast/application/Model"
	. "ZgoMast/application/Service"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
添加主机类别
*/

func HostCategoryCreate(ctx *gin.Context) {
	hostCategory, err := CreateHostCategory(ctx)
	if err != nil || hostCategory.ID < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    Constants.CodeCreateHostCategoryFail,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    Constants.CodeSuccess,
		"message": Constants.Success,
		"data": map[string]Model.HostCategoryInstance{
			"hostCategory": hostCategory,
		},
	})
}
