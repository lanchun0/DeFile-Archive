package general

import (
	"reflect"
	"unsafe"
	// shell "github.com/ipfs/go-ipfs-api"
)

// var sh *shell.Shell

// func BytesTostr(key []byte) (string, error) {
// 	sh = shell.NewShell("localhost:5001")
// 	hash, err := sh.Add(bytes.NewBufferString(string(key)))
// 	if err != nil {
// 		return "N.A.", fmt.Errorf("error: Failed to upload to IPFS: %w", err)
// 	}
// 	return hash, nil
// }

// func StrToBytes(addr string) (key []byte, err error) {
// 	sh = shell.NewShell("localhost:5001")
// 	read, err := sh.Cat(addr)
// 	if err != nil {
// 		return []byte{}, fmt.Errorf("error: Failed to download from IPFS: %w", err)
// 	}
// 	key, err = ioutil.ReadAll(read)
// 	if err != nil {
// 		return []byte{}, fmt.Errorf("error: Failed to download from IPFS: %w", err)
// 	}
// 	return key, nil
// }

func String2Bytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

func Bytes2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
