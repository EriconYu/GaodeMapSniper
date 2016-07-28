package gaodedb

import (
	"Config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/Go-SQL-Driver/MySQL" // 龟腚 Lint风格要求加上注释
	"github.com/cihub/seelog"
)

// DBURL 数据库地址
// DBURL = "root:123456@tcp(localhost:3306)/"
// var DBURL = "root:123456@tcp(192.168.1.60:3306)/"
var DBURL = (*Config.Instance()).MySQLUserName + ":" + (*Config.Instance()).MySQLPassword + "@" + (*Config.Instance()).MySQLAddress + "/"

//DBURL = "webuser:1QAZ1qaz@tcp(rdsoo1h37t3e6b5kbk13.mysql.rds.aliyuncs.com:3306)/"

// DBMgr ...
type DBMgr struct {
	db *sql.DB
}

var dbmgr DBMgr

func init() {
	DBInit()
	seelog.Info("as")
}

// DBInit ...
func DBInit() {
	log.Println("DBInit start...")
	log.Println("DBInit DBURL is ", DBURL)
	var e error
	dbmgr.db, e = sql.Open("mysql", DBURL+Config.Instance().MySQLDataBase+"?charset=utf8")
	if e != nil {
		log.Fatal("mysql error:", e)
	}
	dbmgr.db.SetMaxOpenConns(10)
	// dbmgr.db.SetMaxIdleConns(4)
	err := dbmgr.db.Ping()
	if err != nil {
		log.Println("ping db err...", err)
	}
	log.Println("DBInit end...")
}

// GetDB ...
func GetDB() *sql.DB {
	return dbmgr.db
}

// CloseRows ...
func CloseRows(r *sql.Rows) {
	err := r.Close()
	if err != nil {
		fmt.Println("...!!!...\n\n\n...", err)
		fmt.Println("...Row close failer....")
		return
	}
	// fmt.Println("...Row close successfull....")
}

// ErrRows ...
func ErrRows(r *sql.Rows) {
	err := r.Err()
	if err != nil {
		fmt.Println("...!!!...\n\n\n...Row Next:", err)
		fmt.Println("...Row Next failer....")
		return
	}
	// fmt.Println("...Row Next successfull....")
}
