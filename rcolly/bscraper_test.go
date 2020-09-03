package rcolly

import (
	"fmt"
	"testing"
)

func TestScrape(t *testing.T) {
	// Scrap()
}

func TestVue(t *testing.T) {
	VScrap("")
}

func Test_returnN(t *testing.T) {
	fmt.Println(returnN())
	return
}

func returnN() (result int) {
	defer func() {
		if p := recover(); p != nil {
			result = p.(int)
		}
	}()
	panic(3)
}
