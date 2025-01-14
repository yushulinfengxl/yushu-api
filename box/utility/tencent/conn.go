package tencent

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"yushu/box/utility/singleton"
)

type Connected struct {
	SecretId      string
	SecretKey     string
	DnsPodUrl     string
	Credential    *common.Credential
	ClientProfile *profile.ClientProfile
}

var connectLazySingleton singleton.Lazy

func Conn() *Connected {
	ins := connectLazySingleton.Instance(&Connected{})

	return (*ins).(*Connected)
}

func init() {
	c := Conn()
	c.SecretId = ""
	c.SecretKey = ""
	c.DnsPodUrl = "dnspod.tencentcloudapi.com"
	c.Credential = common.NewCredential(
		c.SecretId,
		c.SecretKey,
	)
	// 实例化一个client选项，可选的，没有特殊需求可以跳过
	c.ClientProfile = profile.NewClientProfile()
}
