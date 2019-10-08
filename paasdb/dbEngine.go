package paasdb

import (
	"errors"
	"strconv"

	"github.com/connext-cs/pub/config"
	"github.com/connext-cs/pub/logs"

	"github.com/go-xorm/xorm"
)

var orm *xorm.Engine

func CloudprojectEngine() *xorm.Engine {
	orm, err := GetEngine()
	if err != nil {
		logs.Error(err)
	}
	return orm
}

func GetEngine() (*xorm.Engine, error) {
	if orm == nil {
		var err error
		orm, err = mysqlEngine()
		if err != nil {
			logs.Error(err.Error())
			return nil, err
		}
	}
	if orm == nil {
		err := errors.New("database init error")
		if err != nil {
			logs.Error(err.Error())
			return nil, err
		}
	}
	orm.ShowSQL()
	return orm, nil
}

func mysqlEngine() (*xorm.Engine, error) {
	Host := config.CMysqlHost()
	Port := uint16(config.CMysqlPort())
	Name := config.CCloudMysqlDatabase()
	User := config.CMysqlUserName()
	Password := config.CMysqlPasswd()
	dburl := User + ":" + Password + "@tcp(" + Host + ":" + strconv.Itoa(int(Port)) + ")/" + Name + "?charset=utf8"
	logs.Info("dburl:", dburl)
	return xorm.NewEngine("mysql", dburl)
	// return xorm.NewEngine("mysql", "connextpaas:connext@0101@tcp(127.0.0.1:3306)/vmwareproject?charset=utf8")
}
