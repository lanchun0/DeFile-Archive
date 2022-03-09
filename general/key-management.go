package general

import (
	"bytes"
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func GenerateECDSAKey() (private, public []byte, addr string) {
	privateKey, _ := crypto.GenerateKey()
	privateKeyBytes := crypto.FromECDSA(privateKey)
	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	return publicKeyBytes, privateKeyBytes, address
}

func MakeSignature(priv, plaintext []byte) (signature []byte, err error) {
	privStr := hexutil.Encode(priv)
	privateKey, err := crypto.HexToECDSA(privStr[2:])
	if err != nil {
		return []byte{}, err
	}
	hash := crypto.Keccak256Hash(plaintext)
	signature, err = crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		return []byte{}, err
	}
	return signature, nil
}

func VerifySignature(pub, plaintext, signature []byte) bool {
	hash := crypto.Keccak256Hash(plaintext)
	sigPublicKey, err := crypto.Ecrecover(hash.Bytes(), signature)
	if err != nil {
		return false
	}
	matches := bytes.Equal(sigPublicKey, pub)
	if !matches {
		return false
	}
	signatureNoRecoverID := signature[:len(signature)-1] // remove recovery id
	return crypto.VerifySignature(pub, hash.Bytes(), signatureNoRecoverID)
}

// func FormatSignature(s entity.Signature) string {
// 	sJson, _ := json.Marshal(s)
// 	sStr := Bytes2String(sJson)
// 	return sStr
// }

// func ParseSignature(addr string) entity.Signature {
// 	sStr := String2Bytes(addr)
// 	var s entity.Signature
// 	json.Unmarshal(sStr, &s)
// 	return s
// }
