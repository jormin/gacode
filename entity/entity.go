package entity

// Data Save data
type Data struct {
	Accounts []*Account `json:"accounts" desc:"account list"`
}

// Account account of Google Authenticator
type Account struct {
	Name       string `json:"name" desc:"account name"`
	Secret     string `json:"secret" desc:"account secret"`
	QRCode     string `json:"qrcode" desc:"account QR code"`
	CreateTime int64  `json:"create_time" desc:"create time"`
	UpdateTime int64  `json:"update_time" desc:"update time"`
}
