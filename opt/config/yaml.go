package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"yushu/opt/file"
	"yushu/opt/logs"
)

func OpenYaml(path string) {
	// 捕获 panic
	defer func() {
		if r := recover(); r != nil {
			logs.Print("Config recovered in f", r)
		}
	}()

	// 创建配置结构体实例
	conf := NewApp()

	// "./res/conf/config.yaml"
	// 默认配置文件路径baseFilename
	baseFilename := "./config.yaml"

	// 判断如果文件不存在则创建默认配置文件
	if !file.Exists(baseFilename) {
		f, err := file.Open(baseFilename, os.O_APPEND)
		if err != nil {
			logs.Print("Failed to open config file: ", err) // 错误处理
		}
		// 去掉前后空行回车符
		_, err = f.WriteString(DefaultConfig)
		if err != nil {
			logs.Print("Failed to write config file: ", err) // 错误处理
		}

		//  将 YAML 数据解码到 conf 结构体中
		err = yaml.Unmarshal([]byte(DefaultConfig), conf)
		// 错误处理
		if err != nil {
			log.Fatalf("Failed to unmarshal YAML: %v", err)
		}

		// 打印解析后的配置
		//datastore
		//log.Printf("Parsed 2 Config: %+v\n", conf.Db.Mysql[0])
	} else {
		// 打开文件
		f, err := file.Open(baseFilename)

		//fileData, err := os.Open("./resource/conf/config.yaml") // 打开配置文件
		// 错误处理
		if err != nil {
			log.Fatalf("Failed to open config file: %v", err)
		}

		// 关闭文件
		defer func(f *file.File) {
			err2 := f.Close()
			if err2 != nil {
				logs.Print("Failed to close config file: ", err2).Ln()
			}
		}(f)

		// 获取文件大小
		fileInfo, _ := f.Stat()
		// 创建一个切片
		size := fileInfo.Size()
		data := make([]byte, size)
		_, err = f.Read(data) // 读取文件内容到 data 切片
		// 错误处理
		if err != nil {
			log.Fatalf("Failed to read config file: %v", err)
		}

		//  将 YAML 数据解码到 conf 结构体中
		err = yaml.Unmarshal(data, conf)
		// 错误处理
		if err != nil {
			log.Fatalf("Failed to unmarshal YAML: %v", err)
		}

		// 打印解析后的配置
		//log.Printf("Parsed Config: %+v\n", conf.Db.Mysql[0])
		if conf == nil {
			// 设置默认值 1
			if conf.Queue.MaxConnNum <= 0 {
				conf.Queue.MaxConnNum = 1
			}
		}
	}

}
