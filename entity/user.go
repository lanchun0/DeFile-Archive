package entity

type User struct {
	Address string `json:"address"`
	Balance uint64 `json:"balance"`
	// PrivateKey string  `json:"privatekey"`
	Assets []Asset `json:"assets"`
	// Logs       []LogInfo
}

// type LogInfo struct {
// 	TimeStamp  time.Time `json:"timestamp"`
// 	PublicKey  string    `json:"publickey"`
// 	Method     string    `json:"method"`
// 	FileID     string    `json:"fileid"`
// 	HashDigest string    `json:"hashdigest"`
// }

type Asset struct {
	FileID     string `json:"fileid"`
	Permission uint8  `json:"permission"`
}
