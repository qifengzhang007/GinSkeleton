package yml_config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"goskeleton/app/core/container"
	"goskeleton/app/global/my_errors"
	"goskeleton/app/global/variable"
	"goskeleton/app/utils/yml_config/ymlconfig_interf"
	"log"
	"sync"
	"time"
)

// 由于 vipver 包本身对于文件的变化事件有一个bug，相关事件会被回调两次
// 常年未彻底解决，相关的 issue 清单：https://github.com/spf13/viper/issues?q=OnConfigChange
// 设置一个内部全局变量，记录配置文件变化时的时间点，如果两次回调事件事件差小于1秒，我们认为是第二次回调事件，而不是人工修改配置文件
// 这样就避免了 viper 包的这个bug

var lastChangeTime time.Time
var containerFactory = container.CreateContainersFactory()

func init() {
	lastChangeTime = time.Now()
}

// CreateYamlFactory 创建一个yaml配置文件工厂
// 参数设置为可变参数的文件名，这样参数就可以不需要传递，如果传递了多个，我们只取第一个参数作为配置文件名
func CreateYamlFactory(fileName ...string) ymlconfig_interf.YmlConfigInterf {

	yamlConfig := viper.New()
	// 配置文件所在目录
	yamlConfig.AddConfigPath(variable.BasePath + "/config")
	// 需要读取的文件名,默认为：config
	if len(fileName) == 0 {
		yamlConfig.SetConfigName("config")
	} else {
		yamlConfig.SetConfigName(fileName[0])
	}
	//设置配置文件类型(后缀)为 yml
	yamlConfig.SetConfigType("yml")

	if err := yamlConfig.ReadInConfig(); err != nil {
		log.Fatal(my_errors.ErrorsConfigInitFail + err.Error())
	}

	return &ymlConfig{
		viper: yamlConfig,
		mu:    new(sync.Mutex),
	}
}

type ymlConfig struct {
	viper *viper.Viper
	mu    *sync.Mutex
}

// ConfigFileChangeListen 监听文件变化
func (y *ymlConfig) ConfigFileChangeListen() {
	y.viper.OnConfigChange(func(changeEvent fsnotify.Event) {
		if time.Now().Sub(lastChangeTime).Seconds() >= 1 {
			if changeEvent.Op.String() == "WRITE" {
				y.clearCache()
				lastChangeTime = time.Now()
			}
		}
	})
	y.viper.WatchConfig()
}

// keyIsCache 判断相关键是否已经缓存
func (y *ymlConfig) keyIsCache(keyName string) bool {
	if _, exists := containerFactory.KeyIsExists(variable.ConfigKeyPrefix + keyName); exists {
		return true
	} else {
		return false
	}
}

// 对键值进行缓存
func (y *ymlConfig) cache(keyName string, value interface{}) bool {
	// 避免瞬间缓存键、值时，程序提示键名已经被注册的日志输出
	y.mu.Lock()
	defer y.mu.Unlock()
	if _, exists := containerFactory.KeyIsExists(variable.ConfigKeyPrefix + keyName); exists {
		return true
	}
	return containerFactory.Set(variable.ConfigKeyPrefix+keyName, value)
}

// 通过键获取缓存的值
func (y *ymlConfig) getValueFromCache(keyName string) interface{} {
	return containerFactory.Get(variable.ConfigKeyPrefix + keyName)
}

// 清空已经缓存的配置项信息
func (y *ymlConfig) clearCache() {
	containerFactory.FuzzyDelete(variable.ConfigKeyPrefix)
}

// Clone 允许 clone 一个相同功能的结构体
func (y *ymlConfig) Clone(fileName string) ymlconfig_interf.YmlConfigInterf {
	// 这里存在一个深拷贝，需要注意，避免拷贝的结构体操作对原始结构体造成影响
	var ymlC = *y
	var ymlConfViper = *(y.viper)
	(&ymlC).viper = &ymlConfViper

	(&ymlC).viper.SetConfigName(fileName)
	if err := (&ymlC).viper.ReadInConfig(); err != nil {
		// 程序启动阶段，zaplog 未初始化，使用系统log打印启动时候发生的异常日志
		if variable.ZapLog == nil {
			log.Fatal(my_errors.ErrorsConfigInitFail+", error: ", err.Error())
		} else {
			variable.ZapLog.Error(my_errors.ErrorsConfigInitFail, zap.Error(err))
		}
	}
	return &ymlC
}

// Get 一个原始值
func (y *ymlConfig) Get(keyName string) interface{} {
	if y.keyIsCache(keyName) {
		return y.getValueFromCache(keyName)
	} else {
		value := y.viper.Get(keyName)
		y.cache(keyName, value)
		return value
	}
}

// GetString 字符串格式返回值
func (y *ymlConfig) GetString(keyName string) string {
	if y.keyIsCache(keyName) {
		return y.getValueFromCache(keyName).(string)
	} else {
		value := y.viper.GetString(keyName)
		y.cache(keyName, value)
		return value
	}

}

// GetBool 布尔格式返回值
func (y *ymlConfig) GetBool(keyName string) bool {
	if y.keyIsCache(keyName) {
		return y.getValueFromCache(keyName).(bool)
	} else {
		value := y.viper.GetBool(keyName)
		y.cache(keyName, value)
		return value
	}
}

// GetInt 整数格式返回值
func (y *ymlConfig) GetInt(keyName string) int {
	if y.keyIsCache(keyName) {
		return y.getValueFromCache(keyName).(int)
	} else {
		value := y.viper.GetInt(keyName)
		y.cache(keyName, value)
		return value
	}
}

// GetInt32 整数格式返回值
func (y *ymlConfig) GetInt32(keyName string) int32 {
	if y.keyIsCache(keyName) {
		return y.getValueFromCache(keyName).(int32)
	} else {
		value := y.viper.GetInt32(keyName)
		y.cache(keyName, value)
		return value
	}
}

// GetInt64 整数格式返回值
func (y *ymlConfig) GetInt64(keyName string) int64 {
	if y.keyIsCache(keyName) {
		return y.getValueFromCache(keyName).(int64)
	} else {
		value := y.viper.GetInt64(keyName)
		y.cache(keyName, value)
		return value
	}
}

// GetFloat64 小数格式返回值
func (y *ymlConfig) GetFloat64(keyName string) float64 {
	if y.keyIsCache(keyName) {
		return y.getValueFromCache(keyName).(float64)
	} else {
		value := y.viper.GetFloat64(keyName)
		y.cache(keyName, value)
		return value
	}
}

// GetDuration 时间单位格式返回值
func (y *ymlConfig) GetDuration(keyName string) time.Duration {
	if y.keyIsCache(keyName) {
		return y.getValueFromCache(keyName).(time.Duration)
	} else {
		value := y.viper.GetDuration(keyName)
		y.cache(keyName, value)
		return value
	}
}

// GetStringSlice 字符串切片数格式返回值
func (y *ymlConfig) GetStringSlice(keyName string) []string {
	if y.keyIsCache(keyName) {
		return y.getValueFromCache(keyName).([]string)
	} else {
		value := y.viper.GetStringSlice(keyName)
		y.cache(keyName, value)
		return value
	}
}
