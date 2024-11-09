package logger

import "log"

// 定义权限常量，每个常量对应一个二进制位
const (
	ErrorType  = 0x000 // ErrorType write file error type
	PanicType  = 0x001 // PanicType write file panic type
	StdoutType = 0x010 // StdoutType output to console
	TestType   = 0x100
)

// Logger 结构体（假设用来记录日志）
type Logger struct {
}

// Info 函数，用于记录日志并统计infoType中的权限个数
func Info(infoType int, args ...interface{}) {
	// 统计infoType中有多少个权限（标志）被设置了
	//countPermissions(infoType)
	// 输出日志
	log.Println(args...)
}

// countPermissions 函数，用于统计infoType中有多少个标志位被设置了
func countPermissions(infoType int) int {
	count := 0

	// 判断各个标志是否设置
	if infoType&ErrorType != 0 {
		count++
	}
	if infoType&PanicType != 0 {
		count++
	}
	if infoType&StdoutType != 0 {
		count++
	}

	return count
}

// SwitchOption 函数，设置日志前缀并执行对应的操作
func SwitchOption(infoType int) {
	// 检查infoType并执行不同的操作
	switch {
	case infoType&ErrorType != 0:
		log.SetPrefix("[ERROR]") // 设置为错误日志
	case infoType&PanicType != 0:
		log.SetPrefix("[PANIC]") // 设置为panic日志
	case infoType&StdoutType != 0:
		log.SetPrefix("[INFO]") // 设置为普通信息日志
	default:
		log.SetPrefix("[UNKNOWN]") // 默认情况下，未知类型
	}
}
