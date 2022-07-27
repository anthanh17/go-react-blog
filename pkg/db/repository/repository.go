package repository

import (
	"github.com/anthanh17/go-react-blog/pkg/db"
)

var MySqlDbConnection *mysql.DB

func SetupDb() (err error) {
	MySqlDbConnection, err = db.GetMySqlDatabase()
	return
}
