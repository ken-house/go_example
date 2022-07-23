package service

import (
	"io"

	"github.com/spf13/cast"

	"github.com/gin-gonic/gin"
	"github.com/ken-house/go-contrib/prototype/excelHandler"

	MysqlModel "github.com/go_example/internal/model/mysql"
)

type ExcelUserService interface {
	ExportUser(*gin.Context, []MysqlModel.User) error
	ImportUser(*gin.Context, io.Reader) ([]MysqlModel.User, error)
}

type excelUserService struct {
	excelExportHandler excelHandler.ExcelExportHandler
	excelImportHandler excelHandler.ExcelImportHandler
}

func NewExcelUserService(
	excelExportHandler excelHandler.ExcelExportHandler,
	excelImportHandler excelHandler.ExcelImportHandler,
) ExcelUserService {
	return &excelUserService{
		excelExportHandler: excelExportHandler,
		excelImportHandler: excelImportHandler,
	}
}

// ImportUser 导入文件
func (svc *excelUserService) ImportUser(c *gin.Context, file io.Reader) ([]MysqlModel.User, error) {
	userList := make([]MysqlModel.User, 0, 100)

	// 表头（用于检测文件内容是否符合要求）
	headerArr := []string{"用户Id", "用户名", "密码", "性别"}

	// 读取文件数据
	rows, err := svc.excelImportHandler.Import(c, file, 0, headerArr)
	if err != nil {
		return userList, err
	}

	// 格式化数据
	for i, row := range rows {
		if i == 0 {
			continue
		}
		user := MysqlModel.User{
			Id:       cast.ToInt(row[0]),
			Username: row[1],
			Password: row[2],
			Gender:   cast.ToInt(row[3]),
		}
		userList = append(userList, user)
	}
	return userList, err
}

// ExportUser 导出用户信息
func (svc *excelUserService) ExportUser(c *gin.Context, userList []MysqlModel.User) error {
	// 文件名
	fileName := "hello.xlsx"
	// 表头
	headerArr := []string{"用户Id", "用户名", "密码", "性别"}

	// 表格数据
	rows := make([][]interface{}, 0, 100)
	for _, user := range userList {
		row := []interface{}{user.Id, user.Username, user.Password, user.Gender}
		rows = append(rows, row)
	}

	// 导出文件
	return svc.excelExportHandler.Export(c, rows, headerArr, fileName)
}
