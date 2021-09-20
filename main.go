package main

import (
	"fmt"
	"time"

	"github.com/reneforever/conf"
	"github.com/reneforever/kafka"
	"github.com/reneforever/taillog"
	"gopkg.in/ini.v1"
)

// logAgent入口

var (
	cfg *conf.AppConf
)

func run(topic string) {
	// 1.读取日志
	for {
		// 2.发送到kafka
		select {
		case line := <-taillog.ReadLogChan():
			kafka.SendToKafka(topic, line.Text)
		default:
			time.Sleep(time.Second)
		}
	}
}

func main() {
	// 0. 加载配置文件
	cfg, iniErr := ini.Load("./conf/config.ini")
	if iniErr != nil {
		fmt.Println("ini load failed, err: ", iniErr)
		return
	}
	// cfg.Section("kafka").Key("address")
	// 1. 初始化kafka连接
	kafkaErr := kafka.Init([]string{cfg.Section("kafka").Key("address").String()})
	if kafkaErr != nil {
		fmt.Println("init kafka failed, err: ", kafkaErr)
		return
	}
	// 2. 打开日志文件准备收集
	tailErr := taillog.Init(cfg.Section("taillog").Key("path").String())
	if tailErr != nil {
		fmt.Println("read log file failed, err: ", tailErr)
		return
	}
	run(cfg.Section("kafka").Key("topic").String())
}
