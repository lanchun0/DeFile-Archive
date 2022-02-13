package entity

import "time"

type Data struct {
	HashDigest      string          `json:"hashdigest"`
	Owner           string          `json:"owner"`
	PermissionLevel PermissionLevel `json:"permissionlevel"`
	MeteData        MeteData
	PermissionList  []PermissionList
	Signature       Signature
	History         []Data
}

type MeteData struct {
	FileName  string    `json:"filename"`
	Size      int       `json:"size"`
	TimeStamp time.Time `json:"timestamp"`
}

type PermissionList struct {
	PublicKey  string     `json:"publickey"`
	Permission Permission `json:"permission"`
}

type Signature struct {
	R []byte `json:"r"`
	S []byte `json:"s"`
}
