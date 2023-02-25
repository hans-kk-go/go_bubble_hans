package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"
)

// conf 全局变量，用来保存程序所有配置信息
var Conf = new(AppConfig)

func Init() (err error) {
	viper.SetConfigName("conf") //指定配置文件名称(不需要带后缀)
	viper.SetConfigType("yaml") //指定配置文件类型
	viper.AddConfigPath("./")   //指定查找路径
	//viper.SetConfigFile("C:\\Users\\hans\\GolandProjects\\awesomeProject\\conf.yaml")//指定绝对路径

	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
		return
	}

	//把读取到的配置信息反序列化到Conf变量中
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.unmarshal failed err:%v\n", err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了...")
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper.unmarshal failed err:%v\n", err)
		}
	}) //viper热加载，时刻监听

	return
}

type AppConfig struct {
	Name         string `mapstructure:"name"`
	Mode         string `mapstructure:"mod"`
	Version      string `mapstructure:"version"`
	StartTime    string `mapstructure:"startTime"`
	MachineID    int64  `mapstructure:"machineId"`
	Port         int    `mapstructure:"port"`
	*LogConfig   `mapstructure:"log"`
	*MysqlConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

type LogConfig struct {
	Level       string `mapstructure:"level"`
	Filename    string `mapstructure:"filename"`
	Max_size    int    `mapstructure:"max_size"`
	Max_age     int    `mapstructure:"max_age"`
	Max_backups int    `mapstructure:"max_backups"`
}

type MysqlConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Dbname   string `mapstructure:"dbname"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	Db       int    `mapstructure:"db"`
}
