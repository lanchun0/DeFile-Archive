package general

import (
	"crypto/sha256"
	"fmt"
	"io"
	"mime/multipart"
)

func MakeHashDigest(f multipart.File) (hash string, err error) {
	ha := sha256.New()
	if _, err := io.Copy(ha, f); err != nil {
		return "", fmt.Errorf("failed to create hash digest: %v", err)
	}
	hash = Bytes2String(ha.Sum(nil))
	return hash, nil
}

// func FillInData(f multipart.File, h multipart.FileHeader, token string) (entity.Data, error) {
// 	pub, priv, err := service.ParseToken(token)
// 	if err != nil {

// 	}

// 	id, err := MakeUUID()
// 	if err != nil {
// 		return entity.Data{}, err
// 	}
// 	time := GenerateTimeStamp()
// 	digest, err := makeHashDigest(f)
// 	if err != nil {
// 		return entity.Data{}, err
// 	}
// 	data := entity.Data{
// 		ID:         id,
// 		HashDigest: digest,
// 		MeteData: entity.MeteData{
// 			FileName:  h.Filename,
// 			Size:      h.Size,
// 			TimeStamp: time,
// 		},
// 	}
// 	return data, nil
// }
