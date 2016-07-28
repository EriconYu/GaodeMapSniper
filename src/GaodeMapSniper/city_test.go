package main

import (
	"fmt"
	"testing"
)

// TestReadFile ..
func TestReadFile(T *testing.T) {
	var c CityList
	filestream := c.ReadFile()

	fmt.Println(filestream)
}
