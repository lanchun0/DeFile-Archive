package smartcontract

import (
	"time"
)

type Permission byte
type PermissionLevel byte

const (
	Administrator Permission = 0b00000001
	Writer        Permission = 0b00000010
	Reader        Permission = 0b00000100
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
	PublicKey  string   `json:"publickey"`
	PrivateKey string   `json:"privatekey"`
	Asset      []string `json:"asset"`
	Logs       []LogInfo
}

type LogInfo struct {
	TimeStamp  time.Time `json:"timestamp"`
	PublicKey  string    `json:"publickey"`
	Method     string    `json:"method"`
	TargetFile string    `json:"targetfile"` // File Hash
}

type MeteData struct {
}

type PermissionList struct {
	PublicKey  string     `json:"publickey"`
	Permission Permission `json:"permission"`
}

type Signature struct {
	R []byte `json:"r"`
	S []byte `json:"s"`
}
