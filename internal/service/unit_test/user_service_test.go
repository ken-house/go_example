package unit_test

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/go_example/internal/meta"
	MysqlModel "github.com/go_example/internal/model/mysql"
	MysqlRepo "github.com/go_example/internal/repository/mysql"
	"github.com/go_example/internal/service"
	"github.com/ken-house/go-contrib/prototype/mysqlClient"
	"github.com/ken-house/go-contrib/prototype/nacosClient"
	"github.com/ken-house/go-contrib/utils/env"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func init() {
	// 从系统环境变量中读取运行环境
	meta.EnvMode = env.Mode()
	// 获取Nacos配置
	getNacosConfig()

	if env.IsDebugging() && !meta.DebugUseConfigCenter { // 本地调试若不使用配置中心则直接读取common.yaml文件
		viper.SetConfigFile("../../../" + meta.CfgFile + "/" + meta.EnvMode + "/common.yaml")
		if err := viper.ReadInConfig(); err != nil {
			log.Fatalln(err)
		}
		if err := viper.Unmarshal(&meta.GlobalConfig); err != nil {
			log.Fatalln(err)
		}
	} else { // 测试环境、生产环境从配置中心读取
		configCenterClient, clean, err := nacosClient.NewConfigClient(meta.NacosConfig)
		if err != nil {
			log.Fatalln(err)
		}
		defer clean()
		globalConfigStr, err := configCenterClient.GetConfig(vo.ConfigParam{
			DataId: meta.NacosConfig.DataId,
			Group:  meta.NacosConfig.Group,
		})
		if err != nil {
			log.Fatalln(err)
		}
		// 将读取到的配置信息转为全局配置
		setGlobalConfigFromData(globalConfigStr)

		// 监听实现自动感知
		configCenterClient.ListenConfig(vo.ConfigParam{
			DataId: meta.NacosConfig.DataId,
			Group:  meta.NacosConfig.Group,
			OnChange: func(namespace, group, dataId, data string) {
				setGlobalConfigFromData(data)
			},
		})
	}
}

// 从文本读取到全局配置
func setGlobalConfigFromData(data string) {
	// 将读取到的配置信息转为全局配置
	viper.SetConfigType("yaml")
	err := viper.ReadConfig(bytes.NewBuffer([]byte(data)))
	if err != nil {
		log.Fatalln(err)
	}
	if err = viper.Unmarshal(&meta.GlobalConfig); err != nil {
		log.Fatalln(err)
	}
}

// 读取nacos配置
func getNacosConfig() {
	viper.SetConfigFile("../../../" + meta.CfgFile + "/" + meta.EnvMode + "/config_center.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln(err)
	}

	// 从配置中心读取项目配置
	if err := viper.Sub("config_center").Unmarshal(&meta.NacosConfig); err != nil {
		log.Fatalln(err)
	}
}

func NewService() service.UserService {
	eg, _, _ := mysqlClient.NewGroupClient(meta.GlobalConfig.Mysql.Group)
	userMysqlRepo := MysqlRepo.NewUserRepository(eg)
	return service.NewUserService(userMysqlRepo)
}

// 单元测试函数
func TestInsertUserList(t *testing.T) {
	userList := []MysqlModel.User{
		{
			Username: "zhangsan",
			Password: "zhangsan",
			Gender:   1,
		},
		{
			Username: "lisi",
			Password: "lisi",
			Gender:   2,
		},
	}
	userService := NewService()
	err := userService.InsertUserList(userList)
	if err != nil {
		panic(err)
	}

	user, err := userService.GetUserInfoByFormData(&gin.Context{}, "zhangsan", "zhangsan")
	if err != nil {
		panic(err)
	}
	assert.Equal(t, "zhangsan", user.Username)
	assert.Equal(t, "zhangsan", user.Password)
}
