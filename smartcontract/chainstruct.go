package smartcontract

import (
	"time"
)

type Permission byte
type PermissionLevel byte

const (
	Owner Permission = 0b10000000 >> iota
	Administrator
	Writer
	Reader
)
const (
	L_0 PermissionLevel = iota
	L_1
	L_2
	L_3
	L_4
)

type Data struct {
	HashDigest      string          `json:"hashdigest"`
	Owner           string          `json:"owner"`
	PermissionLevel PermissionLevel `json:"permissionlevel"`
	MeteData        MeteData
	PermissionList  []PermissionList
	Signature       Signature
	History         []Data
}

type Behavior struct {
	PublicKey  string  `json:"publickey"`
	PrivateKey string  `json:"privatekey"`
	Assets     []Asset `json:"assets"`
	Logs       []LogInfo
}

type LogInfo struct {
	TimeStamp  time.Time `json:"timestamp"`
	PublicKey  string    `json:"publickey"`
	Method     string    `json:"method"`
	TargetFile string    `json:"targetfile"` // File Hash
}

type MeteData struct {
	FileName  string    `json:"filename"`
	Size      int       `json:"size"`
	TimeStamp time.Time `json:"timestamp"`
}

type PermissionList struct {
	PublicKey  string `json:"publickey"`
	Permission byte   `json:"permission"`
}

type Signature struct {
	R []byte `json:"r"`
	S []byte `json:"s"`
}

type Asset struct {
	FileHash   string `json:"filehash"`
	Permission byte   `json:"permission"`
}
