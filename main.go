package main

import (
	"awesomeProject/dao/mysql"
	"awesomeProject/dao/redis"
	"awesomeProject/logger1"
	"awesomeProject/pkg/snowflake"
	"awesomeProject/routes"
	"awesomeProject/settings"
	"fmt"

	"go.uber.org/zap"
)

// go web 通用脚手架
func main() {
	//1,加载配置文件
	if err := settings.Init(); err != nil {
		fmt.Printf("init settings failed, err:%v\n", err)
		return
	}

	//2，加载日志
	if err := logger1.Init(settings.Conf.LogConfig); err != nil {
		fmt.Printf("init settings failed, err:%v\n", err)
		return
	}
	defer zap.L().Sync()
	zap.L().Debug("logger init success...")

	//3，初始化mysql连接
	if err := mysql.Init(settings.Conf.MysqlConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}

	//4，初始化Redis连接
	if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis  failed, err:%v\n", err)
		return
	}
	defer redis.Close()
	snowflake.Init(settings.Conf.StartTime, settings.Conf.MachineID)
	//id := snowflake.GenID()
	//fmt.Println(id)

	//5，注册路由
	r := routes.Setup()

	//6，启动服务
	r.Run()

}
