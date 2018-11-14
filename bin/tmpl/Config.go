//  
//  {{.BoxConf.Name}}Config.go
//  main
//
//  Created by {{.BoxConf.Author}} on {{.Time}}
//
package main

import (
	"{{.MboxPackagePath}}"
	"os"
	."{{.MboxPackagePath}}/boxdef"
)

/*
	以下代码由boxtool自动产生，
	若不清楚了解每条设置参数和含义不建议修改
*/
func init(){
	var cfg mbox.BoxConfig
	
	//工具类别
	cfg.Group                  = "{{.BoxConf.Group}}"
	
	//工具名称
	cfg.Name                   = "{{.BoxConf.Name}}"
	
	//这个工具展示为什么名字
	cfg.ShowName               = "{{.BoxConf.ShowName}}"
	
	//作者
	cfg.Author                 = "{{.BoxConf.Author}}"
	
	//工具的描述
	cfg.Description            = "{{.BoxConf.Description}}"
	
	//http模式时开启的地址和端口
	cfg.SelfHttpServerHost     = "{{.BoxConf.SelfHttpServerHost}}"
	
	//界面操作层路径
	cfg.ShowServerPath         = "{{.BoxConf.ShowServerPath}}"
	
	//nats通信组件server地址列表 ',' 分割
	cfg.NatsServerList         = "" 
	
	//操作界面访问的路径
	cfg.ModeInfo               = "{{.BoxConf.ModeInfo}}"
	
	//nats 接入用户名
	cfg.NatsServerUserName     = ""
	
	//nats 接入密码
	cfg.NatsServerUserPassword = ""
	
	//日志存储文件夹
	cfg.LogDir                 = "{{.BoxConf.LogDir}}"
	
	//工具的输入参数描述 //附加输入参数数组
	cfg.Params = []BoxParam { {
		{{range .BoxConf.Params}}
		Name : "{{.Name}}",
		Label : "{{.Label}}",
		Type : "{{.Type}}",
		Hint : "{{.Hint}}",
		Reg : "{{.Reg}}",
		Value : "{{.Value}}",
		Idx : "{{.Idx}}"},
	    {{end}}
	}
	
	cfg.Version = "1.0.1"
	
	err := mbox.SetBoxConfig(cfg)
	if err != nil {
		mbox.Log("box:",cfg.Name," init error:",err)
		os.Exit(1)
	}
}
