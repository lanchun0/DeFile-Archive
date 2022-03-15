package dto

type Credentials struct {
	PublicKey  string `form:"publickey"`
	PrivateKey string `form:"privatekey"`
}
