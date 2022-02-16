package entity

import "time"

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
	FileID     string    `json:"fileid"`
	HashDigest string    `json:"hashdigest"`
}

type Asset struct {
	FileID     string `json:"fileid"`
	Permission byte   `json:"permission"`
}
