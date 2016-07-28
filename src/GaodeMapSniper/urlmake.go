package main

import (
	"gaodedata"
)

// URI ..
type URI struct {
	key        string
	keywords   string
	types      string
	city       string
	citylimit  string
	children   string
	offset     string
	page       string
	building   string
	floor      string
	extensions string
	sig        string
	output     string
	callback   string
}

// CreateURI ..
func (u *URI) CreateURI(keywords string, city string, page string, extensions string) string {
	u.key = gaodedata.Keqicarweb
	u.keywords = keywords
	u.city = city
	u.children = "1"
	u.offset = "25"
	u.page = page
	u.extensions = extensions
	var uri = "key=" + u.key + "&keywords=" + u.keywords + "&types" + u.types + "&city=" + u.city + "&children=" + u.children + "&offset=" + u.offset + "&page=" + u.page + "&extensions=" + u.extensions
	return uri
}
