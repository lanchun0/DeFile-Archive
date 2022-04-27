package general

import (
	"bytes"
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func GenerateECSDAKey() (pub, priv, addr string) {
	privateKey, _ := crypto.GenerateKey()
	privateKeyBytes := crypto.FromECDSA(privateKey)
	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	pub, priv = hexutil.Encode(publicKeyBytes), hexutil.Encode(privateKeyBytes)
	return pub, priv, address
}

func MakeSignature(priv string, plaintext []byte) (signature string, err error) {
	privateKey, err := crypto.HexToECDSA(priv[2:])
	if err != nil {
		return "", err
	}
	hash := crypto.Keccak256Hash(plaintext)
	sigBytes, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		return "", err
	}
	signature = hexutil.Encode(sigBytes)
	return signature, nil
}

func VerifySignature(pub, signature string, plaintext []byte) bool {
	sigBytes, err := hexutil.Decode(signature)
	if err != nil {
		return false
	}
	pubBytes, err := hexutil.Decode(pub)
	if err != nil {
		return false
	}
	hash := crypto.Keccak256Hash(plaintext)
	sigPublicKey, err := crypto.Ecrecover(hash.Bytes(), sigBytes)
	if err != nil {
		return false
	}
	matches := bytes.Equal(sigPublicKey, pubBytes)
	if !matches {
		return false
	}
	signatureNoRecoverID := sigBytes[:len(sigBytes)-1] // remove recovery id
	return crypto.VerifySignature(pubBytes, hash.Bytes(), signatureNoRecoverID)
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
