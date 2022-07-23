package controller

import (
	"net/http"

	"github.com/go_example/internal/lib/errorAssets"

	"github.com/go_example/internal/service"

	"github.com/go_example/internal/utils/negotiate"

	"github.com/gin-gonic/gin"
)

type ExcelController interface {
	Export(*gin.Context) (int, gin.Negotiate)
	Import(*gin.Context) (int, gin.Negotiate)
}

type excelController struct {
	userSvc  service.UserService
	excelSvc service.ExcelUserService
}

func NewExcelController(
	userSvc service.UserService,
	excelSvc service.ExcelUserService,
) ExcelController {
	return &excelController{
		userSvc:  userSvc,
		excelSvc: excelSvc,
	}
}

func (ctr *excelController) Export(c *gin.Context) (int, gin.Negotiate) {
	userList := ctr.userSvc.GetUserList(c)
	err := ctr.excelSvc.ExportUser(c, userList)
	if err != nil {
		return negotiate.JSON(http.StatusOK, errorAssets.ERR_EXPORT.ToastError())
	}
	return negotiate.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"message": "导出成功",
		},
	})
}

func (ctr *excelController) Import(c *gin.Context) (int, gin.Negotiate) {
	file, _, err := c.Request.FormFile("uploadFile")
	if err != nil {
		return negotiate.JSON(http.StatusOK, errorAssets.ERR_IMPORT.ToastError())
	}
	defer file.Close()
	userList, err := ctr.excelSvc.ImportUser(c, file)
	if err != nil {
		return negotiate.JSON(http.StatusOK, errorAssets.ERR_FILE_PARSE.ToastError())
	}

	// 写入到数据库
	err = ctr.userSvc.InsertUserList(userList)
	if err != nil {
		return negotiate.JSON(http.StatusOK, errorAssets.ERR_SYSTEM.ToastError())
	}

	return negotiate.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"message": "导入成功",
		},
	})
}
