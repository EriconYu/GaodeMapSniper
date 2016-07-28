package gaodedb

import (
	"fmt"
	"testing"
)

func TestSave(T *testing.T) {
	var poidb Poidb
	poidb.ID = "hello"
	poidb.Location = "100000,e87239485"
	poidb.Name = "可其修车厂"
	//poidb.Save()
}

func TestIsExist(T *testing.T) {
	fmt.Println("..............................")
	var poidb Poidb
	poidb.ID = "B000A019C6"
	ok := poidb.IsExist()
	fmt.Println("ok is ", ok)
}
