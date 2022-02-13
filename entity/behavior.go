package entity

import (
	"time"
)

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
