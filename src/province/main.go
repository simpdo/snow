// province project main.go
package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"gopkg.in/redis.v3"
	"net/http"
	"strconv"
	"time"
)

var addr *string
var ssdb *redis.Client

var province2Host map[string]string = map[string]string{
	"01": "223.203.192.1", "02": "103.231.144.2", "03": "101.21.128.3",
	"04": "36.192.28.4", "05": "1.24.64.5", "06": "36.192.51.6",
	"07": "36.192.76.7", "08": "1.188.144.8", "09": "42.159.128.9",
	"10": "43.248.148.10", "11": "39.180.128.11", "12": "36.32.192.12",
	"13": "36.37.36.13", "14": "39.166.216.14", "15": "27.215.128.15",
	"16": "1.198.252.16", "17": "49.210.128.17", "18": "43.224.160.18",
	"19": "61.146.116.19", "20": "222.216.196.20", "21": "124.225.112.21",
	"22": "110.187.192.22", "23": "125.218.190.23", "24": "42.243.224.24",
	"25": "218.206.180.25", "26": "14.197.128.26", "27": "27.225.240.27",
	"28": "61.133.237.28", "29": "36.193.128.29", "30": "49.115.112.30",
	"31": "222.198.179.31",
}

var province2Imsi map[string]string = map[string]string{
	"01": "89860042011494705468", "02": "89860042021494705468",
	"03": "89860042031494705468", "04": "89860042041494705468",
	"05": "89860042051494705468", "06": "89860042061494705468",
	"07": "89860042071494705468", "08": "89860042081494705468",
	"09": "89860042091494705468", "10": "89860042101494705468",
	"11": "89860042111494705468", "12": "89860042121494705468",
	"13": "89860042131494705468", "14": "89860042141494705468",
	"15": "89860042151494705468", "16": "89860042161494705468",
	"17": "89860042171494705468", "18": "89860042181494705468",
	"19": "89860042191494705468", "20": "89860042201494705468",
	"21": "89860042211494705468", "22": "89860042221494705468",
	"23": "89860042231494705468", "24": "89860042241494705468",
	"25": "89860042251494705468", "26": "89860042261494705468",
	"27": "89860042271494705468", "28": "89860042281494705468",
	"29": "89860042291494705468", "30": "89860042301494705468",
	"31": "89860042311494705468",
}

type User struct {
	Channel  string `json:"channel"`
	CpUid    string `json:"cp_uid"`
	Host     string `json:"host"`
	IMEI     string `json:"imei"`
	IMSI     string `json:"imsi"`
	Mac      string `json:"mac"`
	Platform int    `json:"platform"`
	Sim      string `json:"sim"`
	Isp      int    `json:"type"`
	Version  int    `json:"version"`
}

func SsdbClient() *redis.Client {
	if nil == ssdb {
		ssdb = redis.NewClient(&redis.Options{
			Addr:        *addr,
			MaxRetries:  3,
			ReadTimeout: time.Microsecond * 300,
			PoolSize:    100,
			PoolTimeout: time.Millisecond * 100,
		})
	}

	return ssdb
}

func ChangeUserHost(uid string, province string, isp string) (string, error) {
	var strResp string = string("")
	host, ok := province2Host[province]
	if !ok {
		msg := fmt.Sprintf("%s is not a invalid province code!", province)
		return strResp, errors.New(msg)
	}

	imsi, ok := province2Imsi[province]
	if !ok {
		msg := fmt.Sprintf("%s is not a invalid province code!", province)
		return strResp, errors.New(msg)
	}

	client := SsdbClient()
	if nil == client {
		return strResp, errors.New("ssdb client is invalid!")
	}

	key := fmt.Sprintf("client_%s", uid)
	val, err := client.Get(key).Bytes()
	if err != nil {
		msg := fmt.Sprintf("get: %v", err.Error())
		return strResp, errors.New(msg)
	}

	user := User{}
	err = json.Unmarshal(val, &user)
	if err != nil {
		msg := fmt.Sprintf("Unmarshal: %v", err.Error())
		return strResp, errors.New(msg)
	}

	strResp = fmt.Sprintf("%s", string(val))

	user.Host = host
	user.Sim = imsi
	user.Isp, _ = strconv.Atoi(isp)
	val, err = json.Marshal(&user)
	if err != nil {
		msg := fmt.Sprintf("Marshal: %v", err.Error())
		return strResp, errors.New(msg)
	}

	client.Set(key, val, 0)
	strResp = fmt.Sprintf("%s\n%s", strResp, string(val))

	return strResp, nil
}

func ChangeProvinceFunc(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	var isp string
	uid := r.Form["uid"][0]
	province := r.Form["province"][0]
	isps, ok := r.Form["isp"]
	if ok {
		isp = isps[0]
	} else {
		isp = "1"
	}

	msg, err := ChangeUserHost(uid, province, isp)
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		w.Write([]byte(msg))
	}
}

func main() {
	addr = flag.String("s", "127.0.0.1:8888", "ssdb server host:port")
	flag.Parse()

	http.HandleFunc("/", ChangeProvinceFunc)
	http.ListenAndServe(":8999", nil)
}
