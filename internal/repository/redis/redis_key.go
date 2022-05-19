package redis

import (
	"fmt"
	"time"
)

// AuthTokenKey 用户认证令牌
var AuthTokenKey = KeyConf{
	Key: "auth_token_%d",
}

//----------------------------------------------------------------------------------------------

type KeyConf struct {
	Key        string            `json:"key"`
	Structure  map[string]string `json:"structure"`
	Expiration time.Duration     `json:"expiration"`
}

// GetRedisKey 获取redis的配置，并返回所需要的key的全名
func GetRedisKey(keyConf KeyConf, args ...interface{}) string {
	return fmt.Sprintf(keyConf.Key, args...)
}

// GetRedisFields 获取redis的配置，并返回所需要的存储在redis的结构的真实名称
func GetRedisFields(keyConf KeyConf, fields ...string) []string {
	var result []string
	for i := 0; i < len(fields); i++ {
		result = append(result, keyConf.Structure[fields[i]])
	}
	return result
}
