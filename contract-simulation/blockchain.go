package contractsimulation

import (
	"database/sql"
	"dfa/entity"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

const (
	userName = "user0"
	password = "pwd-user0"
	ip       = "127.0.0.1"
	port     = "3306"
	dbName   = "blockchain_simulation"
)

type BehaviorContract interface {
	Register() (entity.Behavior, error)
	Login(publickey string, s entity.Signature) (entity.Behavior, error)
	FindBehavior(publickey string) (entity.Behavior, bool)
	RecordAccess(publickey string, log entity.LogInfo) (entity.Behavior, error)
	Authorize(hostkey, customerkey, fileid string, p entity.Permission) (bool, error)
}

type behaviorchain struct {
	DB *sql.DB
}

func NewBehaviorChain() BehaviorContract {
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	db, err := sql.Open("mysql", path)
	if err != nil {
		panic(err)
	}
	return &behaviorchain{
		DB: db,
	}
}
