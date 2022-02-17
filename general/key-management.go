package general

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"dfa/entity"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"math/big"
)

func GenerateECDSAKey() (private, public []byte, err error) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
	if err != nil {
		panic(err)
	}
	privateBytes, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		return []byte{}, []byte{}, err
	}
	block1 := pem.Block{Type: "FRANK ECDSA PRIVATEKEY", Bytes: privateBytes}
	prvKey := pem.EncodeToMemory(&block1)

	publicKey := privateKey.PublicKey
	X509PublicStream, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		return []byte{}, []byte{}, err
	}
	block2 := &pem.Block{
		Type:  "FRANK ECDSA PUBLICKEY",
		Bytes: X509PublicStream,
	}
	pubKey := pem.EncodeToMemory(block2)
	return prvKey, pubKey, nil
}

func MakeSignature(plainText []byte, pri []byte) (entity.Signature, error) {
	signature := entity.Signature{
		R: []byte{},
		S: []byte{},
	}
	h := sha256.New()
	h.Write(plainText)
	hashed := h.Sum(nil)
	block, _ := pem.Decode(pri)
	if block == nil {
		return signature, fmt.Errorf("error: MakeSignature: private key block is nil")
	}
	privateKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return signature, fmt.Errorf("error: MakeSignature: unable to parse private key: %w", err)
	}
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hashed)
	rText, err := r.MarshalText()
	if err != nil {
		return signature, err
	}
	sText, err := s.MarshalText()
	if err != nil {
		return signature, err
	}
	signature.R, signature.S = rText, sText
	return signature, nil
}
func VerifySignature(plainText []byte, signature entity.Signature, pub []byte) (bool, error) {
	block, _ := pem.Decode(pub)
	if block == nil {
		return false, fmt.Errorf("error: Verify Signature: public key error")
	}
	publicInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return false, fmt.Errorf("error: public key error: %w", err)
	}
	publicKey := publicInterface.(*ecdsa.PublicKey)

	hash256 := sha256.New()
	hash256.Write(plainText)
	hashed := hash256.Sum(nil)
	var r, s big.Int
	r.UnmarshalText(signature.R)
	s.UnmarshalText(signature.S)
	return ecdsa.Verify(publicKey, hashed, &r, &s), nil
}

func FormatSignature(s entity.Signature) string {
	sJson, _ := json.Marshal(s)
	sStr, _ := BytesTostr(sJson)
	return sStr
}

func ParseSignature(addr string) entity.Signature {
	sStr, _ := StrToBytes(addr)
	var s entity.Signature
	json.Unmarshal(sStr, &s)
	return s
}
