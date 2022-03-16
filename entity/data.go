package entity

import "time"

type Data struct {
	ID              string `json:"id"`
	HashDigest      string `json:"hashdigest"`
	Owner           string `json:"owner"`
	PermissionLevel uint8  `json:"permissionlevel"`
	Signature       string `json:"signature"`
	MeteData        MeteData
	PermissionList  []PermissionList
	// History         []DataHistory
}

type MeteData struct {
	FileName  string    `json:"filename"`
	Size      int64     `json:"size"`
	TimeStamp time.Time `json:"timestamp"`
}

type PermissionList struct {
	PublicKey  string `json:"publickey"`
	Permission uint8  `json:"permission"`
}

// type DataHistory struct {
// 	HashDigest string `json:"hashdigest"`
// 	Signer     string `json:"signer"`
// 	MeteData   MeteData
// 	Signature  Signature
// }
