package cache

import (
	"fmt"
	"time"
)

var UserInfoKey = KeyConf{
	Key: "user_info:%d",
}

//----------------------------------------------------------------------------------------------

type KeyConf struct {
	Key        string        `json:"key"`
	Expiration time.Duration `json:"expiration"`
}

// GetCacheKey 获取缓存key
func GetCacheKey(keyConf KeyConf, vals ...interface{}) string {
	return fmt.Sprintf(keyConf.Key, vals...)
}
