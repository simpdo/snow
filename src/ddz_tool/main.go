// xmltool project main.go
package main

import (
	"ddz_tool/config"
	"fmt"
)

func main() {
	err := config.WriteRoundBox2Json("round_box.xlsx", "round_box.json")
	if err != nil {
		fmt.Println(err.Error())
	}
}
