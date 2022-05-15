package dto

type Credentials struct {
	PrivateKey string `json:"privatekey"`
}

type TokenIdentity struct {
	Token string `json:"token"`
}

type TopupRequest struct {
	Token  string `json:"token"`
	Amount uint64 `json:"amount"`
}

type WithDrawRequest struct {
	Token  string `json:"token"`
	Amount uint64 `json:"amount"`
}

type ApproveRequest struct {
	Token string `json:"token"`
	Price uint64 `json:"price"`
}

type UploadFileRequest struct {
	Token           string `json:"token"`
	PermissionLevel string `json:"permissionlevel"`
	Price           uint64 `json:"price"`
	Tradable        bool   `json:"tradable"`
}

type WriteFileRequest struct {
	Token string `json:"token"`
	ID    string `json:"id"`
}

type BuyAFileRequest struct {
	Token string `json:"token"`
	ID    string `json:"id"`
}

type ShareFileRequest struct {
	Token      string `json:"token"`
	ID         string `json:"id"`
	To         string `json:"to"`
	Permission string `json:"permission"`
}

type QueryFileRequest struct {
	ID string `json:"id"`
}
