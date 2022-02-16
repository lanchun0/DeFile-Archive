package contractsimulation

import (
	"dfa/entity"
	"dfa/general"
	"fmt"
	"time"
)

type BehaviorContract interface {
	Register() (entity.Behavior, error)
	Login(publickey string, s entity.Signature) (entity.Behavior, error)
	FindBehavior(publickey string) (entity.Behavior, bool, int)
	RecordAccess(publickey string, log entity.LogInfo) (entity.Behavior, error)
	Authorize(hostkey, customerkey, fileid string, p entity.Permission) (bool, error)
}

type behaviorChain struct {
	behaviorChain []entity.Behavior
}

func NewBehaviorChain() BehaviorContract {
	return &behaviorChain{
		behaviorChain: []entity.Behavior{},
	}
}

func (contract *behaviorChain) Register() (entity.Behavior, error) {
	pri, pub, err := general.GenerateECDSAKey()
	if err != nil {
		return entity.Behavior{}, err
	}

	priStr, err := general.BytesTostr(pri)
	pubStr, err := general.BytesTostr(pub)
	if err != nil {
		return entity.Behavior{}, err
	}

	b := entity.Behavior{
		PublicKey:  pubStr,
		PrivateKey: priStr,
		Assets:     []entity.Asset{},
		Logs:       []entity.LogInfo{},
	}
	l := entity.LogInfo{
		TimeStamp:  time.Now(),
		PublicKey:  pubStr,
		Method:     "register",
		FileID:     "N.A.",
		HashDigest: "N.A.",
	}
	b.Logs = append(b.Logs, l)
	contract.behaviorChain = append(contract.behaviorChain, b)
	return b, err
}

func (contract *behaviorChain) Login(publickey string, s entity.Signature) (entity.Behavior, error) {
	var bh entity.Behavior
	exist := false
	for _, b := range contract.behaviorChain {
		if b.PublicKey == publickey {
			bh, exist = b, true
			break
		}
	}
	if !exist {
		return entity.Behavior{}, fmt.Errorf("error: %s", "the user does not exist")
	}
	priByte, err := general.StrToBytes(bh.PrivateKey)
	pubByte, err := general.StrToBytes(bh.PublicKey)
	if err != nil {
		return entity.Behavior{}, fmt.Errorf("error: %s", "missing keys")
	}
	valid, err := general.VerifySignature(priByte, s, pubByte)
	if !valid || err != nil {
		return entity.Behavior{}, fmt.Errorf("error: login failed on behavior blockchain: %w", err)
	}
	bh.PrivateKey = ""
	return bh, nil
}

func (contract *behaviorChain) FindBehavior(publickey string) (entity.Behavior, bool, int) {
	var behavior entity.Behavior
	var index int
	exist := false
	for i, b := range contract.behaviorChain {
		if b.PublicKey == publickey {
			index, behavior, exist = i, b, true
			break
		}
	}
	if !exist {
		return behavior, false, -1
	}
	behavior.PrivateKey = ""
	return behavior, true, index
}

// do not need index in real contract
func (contract *behaviorChain) RecordAccess(publickey string, log entity.LogInfo) (entity.Behavior, error) {
	behavior, exist, index := contract.FindBehavior(publickey)
	if !exist || index < 0 || index > len(contract.behaviorChain) {
		return behavior, fmt.Errorf("error: cannot find user %s on Behavior Blockchain", publickey)
	}
	condition := make(map[string]entity.Permission)
	condition["download"], condition["modify"], condition["create"] = entity.Reader, entity.Writer, entity.Owner
	exist = false
	// Find the asset and make sure the permission is accessible
	for _, a := range behavior.Assets {
		if a.FileID == log.FileID && a.Permission >= byte(condition[log.Method]) {
			behavior.Logs, exist = append(behavior.Logs, log), true
			break
		}
	}
	if !exist {
		return behavior, fmt.Errorf("error: cannot find asset %s on Behavior Blockchain", log.FileID)
	}
	contract.behaviorChain[index] = behavior
	return behavior, nil

}

func (contract *behaviorChain) Authorize(hostkey, customerkey, fileid string, p entity.Permission) (bool, error) {
	hostB, existB, hosti := contract.FindBehavior(hostkey)
	if !existB {
		return false, fmt.Errorf("error: user %s does not exist", hostkey)
	}
	hostC, existC, customeri := contract.FindBehavior(customerkey)
	if !existC {
		return false, fmt.Errorf("error: user %s does not exist", customerkey)
	}

	asset := entity.Asset{
		FileID:     fileid,
		Permission: byte(p),
	}
	hostinfo := entity.LogInfo{
		TimeStamp:  time.Now(),
		PublicKey:  customerkey,
		Method:     "authorize",
		FileID:     fileid,
		HashDigest: string(p),
	}
	custinfo := entity.LogInfo{
		TimeStamp:  time.Now(),
		PublicKey:  hostkey,
		Method:     "be authorized",
		FileID:     fileid,
		HashDigest: string(p),
	}

	for _, a := range hostB.Assets {
		if a.FileID == fileid {
			if a.Permission < byte(entity.Administrator) || a.Permission < byte(p) {
				return false, fmt.Errorf("error: user %s has no permission", hostkey)
			} else {
				for _, as := range hostC.Assets {
					if as.FileID == fileid && as.Permission >= byte(p) {
						return false, fmt.Errorf("error: user %s has already been authoriszd", customerkey)
					}
				}
				contract.behaviorChain[customeri].Assets = append(contract.behaviorChain[customeri].Assets, asset)
				contract.behaviorChain[customeri].Logs = append(contract.behaviorChain[customeri].Logs, custinfo)
				contract.behaviorChain[hosti].Logs = append(contract.behaviorChain[hosti].Logs, hostinfo)
				return true, nil
			}
		}
	}

	return false, fmt.Errorf("error: user %s has no file", hostkey)
}
