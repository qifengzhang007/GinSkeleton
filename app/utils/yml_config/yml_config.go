package yml_config

import (
	"github.com/spf13/viper"
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

// Get 一个原始值
func (c *ymlConfig) Get(keyName string) interface{} {
	return c.viper.Get(keyName)
}

// GetString
func (c *ymlConfig) GetString(keyName string) string {
	return c.viper.GetString(keyName)
}

// GetBool
func (c *ymlConfig) GetBool(keyName string) bool {
	return c.viper.GetBool(keyName)
}

// GetInt
func (c *ymlConfig) GetInt(keyName string) int {
	return c.viper.GetInt(keyName)
}

// GetInt32
func (c *ymlConfig) GetInt32(keyName string) int32 {
	return c.viper.GetInt32(keyName)
}

// GetInt64
func (c *ymlConfig) GetInt64(keyName string) int64 {
	return c.viper.GetInt64(keyName)
}

// float64
func (c *ymlConfig) GetFloat64(keyName string) float64 {
	return c.viper.GetFloat64(keyName)
}

// GetDuration
func (c *ymlConfig) GetDuration(keyName string) time.Duration {
	return c.viper.GetDuration(keyName)
}

// GetStringSlice
func (c *ymlConfig) GetStringSlice(keyName string) []string {
	return c.viper.GetStringSlice(keyName)
}
