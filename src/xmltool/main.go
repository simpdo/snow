// xmltool project main.go
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/tealeg/xlsx"
	"io/ioutil"
	"net/http"
	"os"
)

type ProvincePayPriorityConfig struct {
	Config []ProvincePayPriority `json:"root"`
}

type ProvincePayPriority struct {
	Channel string        `json:"channel"`
	Isp     int           `json:"isp"`
	Config  []ProvincePay `json:"config"`
}

type ProvincePay struct {
	Province     string        `json:"province"`
	PayPrioritys []PayPriority `json:"pay"`
}

type PayPriority struct {
	Pay      int `json:"type"`
	Priority int `json:"priority"`
}

var row_base int = 2
var column_base int = 1

func main() {
	flag.Args()
	file, err := xlsx.OpenFile("province_pay_priority.xlsx")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	config := ProvincePayPriorityConfig{}

	for _, sheet := range file.Sheets {
		item := ProvincePayPriority{}
		item.Channel = "02, 03"
		if sheet.Name == "移动" {
			item.Isp = 1
		} else if sheet.Name == "联通" {
			item.Isp = 2
		} else if sheet.Name == "电信" {
			item.Isp = 3
		} else {
			break
		}

		var paytypes []int

		for i := row_base; i < len(sheet.Rows); i++ {
			row := sheet.Rows[i]
			var province_pay ProvincePay = ProvincePay{}

			if i == row_base { //获取支付方式
				for j := column_base + 1; j < len(row.Cells); j++ {
					val, _ := row.Cells[j].Int()
					paytypes = append(paytypes, val)
				}
			} else {
				for j := column_base; j < len(row.Cells); j++ {

					if j == column_base {
						province_pay.Province = row.Cells[j].String()
					} else {
						val, _ := row.Cells[j].Int()
						var pay_priority PayPriority
						pay_priority.Pay = paytypes[j-column_base-1]
						pay_priority.Priority = val
						province_pay.PayPrioritys = append(province_pay.PayPrioritys, pay_priority)
					}
				}
				item.Config = append(item.Config, province_pay)
			}
		}
		config.Config = append(config.Config, item)
	}

	data, _ := json.Marshal(config)
	fmt.Println(string(data))
	ioutil.WriteFile("province_pay.json", data, os.ModePerm)
}
