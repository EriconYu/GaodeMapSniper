package main

import (
	"Config"
	"fmt"
)

func init() {
	Config.Instance().GetConfigure()
	fmt.Println()
}

func main() {
	fmt.Println("start!")
	fmt.Println("Configure is ", *Config.Instance())
	var c CityList
	c.ReadFile()

	var citylist CityList
	city := citylist.ReadFile()
	citylist.ParseData(city)
	count := len(citylist.CityList)

	keycount := len(Config.Instance().KeyWord)

	for i := 0; i < keycount; i++ {
		fmt.Println("i is ", i, "keyword is ", Config.Instance().KeyWord[i].Key)
		for j := 0; j < count; j++ {
			cityname := citylist.CityList[j].Name
			Run(Config.Instance().KeyWord[i].Key, cityname)
		}
	}
	fmt.Println("end!")
}
