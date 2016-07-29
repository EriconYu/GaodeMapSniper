package gaodedb

import (
	"Config"
	"database/sql"
	"fmt"
)

// Poidb 存储到数据库中的结构体
type Poidb struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Img      string `json:"img"`
}

// IsExist :查重
func (p *Poidb) IsExist() (ok bool) {
	//查重，存在返回true，不存在返回false
	db := GetDB()
	var id string
	var cmd = fmt.Sprintf(`select id from `+Config.Instance().MySQLDataSchem+` where id='%s';`, p.ID)
	fmt.Println("cmd is ", cmd)

	err := db.QueryRow(fmt.Sprintf(`select id from `+Config.Instance().MySQLDataSchem+` where id='%s';`, p.ID)).Scan(&id)
	fmt.Println("id is ", id)
	if err != nil {
		fmt.Println("IsExist err is ", err)
	}
	if id == "" {
		return false
	}
	return true
}

// Save ...
func (p *Poidb) Save() {
	var e error
	var stmt *sql.Stmt
	db := GetDB()

	if IsExist := p.IsExist(); IsExist == true { //已存在
		//更新原有的记录
		stmt, e = db.Prepare("update " + Config.Instance().MySQLDataSchem + " set name=?,location=?,img=? where id=?")
		defer stmt.Close()
		if e != nil {
			fmt.Println("Poidb Save db.Prepare e :", e)
			return
		}
		_, e = stmt.Exec(p.Name, p.Location, p.Img, p.ID)
		if e != nil {
			fmt.Println("Poidb Save stmt.Exec e :", e)
		}
		return
	}

	//插入一条新纪录
	stmt, e = db.Prepare("insert into " + Config.Instance().MySQLDataSchem + "(id,name,location,img)values(?,?,?,?)")
	defer stmt.Close()
	if e != nil {
		fmt.Println("Piodb Save Prepare error:", e)
		return
	}
	_, e = stmt.Exec(p.ID, p.Name, p.Location, p.Img)
	if e != nil {
		fmt.Println("Piodb Save Exec error:", e, ",ID is", p.ID)

	}
}
