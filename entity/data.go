package entity

import "time"

type Data struct {
	ID              string `json:"id"`
	HashDigest      string `json:"hashdigest"`
	Owner           string `json:"owner"`
	PermissionLevel uint   `json:"permissionlevel"`
	MeteData        MeteData
	PermissionList  []PermissionList
	Signature       Signature
	History         []DataHistory
}

type MeteData struct {
	FileName  string    `json:"filename"`
	Size      int       `json:"size"`
	TimeStamp time.Time `json:"timestamp"`
}

type PermissionList struct {
	PublicKey  string `json:"publickey"`
	Permission uint   `json:"permission"`
}

type Signature struct {
	R []byte `json:"r"`
	S []byte `json:"s"`
}

type DataHistory struct {
	HashDigest string `json:"hashdigest"`
	Signer     string `json:"signer"`
	MeteData   MeteData
	Signature  Signature
}
