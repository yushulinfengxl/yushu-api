package whois

import (
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	dnspod "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dnspod/v20210323"
	"yushu/opt/utility/tencent"
)

func GET(domain string) (res string, err error) {
	// 拿到实例对象
	conn := tencent.Conn()
	credential := conn.Credential
	// 实例化一个client选项，可选的，没有特殊需求可以跳过
	cpf := conn.ClientProfile
	// 连接接口地址
	cpf.HttpProfile.Endpoint = conn.DnsPodUrl
	// 实例化要请求产品的client对象,clientProfile是可选的
	client, _ := dnspod.NewClient(credential, "", cpf)

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := dnspod.NewDescribeDomainWhoisRequest()

	request.Domain = common.StringPtr(domain)

	// 返回的resp是一个DescribeDomainWhoisResponse的实例，与请求对象对应
	response, err := client.DescribeDomainWhois(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}

	res = response.ToJsonString()
	return
}
