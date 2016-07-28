package gaodedata

import "encoding/json"

//GaodeData 高德地图返回的数据
type GaodeData struct {
	Status   string `json:"status"`
	Count    string `json:"count"`
	Info     string `json:"info"`
	Infocode string `json:"infocode"`
	// Suggestion string    `json:"suggestion"`
	Pois []PoiData `json:"pois"`
}

//Photo 商家图片
type Photo struct {
	URL string `json:"url"`
}

//PoiData Poi数据
type PoiData struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	TypeCode string `json:"typecode"`
	// BizType  string `json:"biz_type"`
	Address  string `json:"address"`
	Location string `json:"location"`
	Tel      string `json:"tel"`
	// Distance string `json:"distance"`
	// BizExt   string `json:"biz_ext"`
	PName    string `json:"pname"`
	CityName string `json:"cityname"`
	AdName   string `json:"adname"`
	// ShopID   string `json:"shopid"`
	Photos []Photo `json:"photos"`
}

// ParseData 解析从高德回来的数据
func (g *GaodeData) ParseData(data string) {
	json.Unmarshal([]byte(data), g)
}
