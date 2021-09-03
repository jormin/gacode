package entity

// Data
type Data struct {
	Accounts []Account `json:"accounts" desc:"账户信息"`
}

// Account 账户
type Account struct {
	Name       string `json:"name" desc:"账户名称"`
	Secret     string `json:"secret" desc:"秘钥"`
	QRCode     string `json:"qrcode" desc:"二维码"`
	CreateTime int64  `json:"create_time" desc:"创建时间"`
	UpdateTime int64  `json:"update_time" desc:"更新时间"`
}
