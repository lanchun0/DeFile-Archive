package entity

import "time"

type Data struct {
	ID              string `json:"id"`
	Owner           string `json:"owner"`
	PermissionLevel uint8  `json:"permissionlevel"`
	Tradable        bool   `json:"tradable"`
	Price           uint64 `json:"price"`
	Signature       string `json:"signature"`
	MeteData        MeteData
	PermissionList  []PermissionList
	// History         []DataHistory
}

type MeteData struct {
	FileName   string    `json:"filename"`
	HashDigest string    `json:"hashdigest"`
	Signer     string    `json:"signer"`
	Size       uint64    `json:"size"`
	TimeStamp  time.Time `json:"timestamp"`
}

type PermissionList struct {
	Address    string `json:"address"`
	Permission uint8  `json:"permission"`
}

// type DataHistory struct {
// 	HashDigest string `json:"hashdigest"`
// 	Signer     string `json:"signer"`
// 	MeteData   MeteData
// 	Signature  Signature
// }
