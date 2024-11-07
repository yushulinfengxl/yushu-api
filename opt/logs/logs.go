package logs

import (
	"fmt"
	"time"
)

type Values []interface{}

func (v *Values) SetVal(a ...interface{}) *Values {
	*v = append(*v, a...)
	return v
}

func (v *Values) Ln() *Values {
	v.Print("\n")
	return v
}

func (v *Values) Print(a ...interface{}) *Values {
	Print(a...)
	return v
}

func Print(a ...interface{}) *Values {
	v := &Values{}
	v.SetVal(a...)
	fmt.Print(*v...)
	return v
}

func Date(a ...interface{}) *Values {
	v := &Values{}
	// 现在时间精确毫秒
	nowTime := time.Now().Format("2006-01-02 15:04:05.000")

	r := append([]interface{}{nowTime, "\t"}, a...)

	v.SetVal(r...)
	Print(*v...)

	return v
}

func Title(v ...interface{}) *Values {
	logs := &Values{}

	fmt.Print(v...)

	return logs
}
