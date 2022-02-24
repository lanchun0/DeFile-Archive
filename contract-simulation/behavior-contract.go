package contractsimulation

import (
	"dfa/entity"
	"dfa/general"
	"fmt"
)

func (contract *behaviorchain) Register() (entity.Behavior, error) {
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
		TimeStamp: general.GenerateTimeStamp(),
		PublicKey: pubStr,
		Method:    "register",
	}
	b.Logs = append(b.Logs, l)

	ok, err := insertBehavior(contract, b)
	if err != nil {
		return entity.Behavior{}, fmt.Errorf("behaviorchain: register: %w", err)
	}
	if !ok {
		return entity.Behavior{}, fmt.Errorf("behaviorchain: register falied")
	}
	if ok {
		fmt.Println("Success register!")
	}
	return b, nil

}

func (contract *behaviorchain) Login(publickey string, s entity.Signature) (entity.Behavior, error) {

	bh, exist := selectBehavior(contract, publickey)
	if !exist {
		return entity.Behavior{}, fmt.Errorf("behavior chain: login: user doesn't exist")
	}

	priByte, err := general.StrToBytes(bh.PrivateKey)
	pubByte, err := general.StrToBytes(bh.PublicKey)
	if err != nil {
		return entity.Behavior{}, fmt.Errorf("error: %s", "missing keys")
	}
	valid, err := general.VerifySignature(priByte, s, pubByte)
	if !valid || err != nil {
		return entity.Behavior{}, fmt.Errorf("behavior chain: login: faked signature or %w", err)
	}
	bh.PrivateKey = ""
	return bh, nil
}

func (contract *behaviorchain) FindBehavior(publickey string) (entity.Behavior, bool) {
	return selectBehavior(*&contract, publickey)
}

func (contract *behaviorchain) RecordAccess(publickey string, log entity.LogInfo) (entity.Behavior, error) {
	behavior, exist := contract.FindBehavior(publickey)
	if !exist {
		return behavior, fmt.Errorf("behaviorchain: record access: cannot find user %s on Behavior Blockchain", publickey)
	}
	condition := make(map[string]entity.Permission)
	condition["download"], condition["modify"], condition["create"] = entity.Reader, entity.Writer, entity.Owner

	// Find the asset and make sure the permission is accessible
	asset := entity.Asset{
		FileID: log.FileID,
	}
	asset, exist = selectAsset(contract, publickey, asset)
	if !exist {
		return behavior, fmt.Errorf("behaviorchain: record access: cannot find asset %s on Behavior Blockchain", log.FileID)
	}
	if asset.Permission < uint(condition[log.Method]) {
		return behavior, fmt.Errorf("behaviorchain: record access: permission denied")
	}
	// update the logs
	as := entity.Asset{
		FileID:     log.FileID,
		Permission: uint(condition[log.Method]) | asset.Permission,
	}
	ok, _ := updateAsset(contract, publickey, as)
	if !ok {
		return behavior, fmt.Errorf("behaviorchain: record access: fail updating asset")
	}

	behavior.Logs = append(behavior.Logs, log)
	return behavior, nil
}

func (contract *behaviorchain) Authorize(hostkey, customerkey, fileid string, p entity.Permission) (bool, error) {
	// make sure host and customer exist
	_, existB := contract.FindBehavior(hostkey)
	if !existB {
		return false, fmt.Errorf("behavior chain: authorize: user %s does not exist", hostkey)
	}
	_, existC := contract.FindBehavior(customerkey)
	if !existC {
		return false, fmt.Errorf("behavior chain: authorize: user %s does not exist", customerkey)
	}

	// make sure customer doesn't have the file or have insufficient permission
	asset := entity.Asset{
		FileID: fileid,
	}
	a, exist := selectAsset(contract, customerkey, asset)
	if !exist {
		asset = entity.Asset{
			FileID:     fileid,
			Permission: uint(p),
		}
		return insertAsset(contract, customerkey, asset)
	} else if exist && a.Permission < uint(p) {
		asset = entity.Asset{
			FileID:     fileid,
			Permission: uint(p) | a.Permission,
		}
		return updateAsset(contract, customerkey, asset)
	} else {
		return false, fmt.Errorf("behavior chain: authorize: user %s already has the permission", customerkey)
	}

}

func insertBehavior(contract *behaviorchain, b entity.Behavior) (bool, error) {
	tx, err := contract.DB.Begin()
	if err != nil {
		return false, fmt.Errorf("behaviorchain: insert: %w", err)
	}

	stmt, err := tx.Prepare("INSERT INTO behavior_table (`publickey`, `privatekey`) VALUES (?,?)")
	if err != nil {
		return false, fmt.Errorf("behaviorchain: insert: prepare fail %w", err)
	}
	_, err = stmt.Exec(b.PublicKey, b.PrivateKey)
	if err != nil {
		return false, fmt.Errorf("behaviorchain: insert: exec fail %w", err)
	}
	// //
	// tx.Commit()
	// id, _ := res.LastInsertId()
	// fmt.Println(id)
	// //
	return true, nil
}

func selectBehavior(contract *behaviorchain, pub string) (entity.Behavior, bool) {
	b := entity.Behavior{
		Logs:   []entity.LogInfo{},
		Assets: []entity.Asset{},
	}

	err := contract.DB.QueryRow("SELECT * from behavior_table WHERE publickey = ?", pub).Scan(&b.PublicKey, &b.PrivateKey)
	if err != nil {
		return entity.Behavior{}, false
	}

	rows, err := contract.DB.Query("SELECT * from loginfo_table WHERE behavior_publickey = ?", pub)
	if err != nil {
		fmt.Printf("select behavior: %s", err)
	}
	for rows.Next() {
		var ignore string
		var l entity.LogInfo
		var timeString string
		err := rows.Scan(&ignore, &timeString, &l.PublicKey, &l.Method, &l.FileID, l.HashDigest)
		if err != nil {
			fmt.Println(err)
		}
		l.TimeStamp, _ = general.Str2Timestamp(timeString)
		b.Logs = append(b.Logs, l)
	}

	rows, _ = contract.DB.Query("SELECT * from asset_table WHERE behavior_publickey = ?", pub)
	for rows.Next() {
		var ignore string
		var a entity.Asset
		err := rows.Scan(&ignore, &a.FileID, &a.Permission)
		if err != nil {
			fmt.Println(err)
		}
		b.Assets = append(b.Assets, a)
	}
	return b, true
}

func selectAsset(contract *behaviorchain, pub string, a entity.Asset) (entity.Asset, bool) {
	row := contract.DB.QueryRow("SELECT * from behavior_table WHERE behavior_publickey = ? AND fileid = ?", pub, a.FileID)
	var ignore string
	var asset entity.Asset
	err := row.Scan(&ignore, &asset.FileID, &asset.Permission)
	if err != nil {
		return asset, false
	}
	return asset, true
}

func insertLoginfo(contract *behaviorchain, pub string, log entity.LogInfo) (bool, error) {
	tx, err := contract.DB.Begin()
	if err != nil {
		return false, fmt.Errorf("behavior chain: update loginfo: %w", err)
	}
	stmt, err := tx.Prepare("INSERT INTO loginfo_table (`behavior_publickey`,`timestamp`,`publickey`,`method`,`fileid`,`hashdigest`) VALUES (?,?,?,?,?,?)")
	if err != nil {
		return false, fmt.Errorf("behavior chain: update loginfo: %w", err)
	}

	_, err = stmt.Exec(pub, general.Timestamp2Str(log.TimeStamp), log.PublicKey, log.Method, log.FileID, log.HashDigest)
	if err != nil {
		return false, fmt.Errorf("behavior chain: update loginfo: %w", err)
	}
	tx.Commit()
	return true, nil

}

func insertAsset(contract *behaviorchain, pub string, a entity.Asset) (bool, error) {
	tx, err := contract.DB.Begin()
	if err != nil {
		return false, fmt.Errorf("behavior chain: insert asset: %w", err)
	}
	stmt, err := tx.Prepare("INSERT INTO asset_table (`behavior_publickey`,`fileid`,`permission`) VALUES (?,?,?)")
	if err != nil {
		return false, fmt.Errorf("behavior chain: insert asset: %w", err)
	}

	_, err = stmt.Exec(pub, a.FileID, a.Permission)
	if err != nil {
		return false, fmt.Errorf("behavior chain: insert asset: %w", err)
	}
	tx.Commit()
	return true, nil
}

func updateAsset(contract *behaviorchain, pub string, a entity.Asset) (bool, error) {
	tx, err := contract.DB.Begin()
	if err != nil {
		return false, fmt.Errorf("behavior chain: update asset: %w", err)
	}
	stmt, err := tx.Prepare("UPDATE asset_table SET permission = ? WHERE behavior_publickey = ? AND fileid = ?")
	if err != nil {
		return false, fmt.Errorf("behavior chain: update asset: %w", err)
	}
	_, err = stmt.Exec(pub, a.FileID, a.Permission)
	if err != nil {
		return false, fmt.Errorf("behavior chain: update asset: %w", err)
	}
	tx.Commit()
	return true, nil

}
