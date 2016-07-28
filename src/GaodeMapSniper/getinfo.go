package main

import (
	"Config"
	"encoding/json"
	"fmt"
	"gaodedata"
	"gaodedb"
	"strconv"
	"time"

	"github.com/astaxie/beego/httplib"
)

//GetData 发送给高德地图的数据
type GetData struct {
	Keywords   string `json:"keywords"`
	Types      string `json:"types"`
	City       string `json:"city"`
	Children   string `json:"children"`
	Offset     string `json:"offset"`
	Page       string `json:"page"`
	Extensions string `json:"extensions"`
}

//GetFromGaode Post数据给高德
func (g *GetData) GetFromGaode(url string) string {
	fmt.Println("GetFromGaode Start")
	req := httplib.Get(url)
	req.Param("keywords", g.Keywords)
	req.Param("types", g.Types)
	req.Param("city", g.City)
	req.Param("children", g.Children)
	req.Param("offset", g.Offset)
	req.Param("page", g.Page)
	req.Param("extensions", g.Extensions)
	req.Param("key", Config.Instance().OpenKey)
	str, err := req.String()
	if err != nil {
		// error
		fmt.Println("GetFromGaode PostToGaode err is ", err)
	}
	fmt.Println("GetFromGaode End")
	return str
}

// Run 跑起来吧
func Run(keywords string, city string) {
	var postdata GetData
	postdata.Keywords = keywords
	postdata.Types = ""
	postdata.City = city
	postdata.Children = "1"
	postdata.Offset = "25"
	postdata.Page = "1"
	// postdata.Extensions = "base"
	postdata.Extensions = "all"

	for i := 1; ; i++ {

		postdata.Page = strconv.Itoa(i)

		data := postdata.GetFromGaode(gaodedata.KeySearchAPI)

		var gaodeData gaodedata.GaodeData
		gaodeData.ParseData(data)

		if gaodeData.Status == "1" && gaodeData.Count != "0" {
			for i := 0; i < len(gaodeData.Pois); i++ {
				fmt.Println("[ time is ", time.Now(), "]")
				fmt.Println("i is ", i, "Name is ", city, ":", gaodeData.Pois[i].Name, "ID is", gaodeData.Pois[i].ID, "img is", gaodeData.Pois[i].Photos)
				var poidb gaodedb.Poidb
				poidb.ID = gaodeData.Pois[i].ID
				poidb.Location = gaodeData.Pois[i].Location
				poidb.Name = city + ":" + gaodeData.Pois[i].Name
				if len(gaodeData.Pois[i].Photos) != 0 {
					ImgB, _ := json.Marshal(gaodeData.Pois[i].Photos)
					poidb.Img = string(ImgB)
				}
				poidb.Save()
			}
		} else if gaodeData.Count == "0" {
			fmt.Println("Run Err,Count is 0", keywords, "-", city, "page is ", postdata.Page)
			return
		} else {
			fmt.Println("Run Err,Status is 0", keywords, "-", city, "page is ", postdata.Page)
		}
	}
}
