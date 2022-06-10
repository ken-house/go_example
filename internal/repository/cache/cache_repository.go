package cache

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/spf13/viper"
)

// GetCacheKey 获取缓存key
func GetCacheKey(key string, vals ...interface{}) (string, error) {
	cacheKey := viper.GetString("cacheKeys." + key)
	if cacheKey == "" {
		return "", errors.New("key不存在")
	}
	return fmt.Sprintf(cacheKey, vals...), nil
}
