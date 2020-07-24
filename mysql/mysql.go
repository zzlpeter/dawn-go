package mysql

import (
	"github.com/jinzhu/gorm"
	"github.com/zzlpeter/dawn-go/libs"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"log"
	"sync"
)

var dbMap = map[string]*gorm.DB{}
var once sync.Once

func getDbConn() {
	once.Do(func() {
		makeDbConn()
	})
}

func makeDbConn() {
	dbConf := libs.Config{}.MysqlConfS()
	for alias, conf := range dbConf {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%ds",
			conf["username"], conf["password"], conf["host"],
			conf["port"], conf["db"], conf["timeout"])
		db, err := gorm.Open("mysql", dsn)
		if err != nil {
			log.Fatalf("connect mysql: %v err: %v", alias, err.Error())
		}
		maxConn := conf["max_con"].(int)
		db.DB().SetMaxOpenConns(maxConn)
		maxIdle := conf["max_idle"].(int)
		db.DB().SetMaxIdleConns(maxIdle)
		dbMap[alias] = db
	}
}

func GetDbConn(db string) *gorm.DB {
	getDbConn()
	return dbMap[db]
}