package initialize

import (
	"github.com/smartwalle/alipay/v3"
	"github.com/zhenghuajing/fresh_shop/global"
)

func Alipay() {
	alipayCfg := global.Config.Alipay
	// appId
	appID := alipayCfg.AppID
	// 应用私钥
	privateKey := alipayCfg.PrivateKey

	global.AlipayClient, _ = alipay.New(appID, privateKey, false)

	// 载入证书文件
	global.AlipayClient.LoadAppPublicCertFromFile(alipayCfg.AppPublicCertPath)
	global.AlipayClient.LoadAliPayPublicCertFromFile(alipayCfg.AliPayPublicCertPath)
	global.AlipayClient.LoadAliPayRootCertFromFile(alipayCfg.AliPayRootCertPath)
}
