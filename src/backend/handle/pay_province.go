package handle

import (
	"backend/config"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func ProvincePayHandle(w http.ResponseWriter, r *http.Request) {
	log.Println(r.RequestURI)

	err := r.ParseForm()
	if err != nil {
		w.Write([]byte("parse form failed!"))
		return
	}

	var resp []byte
	if r.Method == "GET" {
		resp = GetProvincePay(r)
	} else if r.Method == "POST" {
		resp = UpdateProvincePay(r)
	}

	log.Println(string(resp))
	w.Write(resp)
}

func hashProvincePay(arr []config.ProvincePay) map[string][]config.PayPriority {
	province2PayPriority := make(map[string][]config.PayPriority)
	for i := 0; i < len(arr); i++ {
		province2PayPriority[arr[i].Province] = arr[i].PayPrioritys
	}

	return province2PayPriority
}

func UpdateProvincePay(r *http.Request) []byte {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return []byte(err.Error())
	}

	log.Println(string(body))

	param := config.ProvincePayPriority{}
	err = json.Unmarshal(body, &param)
	if err != nil {
		return []byte(err.Error())
	}

	log.Println(param)
	province_pay := hashProvincePay(param.Config)
	err = config.UpdateProvincePay(param.Channel, param.Isp, province_pay)
	if err != nil {
		return []byte(err.Error())
	}

	return []byte("0")
}

func GetProvincePay(r *http.Request) []byte {
	channel, ok := r.Form["channel"]
	if !ok {
		return []byte("channel is null")

	}

	isps, ok := r.Form["isp"]
	if !ok {
		return []byte("isp is null")
	}

	isp, err := strconv.Atoi(isps[0])
	if err != nil {
		return []byte("isp data is wrong")
	}

	data, err := config.GetProvincePay(channel[0], isp)
	if err != nil {
		log.Println(err.Error())
		return []byte(err.Error())
	}

	return data
}
