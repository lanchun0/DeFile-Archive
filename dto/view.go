package dto

import (
	"dfa/entity"

	"fmt"
	"strings"
)

type ViewData struct {
	ID              string `json:"id"`
	HashDigest      string `json:"hashdigest"`
	Owner           string `json:"owner"`
	PermissionLevel string `json:"permissionlevel"`
	Tradable        bool   `json:"tradable"`
	MeteData        entity.MeteData
	PermissionList  []PermissionList
}

type PermissionList struct {
	Address    string `json:"address"`
	Permission string `json:"permission"`
}

type ViewUser struct {
	Address string `json:"address"`
	Balance uint64 `json:"balance"`
	// PrivateKey string  `json:"privatekey"`
	Assets []Asset `json:"assets"`
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
		str = str + " [owner]"
	}
	if peimission&uint8(entity.Administrator) != 0 {
		str = str + " [administrator]"
	}
	if peimission&uint8(entity.Writer) != 0 {
		str = str + " [writer]"
	}
	if peimission&uint8(entity.Reader) != 0 {
		str = str + " [reader]"
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
	v := ViewData{
		ID:              data.ID,
		Owner:           data.Owner,
		PermissionLevel: PL2String[data.PermissionLevel],
		Tradable:        data.Tradable,
		MeteData:        data.MeteData,
		PermissionList:  []PermissionList{},
	}
	for _, pl := range data.PermissionList {
		view_pl := PermissionList{
			Address:    pl.Address,
			Permission: Psermission2String(pl.Permission),
		}
		v.PermissionList = append(v.PermissionList, view_pl)
	}
	return v
}

func Behavior2View(u entity.User) ViewUser {
	v := ViewUser{
		Address: u.Address,
		Balance: u.Balance,
		Assets:  []Asset{},
	}
	for _, a := range u.Assets {
		view_asset := Asset{
			FileID:     a.FileID,
			Permission: Psermission2String(a.Permission),
		}
		v.Assets = append(v.Assets, view_asset)
	}
	return v
}
