package main

import (
	"Config"
	"encoding/json"
	"fmt"
	"os"
)

// City ..
type City struct {
	Name string `json:"name"`
}

// CityList ..
type CityList struct {
	CityList []City `json:"CityList"`
}

//ReadFile ..
func (c *CityList) ReadFile() []byte {
	fmt.Println("ReadFile start")
	f, e := os.Open(Config.Instance().CityPath)
	if e != nil {
		fmt.Println("ReadFile OpenFile err:", e)
		return []byte("")
	}
	defer f.Close()

	fileInfo, _ := f.Stat()
	filestream := make([]byte, fileInfo.Size(), fileInfo.Size())
	n, _ := f.Read(filestream)
	fmt.Println("ReadFile end , n is ", n)
	return filestream
}

// ParseData ..
func (c *CityList) ParseData(filestream []byte) {
	json.Unmarshal(filestream, c)
}
