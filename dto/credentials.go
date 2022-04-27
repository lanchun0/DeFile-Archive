package dto

import (
	"dfa/entity"
	"dfa/general"

	"fmt"
	"strings"
)

type Credentials struct {
	PublicKey  string `form:"publickey"`
	PrivateKey string `form:"privatekey"`
}

type FileCreateRequest struct {
	Token      string `form:"token"`
	Permission string `form:"permission"`
}

type OneFileRequest struct {
	ID string `json:"id"`
}

type ShareRequest struct {
	Token      string `form:"token"`
	ID         string `json:"id"`
	To         string `json:"to"`
	Permission string `json:"permission"`
}

type ViewData struct {
	ID              string `json:"id"`
	HashDigest      string `json:"hashdigest"`
	Owner           string `json:"owner"`
	PermissionLevel string `json:"permissionlevel"`
	Signature       string `json:"signature"`
	MeteData        MeteData
	PermissionList  []PermissionList
}
type MeteData struct {
	FileName  string `json:"filename"`
	Size      int64  `json:"size"`
	TimeStamp string `json:"timestamp"`
}

type PermissionList struct {
	PublicKey  string `json:"publickey"`
	Permission string `json:"permission"`
}

type ViewBehavior struct {
	PublicKey  string  `json:"publickey"`
	PrivateKey string  `json:"privatekey"`
	Assets     []Asset `json:"assets"`
	Logs       []LogInfo
}

type LogInfo struct {
	TimeStamp  string `json:"timestamp"`
	PublicKey  string `json:"publickey"`
	Method     string `json:"method"`
	FileID     string `json:"fileid"`
	HashDigest string `json:"hashdigest"`
}

type Asset struct {
	FileID     string `json:"fileid"`
	Permission string `json:"permission"`
}

var PL2String = map[uint8]string{
	uint8(entity.L_0): "L_0 (private)",
	uint8(entity.L_1): "L_1 (permissioned, no administrator)",
	uint8(entity.L_2): "L_2 (permissioned, multiple administrators)",
	uint8(entity.L_3): "L_3 (permissioned, no administrator)",
	uint8(entity.L_4): "L_4 (permissioned, multiple administrators)",
}

func Psermission2String(peimission uint8) string {
	str := ""
	if peimission&uint8(entity.Owner) != 0 {
		str = str + "owner"
	}
	if peimission&uint8(entity.Administrator) != 0 {
		str = str + "| administrator"
	}
	if peimission&uint8(entity.Writer) != 0 {
		str = str + "| writer"
	}
	if peimission&uint8(entity.Reader) != 0 {
		str = str + "| reader"
	}
	return str
}

func String2Permission(permission string) (uint8, error) {
	if strings.EqualFold("owner", permission) {
		return uint8(entity.Owner), nil
	} else if strings.EqualFold("administrator", permission) || strings.EqualFold("admin", permission) {
		return uint8(entity.Administrator), nil
	} else if strings.EqualFold("writer", permission) {
		return uint8(entity.Writer), nil
	} else if strings.EqualFold("reader", permission) {
		return uint8(entity.Reader), nil
	} else {
		return 0, fmt.Errorf("invalid permission: %s", permission)
	}
}

func String2PermissionLevel(pL string) (uint8, error) {
	if strings.EqualFold("L_0", pL) || strings.EqualFold("private", pL) {
		return uint8(entity.L_0), nil
	} else if strings.EqualFold("L_1", pL) {
		return uint8(entity.L_1), nil
	} else if strings.EqualFold("L_2", pL) {
		return uint8(entity.L_2), nil
	} else if strings.EqualFold("L_3", pL) {
		return uint8(entity.L_3), nil
	} else if strings.EqualFold("L_4", pL) {
		return uint8(entity.L_4), nil
	} else {
		return 0, fmt.Errorf("invalid permission level: %s", pL)
	}
}

func Data2View(data entity.Data) ViewData {
	time := general.Timestamp2Str(data.MeteData.TimeStamp)
	v := ViewData{
		ID:              data.ID,
		HashDigest:      data.HashDigest,
		Owner:           data.Owner,
		PermissionLevel: PL2String[data.PermissionLevel],
		Signature:       data.Signature,
		MeteData: MeteData{
			FileName:  data.MeteData.FileName,
			TimeStamp: time,
			Size:      data.MeteData.Size,
		},
		PermissionList: []PermissionList{},
	}
	for _, pl := range data.PermissionList {
		view_pl := PermissionList{
			PublicKey:  pl.PublicKey,
			Permission: Psermission2String(pl.Permission),
		}
		v.PermissionList = append(v.PermissionList, view_pl)
	}
	return v
}

func Behavior2View(b entity.Behavior) ViewBehavior {
	v := ViewBehavior{
		PublicKey:  b.PublicKey,
		PrivateKey: b.PrivateKey,
		Assets:     []Asset{},
		Logs:       []LogInfo{},
	}
	for _, a := range b.Assets {
		view_asset := Asset{
			FileID:     a.FileID,
			Permission: Psermission2String(a.Permission),
		}
		v.Assets = append(v.Assets, view_asset)
	}

	for _, l := range b.Logs {
		view_log := LogInfo{
			TimeStamp:  general.Timestamp2Str(l.TimeStamp),
			PublicKey:  l.PublicKey,
			Method:     l.Method,
			FileID:     l.FileID,
			HashDigest: l.HashDigest,
		}
		v.Logs = append(v.Logs, view_log)
	}
	return v
}
