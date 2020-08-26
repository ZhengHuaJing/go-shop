package util

import (
	"fmt"
	"github.com/smartwalle/alipay/v3"
	"github.com/zhenghuajing/fresh_shop/global"
)

// 网站扫码支付
func WebPageAlipay(returnUrl, orderNo, nonceStr, money string) string {
	serverCfg := global.Config.Server
	pay := alipay.TradePagePay{}
	pay.NotifyURL = fmt.Sprintf("%s:%d/api/v1/pay", serverCfg.Host, serverCfg.HttpPort)
	pay.ReturnURL = returnUrl
	pay.Subject = "生鲜超市付款"
	pay.OutTradeNo = orderNo
	pay.ProductCode = "FAST_INSTANT_TRADE_PAY"
	pay.TotalAmount = money
	pay.PassbackParams = nonceStr

	url, err := global.AlipayClient.TradePagePay(pay)
	if err != nil {
		global.Log.Error(err)
	}

	return url.String()
}

// 手机客户端支付
func WapAlipay(returnUrl, orderNo, nonceStr, money string) string {
	serverCfg := global.Config.Server
	pay := alipay.TradeWapPay{}
	pay.NotifyURL = fmt.Sprintf("%s:%d/api/v1/pay", serverCfg.Host, serverCfg.HttpPort)
	pay.ReturnURL = returnUrl
	pay.Subject = "生鲜超市付款"
	pay.OutTradeNo = orderNo
	pay.TotalAmount = money
	pay.PassbackParams = nonceStr

	url, err := global.AlipayClient.TradeWapPay(pay)
	if err != nil {
		global.Log.Error(err)
	}

	return url.String()
}
