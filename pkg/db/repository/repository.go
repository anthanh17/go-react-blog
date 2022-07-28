package repository

import (
	"github.com/anthanh17/go-react-blog/pkg/db"
)

func SetupDb() (err error) {
	//MySqlDbConnection, err = db.GetMySqlDatabase()
	_, _ = db.GetMySqlDatabase()
	return
}
