package paasdb

import (
	"fmt"
	"time"

	"github.com/connext-cs/pub/config"
	logs "github.com/connext-cs/pub/logs"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	// DB is global db pool
	db       *sqlx.DB
	database *config.Database
)

//func DbConfigSet(outdb *config.Database) {
//	database = new(config.Database)
//	database.Host = outdb.Host
//	database.Port = outdb.Port
//	database.User = outdb.User
//	database.Password = outdb.Password
//	database.Name = outdb.Name
//	fmt.Println("paasdb DbConfigSet, database:",database)
//}
//
//func init() {
//	database = new(config.Database)
//	database.Host = config.CMysqlHost()
//	database.Port = uint16(config.CMysqlPort())
//	database.Name = config.CCloudMysqlDatabase()
//	database.User = config.CMysqlUserName()
//	database.Password = config.CMysqlPasswd()
//	fmt.Println("paasdb init database:", database)
//}

// InitDB init global db loc=Local&
func InitDB() (err error) {
	if db == nil {
		dbURI := fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?parseTime=true&charset=utf8",
			database.User,
			database.Password,
			database.Host,
			database.Port,
			database.Name,
		)
		logs.Info("dbURI:", dbURI)
		db, err = sqlx.Open("mysql", dbURI)
		fmt.Println("db:", db)
		if err != nil {
			fmt.Println("Create DB error: ", err.Error())
			return err
		}
		if database.MaxIdleTime > 0 {
			db.SetConnMaxLifetime(time.Millisecond * time.Duration(database.MaxIdleTime))
		}
		if database.MaxIdle > 0 {
			db.SetMaxIdleConns(database.MaxIdle)
			db.SetMaxOpenConns(database.MaxOverflow + database.MaxIdle)
		} else {
			db.SetMaxIdleConns(10)
			db.SetMaxOpenConns(database.MaxOverflow + 10)
		}

		err = db.Ping()
		if err != nil {
			return
		}
	}
	return
}

func GetDB() *sqlx.DB {
	fmt.Printf("start init db")
	err := InitDB()
	if err != nil {
		fmt.Printf("init db error: %v", err)
	}

	return db
}
