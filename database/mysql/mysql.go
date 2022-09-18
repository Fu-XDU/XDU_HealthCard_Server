package mysql

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"xdu-health-card/model"
)

var db *gorm.DB

func Connect(IP string, port int, username string, passwd string, dbname string) (Db *gorm.DB, err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, passwd, IP, port, dbname)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		log.Errorf("Connect mysql '%s'@'%s:%d' failed", username, IP, port)
		return
	}

	log.Infof("Successfully connected to mysql '%s'@'%s:%d'", username, IP, port)

	err = initMysql()
	if err != nil {
		log.Errorf("Initialize mysql database failed", username, IP, port)
		return
	}

	log.Infof("Successfully initialize mysql")
	Db = db
	return
}

func initMysql() (err error) {
	err = db.AutoMigrate(&model.HealthCard{}, &model.ThreeCheck{})
	return
}
