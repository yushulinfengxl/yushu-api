package config

import (
	"yushu/box/datastore"
	"yushu/box/utility/singleton"
)

// Config 配置结构体
type Config struct {
	Server   `json:"server" yaml:"server"`     // Server 服务配置
	Logger   `json:"logger" yaml:"logger"`     // Logger 日志配置
	Database `json:"database" yaml:"database"` // Database 数据库配置
	Tencent  `json:"tencent" yaml:"tencent"`   // Tencent 腾讯云配置
}

// Server 服务配置结构体
type Server struct {
	Address           string `json:"address" yaml:"address"`                           // 服务监听地址
	MaxConn           int    `json:"max_conn" yaml:"max_conn"`                         // 最大连接数
	AccessKey         string `json:"access_key" yaml:"access_key"`                     // 访问密钥
	ReadTimeout       string `json:"read_timeout" yaml:"read_timeout"`                 // 读超时时间
	StaticBaseUri     string `json:"static_base_uri" yaml:"static_base_uri"`           // 静态资源路由前缀
	StaticPath        string `json:"static_path" yaml:"static_path"`                   // 静态资源目录
	OpenapiPath       string `json:"openapi_path" yaml:"openapi_path"`                 // OpenAPI 路径
	SwaggerPath       string `json:"swagger_path" yaml:"swagger_path"`                 // Swagger UI 路径
	MaxHeaderBytes    string `json:"max_header_bytes" yaml:"max_header_bytes"`         // 请求头的最大字节数
	ClientMaxBodySize string `json:"client_max_body_size" yaml:"client_max_body_size"` // 请求体最大字节数
	// Logging 配置
	LogFilePath      string `json:"log_filepath" yaml:"log_filepath"`             // 日志文件存储目录路径，建议使用绝对路径。默认为空，表示关闭
	LogStdout        bool   `json:"log_stdout" yaml:"log_stdout"`                 // 日志是否输出到终端。默认为true
	ErrorStack       bool   `json:"error_stack" yaml:"error_stack"`               // 当Server捕获到异常时是否记录堆栈信息到日志中。默认为true
	ErrorLogEnabled  bool   `json:"error_log_enabled" yaml:"error_log_enabled"`   // 是否记录异常日志信息到日志中。默认为true
	ErrorLogPattern  string `json:"error_log_pattern" yaml:"error_log_pattern"`   // 异常错误日志文件格式
	AccessLogEnabled bool   `json:"access_log_enabled" yaml:"access_log_enabled"` // 是否记录访问日志
	AccessLogPattern string `json:"access_log_pattern" yaml:"access_log_pattern"` // 访问日志文件格式
}

// Logger 配置
type Logger struct {
	Path   string `json:"path" yaml:"path"`     // 日志文件存储目录路径，建议使用绝对路径。默认为空，表示关闭
	File   string `json:"file" yaml:"file"`     // 日志文件格式，支持日期
	Level  string `json:"level" yaml:"level"`   // 日志级别
	Stdout bool   `json:"stdout" yaml:"stdout"` // 是否输出到终端
}

// Database 数据库配置
type Database struct {
	Default string `json:"default" yaml:"default"` // 默认数据库
	Mysql   Mysql  `json:"mysql" yaml:"mysql"`     // mysql 数据库配置
	Redis   Redis  `json:"redis" yaml:"redis"`     // redis 配置
}

// Mysql 数据库配置结构体
type Mysql struct {
	TablePrefix           string      `json:"table_prefix" yaml:"table_prefix"`                         // 表前缀
	MaxIdleConn           int         `json:"max_idle_conn" yaml:"max_idle_conn"`                       // 最大空闲连接数
	MaxOpenConn           int         `json:"max_open_conn" yaml:"max_open_conn"`                       // 最大连接数
	ConnMaxLifetimeSecond int         `json:"conn_max_lifetime_second" yaml:"conn_max_lifetime_second"` // 连接最大存活时间
	List                  []MysqlItem `json:"list" yaml:"list"`                                         // 数据库配置列表
}

// MysqlItem MysqlList 数据库配置列表
type MysqlItem struct {
	Host     string `json:"host" yaml:"host"`         // 数据库地址
	User     string `json:"user" yaml:"user"`         // 数据库用户名
	Password string `json:"password" yaml:"password"` // 数据库密码
	Database string `json:"database" yaml:"database"` // 数据库名
	Port     string `json:"port" yaml:"port"`         // 数据库端口
	Timeout  string `json:"timeout" yaml:"timeout"`   // 连接超时时间
}

// Redis 配置结构体
type Redis struct {
	Host            string `json:"host" yaml:"host"`                         // 地址
	Db              int    `json:"db" yaml:"db"`                             // 数据库
	IdleTimeout     string `json:"idle_timeout" yaml:"idle_timeout"`         // 连接最大空闲时间，使用时间字符串例如30s/1m/1d
	MaxConnLifetime string `json:"max_connLifetime" yaml:"max_connLifetime"` // 连接最长存活时间，使用时间字符串例如30s/1m/1d
	WaitTimeout     string `json:"wait_timeout" yaml:"wait_timeout"`         // 等待连接池连接的超时时间，使用时间字符串例如30s/1m/1d
	DialTimeout     string `json:"dial_timeout" yaml:"dial_timeout"`         // TCP连接的超时时间，使用时间字符串例如30s/1m/1d
	ReadTimeout     string `json:"read_timeout" yaml:"read_timeout"`         // TCP的Read操作超时时间，使用时间字符串例如30s/1m/1d
	WriteTimeout    string `json:"write_timeout" yaml:"write_timeout"`       // TCP的Write操作超时时间，使用时间字符串例如30s/1m/1d
	MaxActive       int    `json:"max_active" yaml:"max_active"`             // 最大连接数
}

// Tencent 配置结构体
type Tencent struct {
	Dnspod Dnspod `json:"dnspod" yaml:"dnspod"` // 腾讯云DNS配置
}

// Dnspod 腾讯云DNS配置结构体
type Dnspod struct {
	SecretId  string `json:"secretId" yaml:"secretId"`   // SecretId 腾讯云SecretId
	SecretKey string `json:"secretKey" yaml:"secretKey"` // SecretKey 腾讯云SecretKey
	DnsPodUrl string `json:"dnsPodUrl" yaml:"dnsPodUrl"` // DnsPodUrl 腾讯云DNS API地址
}

var configLazySingleton singleton.Lazy

func New() *Config {
	ins := configLazySingleton.Instance(&Config{})
	return (*ins).(*Config)
}

func init() {
	// 拿到配置实例
	conf := New()
	// 初始化配置文件
	Yaml(conf, "config.yaml")
	// 向datastore中注册配置实例
	datastore.Set("config", conf)
}
