package main

import "testing"

func TestGetFromGaode(t *testing.T) {
	var postdata GetData
	postdata.Keywords = "修车厂"
	postdata.Types = ""
	postdata.City = "青岛"
	postdata.Children = "1"
	postdata.Offset = "25"
	postdata.Page = "1"
	postdata.Extensions = "base"
	//	var data = postdata.GetFromGaode(gaodedata.KeySearchAPI)
	//	t.Log("data is ", data)
}
