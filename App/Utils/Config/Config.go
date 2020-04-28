package Config

import (
	"GinSkeleton/App/Global/MyErrors"
	"GinSkeleton/App/Global/Variable"
	"github.com/spf13/viper"
	"log"
	"time"
)

// 创建一个yaml配置文件工厂
func CreateYamlFactory() *ConfigYml {

	yaml_config := viper.New()
	yaml_config.AddConfigPath(Variable.BASE_PATH + "/Config")
	// 需要读取的文件名
	yaml_config.SetConfigName("config")
	//设置配置文件类型
	yaml_config.SetConfigType("yaml")

	if err := yaml_config.ReadInConfig(); err != nil {
		log.Fatal(MyErrors.Errors_Config_Init_Fail + err.Error())
	}

	return &ConfigYml{
		yaml_config,
	}
}

type ConfigYml struct {
	viper *viper.Viper
}

// get 一个原始值
func (c *ConfigYml) Get(keyname string) interface{} {
	return c.viper.Get(keyname)
}

// getstring
func (c *ConfigYml) GetString(keyname string) string {
	return c.viper.GetString(keyname)
}

// getbool
func (c *ConfigYml) GetBool(keyname string) bool {
	return c.viper.GetBool(keyname)
}

// getint
func (c *ConfigYml) GetInt(keyname string) int {
	return c.viper.GetInt(keyname)
}

// getint32
func (c *ConfigYml) GetInt32(keyname string) int32 {
	return c.viper.GetInt32(keyname)
}

// getint64
func (c *ConfigYml) GetInt64(keyname string) int64 {
	return c.viper.GetInt64(keyname)
}

// float64
func (c *ConfigYml) GetFloat64(keyname string) float64 {
	return c.viper.GetFloat64(keyname)
}

// GetDuration
func (c *ConfigYml) GetDuration(keyname string) time.Duration {
	return c.viper.GetDuration(keyname)
}

// GetStringSlice
func (c *ConfigYml) GetStringSlice(keyname string) []string {
	return c.viper.GetStringSlice(keyname)
}
