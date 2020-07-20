package mysql

import (
	"database/sql"
	"github.com/zzlpeter/dawn-go/libs"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"log"
	"sync"
)


var dbMap map[string]*sql.DB
var once sync.Once

func getDbConn() {
	once.Do(func() {
		makeDbConn()
	})
}

func makeDbConn() {
	dbConf := libs.Config{}.MysqlConfS()
	for alias, conf := range dbConf {
		dataSourceName := fmt.Sprintf("%s:@tcp(%s:%v)/%s?charset=%s", conf["username"], conf["host"], conf["port"], conf["db"], conf["charset"])
		db, err := sql.Open("mysql", dataSourceName)
		if err != nil {
			log.Fatalf("connect mysql: %v err: %v", alias, err.Error())
		}
		maxConn := conf["max_conn"].(int)
		db.SetMaxOpenConns(maxConn)
		maxIdle := conf["max_idle"].(int)
		db.SetMaxIdleConns(maxIdle)
		dbMap[alias] = db
	}
}

func GetDbConn(db string) *sql.DB {
	getDbConn()
	return dbMap[db]
}