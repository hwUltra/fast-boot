package pay

// WxConf 微信支付配置
type WxConf struct {
	AppId      string `json:"app_id"`      //公众号 APPID
	AppMiniId  string `json:"app_mini_id"` //小程序 APPID
	APPId      string `json:"appId"`       //APP   APPID
	MchId      string `json:"mchId"`       //微信商户id
	SerialNo   string `json:"serialNo"`    //商户证书的证书序列号
	ApiV3Key   string `json:"apiV3Key"`    //apiV3Key，商户平台获取
	PrivateKey string `json:"privateKey"`  //privateKey：私钥 apiclient_key.pem 读取后的内容
	AppSecret  string `json:"appSecret"`
	NotifyUrl  string `json:"notifyUrl"` //支付通知回调服务端地址
}

// AliConf 微信支付配置
type AliConf struct {
	AppId      string `json:"app_id"`
	PrivateKey string `json:"private_key"`
	PublicKey  string `json:"public_key"`
	IsProd     bool   `json:"is_prod"`
	NotifyUrl  string `json:"notifyUrl"`
}
