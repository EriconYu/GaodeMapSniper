package Config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// Configure ..
var configure *Config

type keyword struct {
	Key string `json:"key"`
}

//Config 系统配置选项
type Config struct {
	//	KeyWord        string `json:"keyword"`
	KeyWord        []keyword `json:"keyword"`
	CityPath       string    `json:"citypath"`
	MySQLAddress   string    `json:"mysqladdress"`
	MySQLUserName  string    `json:"mysqlusername"`
	MySQLPassword  string    `json:"mysqlpassword"`
	MySQLDataBase  string    `json:"mysqldatabase"`
	MySQLDataSchem string    `json:"mysqldataschem"`
	OpenKey        string    `json:"openkey"`
}

//Instance 单例模式
func Instance() *Config {
	if configure == nil {
		configure = &Config{
			CityPath:       "",
			MySQLUserName:  "",
			MySQLPassword:  "",
			MySQLDataBase:  "",
			MySQLDataSchem: "",
			OpenKey:        "",
		}
		configure.GetConfigure()
	}
	return configure
}

//GetConfigure 获取配置项
func (c *Config) GetConfigure() {
	f, e := os.Open("./config")
	if e != nil {
		fmt.Println("ReadFile OpenFile err:", e)
		return
	}
	defer f.Close()
	fileInfo, _ := f.Stat()
	filestream := make([]byte, fileInfo.Size(), fileInfo.Size())
	f.Read(filestream)
	e = json.Unmarshal(filestream, &configure)
	if e != nil {
		log.Fatal("init json Umarshal e:", e)
		return
	}
}
