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
	yamlConfig.SetConfigType("yaml")

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

// get 一个原始值
func (c *ymlConfig) Get(keyname string) interface{} {
	return c.viper.Get(keyname)
}

// getstring
func (c *ymlConfig) GetString(keyname string) string {
	return c.viper.GetString(keyname)
}

// getbool
func (c *ymlConfig) GetBool(keyname string) bool {
	return c.viper.GetBool(keyname)
}

// getint
func (c *ymlConfig) GetInt(keyname string) int {
	return c.viper.GetInt(keyname)
}

// getint32
func (c *ymlConfig) GetInt32(keyname string) int32 {
	return c.viper.GetInt32(keyname)
}

// getint64
func (c *ymlConfig) GetInt64(keyname string) int64 {
	return c.viper.GetInt64(keyname)
}

// float64
func (c *ymlConfig) GetFloat64(keyname string) float64 {
	return c.viper.GetFloat64(keyname)
}

// GetDuration
func (c *ymlConfig) GetDuration(keyname string) time.Duration {
	return c.viper.GetDuration(keyname)
}

// GetStringSlice
func (c *ymlConfig) GetStringSlice(keyname string) []string {
	return c.viper.GetStringSlice(keyname)
}
