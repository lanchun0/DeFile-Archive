package contractsimulation

import (
	"dfa/entity"
	"dfa/general"
	"fmt"
)

type DataContract interface {
}

func (contract *datachain) FindAllData() ([]entity.Data, int) {
	res, count := []entity.Data{}, 0
	// rows, err := contract.DB.Query("SELECT * from data_table")

	return res, count
}

func selectAllData(contract *datachain) {

}

// exlude hashdigest
func selectHistories(contract *datachain, id, hashdigest string) ([]entity.DataHistory, bool) {
	res := []entity.DataHistory{}

	q := "SELECT history_table.hash_digest, history_table.signer, history_table.signature, "
	q = q + "metedata_table.filename, metedata_table.size, metedata_table.timestamp "
	q = q + "FROM history_table, metedata_table "
	q = q + "WHERE history_table.id = ? AND history_table.hash_digest = ? "
	q = q + "AND (history_table.id = metedata_table.id AND history_table.hash_digest = metedata_table.hash_digest) "

	rows, err := contract.DB.Query(q, id, hashdigest)
	if err != nil {
		fmt.Printf("select histories: %s", err)
		return res, false
	}
	for rows.Next() {
		var h entity.DataHistory
		var s, t string
		rows.Scan(&h.HashDigest, &h.Signer, &s, &h.MeteData.FileName, &h.MeteData.Size, &t)
		h.Signature = general.ParseSignature(s)
		h.MeteData.TimeStamp, _ = general.Str2Timestamp(t)
		res = append(res, h)
	}
	return res, true
}
