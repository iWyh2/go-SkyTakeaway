package vo

// OrderPaymentVO 订单支付返回数据模型
type OrderPaymentVO struct {
	// 随机字符串
	NonceStr string `json:"nonceStr"`
	// 签名
	PaySign string `json:"paySign"`
	// 时间戳
	TimeStamp string `json:"timeStamp"`
	// 签名算法
	SignType string `json:"signType"`
	// 统一下单接口返回的 prepay_id 参数值
	PackageStr string `json:"packageStr"`
}
