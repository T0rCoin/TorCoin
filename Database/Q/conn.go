package Db

import (
	"TorCoin/conf"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (

	Dsn string
	Host, Port, Username, Password, Dbname string

)

func Conn(NodeName string) (*gorm.DB,error) {
	var err error
	Cfg := conf.Load()
	Host, err = Cfg.GetValue(NodeName,"Host")
	Port, err = Cfg.GetValue(NodeName,"Port")
	Username, err = Cfg.GetValue(NodeName,"Username")
	Password, err = Cfg.GetValue(NodeName,"Password")
	Dbname, err = Cfg.GetValue(NodeName,"Dbname")
	if err != nil {
		return nil, err
	}
	Dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", Username, Password, Host, Port, Dbname)
	db, err := gorm.Open(mysql.Open(Dsn),&gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
