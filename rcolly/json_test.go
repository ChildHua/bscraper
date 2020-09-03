package rcolly

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestTransmitJson(t *testing.T) {
	TransmitJson()
}

func TransmitJson() {
	// a := []int{1,2,3,4,5}
	b := map[string]string{
		"1": "1",
		"2": "1",
		"3": "1",
		"4": "1",
	}
	tb, _ := json.Marshal(b)
	fmt.Println(string(tb))
}
