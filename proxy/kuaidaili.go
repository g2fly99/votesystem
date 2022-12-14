package proxy

import (
	"github.com/beego/beego/v2/core/logs"
)

// 私密代理使用示例

// 接口鉴权说明：
// 接口鉴权方式为必填项, 目前支持的鉴权方式有"token" 和 "hmacsha1"两种
// 可选值为TOKEN和HmacSha1 或直接传"token"或"hmacsha1"

// 返回值说明:
// 所有返回值都包括两个值，第一个为目标值，第二个为error类型, 值为nil说明成功，不为nil说明失败

func UseDps(appId, appKey string) {
	auth := Auth{SecretID: appId, SecretKey: appKey}
	client := Client{Auth: auth}

	// 提取私密代理, 参数有: 提取数量、鉴权方式及其他参数(放入map[string]interface{}中, 若无则传入nil)
	// (具体有哪些其他参数请参考帮助中心: "https://www.kuaidaili.com/doc/api/getdps/")
	params := map[string]interface{}{"format": "json", "area": "北京,上海", "f_loc": 1, "dedup": 1}
	ips, err := client.GetDps(10, HmacSha1, params)
	if err != nil {
		logs.Error(err)
	}

	logs.Debug("ips: %v", ips)

	// 获取订单访问代理IP的鉴权信息
	proxyAuthorization, err := client.GetProxyAuthorization(1, HmacSha1)
	logs.Debug("proxyAuthorization: %v", proxyAuthorization)

