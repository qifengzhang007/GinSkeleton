package yml_config

import (
	"github.com/spf13/viper"
	"goskeleton/app/core/container"
	"goskeleton/app/global/my_errors"
	"goskeleton/app/global/variable"
	"log"
	"time"
)

// 创建一个yaml配置文件工厂
func CreateYamlFactory() *ymlConfig {

	yamlConfig := viper.New()
	yamlConfig.AddConfigPath(variable.BasePath + "/config")
	// 需要读取的文件名
	yamlConfig.SetConfigName("config")
	//设置配置文件类型
	yamlConfig.SetConfigType("yml")

	if err := yamlConfig.ReadInConfig(); err != nil {
		log.Fatal(my_errors.ErrorsConfigInitFail + err.Error())
	}

	return &ymlConfig{
		yamlConfig,
	}
}

type ymlConfig struct {
	viper *viper.Viper
}

// 判断相关键是否已经缓存
func (c *ymlConfig) keyIsCache(keyName string) bool {
	if _, exists := container.CreateContainersFactory().KeyIsExists(variable.ConfigKeyPrefix + keyName); exists {
		return true
	} else {
		return false
	}
}

// 对键值进行缓存
func (c *ymlConfig) cache(keyName string, value interface{}) bool {
	return container.CreateContainersFactory().Set(variable.ConfigKeyPrefix+keyName, value)
}

// 通过键获取缓存的值
func (c *ymlConfig) getValueFromCache(keyName string) interface{} {
	return container.CreateContainersFactory().Get(variable.ConfigKeyPrefix + keyName)
}

// Get 一个原始值
func (c *ymlConfig) Get(keyName string) interface{} {
	if c.keyIsCache(keyName) {
		return c.getValueFromCache(keyName)
	} else {
		value := c.viper.Get(keyName)
		c.cache(keyName, value)
		return value
	}
}

// GetString
func (c *ymlConfig) GetString(keyName string) string {
	if c.keyIsCache(keyName) {
		return c.getValueFromCache(keyName).(string)
	} else {
		value := c.viper.GetString(keyName)
		c.cache(keyName, value)
		return value
	}

}

// GetBool
func (c *ymlConfig) GetBool(keyName string) bool {
	if c.keyIsCache(keyName) {
		return c.getValueFromCache(keyName).(bool)
	} else {
		value := c.viper.GetBool(keyName)
		c.cache(keyName, value)
		return value
	}
}

// GetInt
func (c *ymlConfig) GetInt(keyName string) int {
	if c.keyIsCache(keyName) {
		return c.getValueFromCache(keyName).(int)
	} else {
		value := c.viper.GetInt(keyName)
		c.cache(keyName, value)
		return value
	}
}

// GetInt32
func (c *ymlConfig) GetInt32(keyName string) int32 {
	if c.keyIsCache(keyName) {
		return c.getValueFromCache(keyName).(int32)
	} else {
		value := c.viper.GetInt32(keyName)
		c.cache(keyName, value)
		return value
	}
}

// GetInt64
func (c *ymlConfig) GetInt64(keyName string) int64 {
	if c.keyIsCache(keyName) {
		return c.getValueFromCache(keyName).(int64)
	} else {
		value := c.viper.GetInt64(keyName)
		c.cache(keyName, value)
		return value
	}
}

// float64
func (c *ymlConfig) GetFloat64(keyName string) float64 {
	if c.keyIsCache(keyName) {
		return c.getValueFromCache(keyName).(float64)
	} else {
		value := c.viper.GetFloat64(keyName)
		c.cache(keyName, value)
		return value
	}
}

// GetDuration
func (c *ymlConfig) GetDuration(keyName string) time.Duration {
	if c.keyIsCache(keyName) {
		return c.getValueFromCache(keyName).(time.Duration)
	} else {
		value := c.viper.GetDuration(keyName)
		c.cache(keyName, value)
		return value
	}
}

// GetStringSlice
func (c *ymlConfig) GetStringSlice(keyName string) []string {
	if c.keyIsCache(keyName) {
		return c.getValueFromCache(keyName).([]string)
	} else {
		value := c.viper.GetStringSlice(keyName)
		c.cache(keyName, value)
		return value
	}
}
