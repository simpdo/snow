// xml2json project main.go
package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	buff, err := ioutil.ReadFile("paytype.xml")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var pays XmlPayTypeManager
	err = xml.Unmarshal(buff, &pays)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	pay_manage := PayTypeManager{}
	app_manage := AppConfigManager{}

	for i, _ := range pays.Pays {
		app := AppConfig{}
		app.Type = pays.Pays[i].Type
		app.Apps = pays.Pays[i].Apps

		pay_manage.Pays = append(pay_manage.Pays, pays.Pays[i].PayType)
		app_manage.Configs = append(app_manage.Configs, app)
	}

	buff, err = json.Marshal(&pay_manage)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	ioutil.WriteFile("paytype.json", buff, os.ModePerm)

	buff, err = json.Marshal(&app_manage)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	ioutil.WriteFile("paytype_province.json", buff, os.ModePerm)
}
