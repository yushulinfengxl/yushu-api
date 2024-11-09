package utility

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"strconv"
	"yushu/box/config"
	"yushu/box/filesystem"
	"yushu/box/logger"
)

// yaml 默认配置文件模板
var yamlFileDefaultTemplateText = `# Server 配置
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
  log_filepath: "resource/log/server" # 日志文件存储目录路径，建议使用绝对路径。默认为空，表示关闭
  log_stdout: true # 日志是否输出到终端。默认为true
  error_stack: true # 当Server捕获到异常时是否记录堆栈信息到日志中。默认为true
  error_log_enabled: true # 是否记录异常日志信息到日志中。默认为true
  error_log_pattern: "error-{Ymd}.log" # 异常错误日志文件格式。默认为"error-{Ymd}.log"
  access_log_enabled: true # 是否记录访问日志。默认为false
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



# tencent 配置
tencent:
  dnspod:
    secretId: ""
    secretKey: ""
    dnsPodUrl: "dnspod.tencentcloudapi.com"`
// Yaml 读取配置文件
func Yaml(conf *config.Config, yamlFilepath string) {
	// 判断配置文件是否存在
	isYamlFileExists := filesystem.Exists(yamlFilepath)
	// 拿到默认配置文件字节
	yamlFileByte := []byte(yamlFileDefaultTemplateText)
	// 拿到yaml文件路径
	yamlFilepath = filesystem.ExecPath(yamlFilepath)
	// 创建或拿到文件
	file, err := filesystem.Create(yamlFilepath)
	if err != nil {
		// 创建或获取配置文件发生错误
		logger.Info(logger.ErrorType|logger.StdoutType, "创建或获取配置文件发生错误", err)
		return
	}
	// 判断配置文件是否存在
	if !isYamlFileExists {
		_, err = file.Write(yamlFileByte)
		if err != nil {
			logger.Info(logger.ErrorType|logger.StdoutType, "配置文件写入数据发生搓揉", err)
			return
		}
	}
	file, err = filesystem.Create(yamlFilepath)
	if err != nil {
		logger.Info(logger.ErrorType|logger.StdoutType, "获取配置文件发生错误", err)
		return
	}
	fileInfo, _ := file.Stat()
	buf := make([]byte, fileInfo.Size())
	_, err = file.Read(buf)
	if err != nil {
		logger.Info(logger.ErrorType|logger.StdoutType, "读取配置文件发生错误", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			logger.Info(logger.ErrorType|logger.StdoutType, "关闭配置文件发生错误", err)
		}
	}(file)
	log.Println("buf: ", strconv.Itoa(len(buf)))
	err = yaml.Unmarshal(buf, conf)
	if err != nil {
		logger.Info(logger.ErrorType|logger.StdoutType, "yaml配置文件解析发生错误", err)
		return
	}
}