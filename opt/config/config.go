package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"yushu/opt/file"
	"yushu/opt/logs"
	"yushu/opt/utility/singleton"
)

var DefaultConfig = `# Server 配置
server:
  address:     ":5000" # 服务监听地址
  max_conn:     100 # 最大连接数
  read_timeout:  "60s" # 读超时时间
  static_base_uri: "/static" # 静态资源路由前缀
  static_path: "/resource/public" # 静态资源目录
#  dumpRouterMap: false # 是否打印路由表
#  routeOverWrite: true # 是否覆盖路由
  openapi_path: "/api.json" # OpenAPI 路径
  swagger_path: "/swagger" # Swagger UI 路径
#  NameToUriType: 3 # 路由名称转URI 的策略
  max_header_bytes: "20KB" # 请求头的最大字节数
  client_max_body_size: "50MB" # 请求体最大字节数
  # Logging 配置
  log_path: "resource/log/server"                 # 日志文件存储目录路径，建议使用绝对路径。默认为空，表示关闭
  log_stdout: true               # 日志是否输出到终端。默认为true
  error_stack: true               # 当Server捕获到异常时是否记录堆栈信息到日志中。默认为true
  error_log_enabled: true               # 是否记录异常日志信息到日志中。默认为true
  error_log_pattern: "error-{Ymd}.log"  # 异常错误日志文件格式。默认为"error-{Ymd}.log"
  access_log_enabled: true              # 是否记录访问日志。默认为false
  access_log_pattern: "access-{Ymd}.log" # 访问日志文件格式。默认为"access-{Ymd}.log"


# Log 配置
logger:
  path: "resource/log/run" # 日志文件存储目录路径，建议使用绝对路径。默认为空，表示关闭
  file: "{Y-m-d}.log" # 日志文件格式，支持日期变量，例如：access-{Y-m-d}.log
  level: "all" # 日志级别，支持debug、info、warn、error、fatal、panic
  stdout: true # 是否输出到终端



# Database 配置
database:
  default: "mysql" # 默认数据库
  # Mysql 数据库配置示例
  mysql:
    # 多实例配置示例
    table_prefix: "yushu_" # 表前缀
    max_idle_conn: 10 # 最大空闲连接数
    max_open_conn: 100 # 最大连接数
    conn_max_lifetime_second: 3600 # 连接最大存活时间
    list:
      - host: "123.249.46.205"
        user: "yushu"
        password: "ZehsPXyKKYyec7az"
        database: "yushu"
        port: 3306
        timeout: "30s"
      - host: "localhost"
        user: "root"
        password: "admin"
        database: "dbname"
        port: 3306
        timeout: "30s"
      - host: "localhost2"
        user: "root"
        password: "admin"
        database: "dbname"
        port: 3306
        timeout: "30s"
  # Redis 配置示例
  redis:
    address: "127.0.0.1:6379" # 地址
    db: 1 # 数据库
    idle_timeout: "60s" # 连接最大空闲时间，使用时间字符串例如30s/1m/1d
    max_connLifetime: "90s" # 连接最长存活时间，使用时间字符串例如30s/1m/1d
    wait_timeout: "60s" # 等待连接池连接的超时时间，使用时间字符串例如30s/1m/1d
    dial_timeout: "30s" # TCP连接的超时时间，使用时间字符串例如30s/1m/1d
    read_timeout: "30s" # TCP的Read操作超时时间，使用时间字符串例如30s/1m/1d
    write_timeout: "30s" # TCP的Write操作超时时间，使用时间字符串例如30s/1m/1d
    max_active: 100 # 最大连接数

`

// Config 配置结构体
type Config struct {
}

// Server 服务配置结构体
type Server struct {
	Address           string `json:"address" yaml:"address"`                           // 服务监听地址
	MaxConn           int    `json:"max_conn" yaml:"max_conn"`                         // 最大连接数
	ReadTimeout       string `json:"read_timeout" yaml:"read_timeout"`                 // 读超时时间
	StaticBaseUri     string `json:"static_base_uri" yaml:"static_base_uri"`           // 静态资源路由前缀
	StaticPath        string `json:"static_path" yaml:"static_path"`                   // 静态资源目录
	OpenapiPath       string `json:"openapi_path" yaml:"openapi_path"`                 // OpenAPI 路径
	SwaggerPath       string `json:"swagger_path" yaml:"swagger_path"`                 // Swagger UI 路径
	MaxHeaderBytes    string `json:"max_header_bytes" yaml:"max_header_bytes"`         // 请求头的最大字节数
	ClientMaxBodySize string `json:"client_max_body_size" yaml:"client_max_body_size"` // 请求体最大字节数
}

var configLazySingleton singleton.Lazy

func NewApp() *Config {
	ins := configLazySingleton.Instance(&Config{})
	return (*ins).(*Config)
}

func init() {
	defer func() {
		if r := recover(); r != nil {
			logs.Print("Config recovered in f", r)
		}
	}()

	conf := NewApp() // 创建配置结构体实例

	// "./res/conf/config.yaml"
	filePath := "./config.yaml"

	// 判断如果文件不存在则创建默认配置文件
	if !file.Exists(filePath) {
		f, err := file.Open(filePath, os.O_APPEND)
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
		f, err := file.Open(filePath)

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
