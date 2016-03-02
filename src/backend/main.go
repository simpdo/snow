// config_server project main.go
package main

import (
	"backend/config"
	"backend/handle"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime"
)

type ServerConfig struct {
	Port int `json:"port"`
}

func initLog() {
	file, _ := os.OpenFile("server.log", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
	log.SetOutput(file)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func initServerConfig(config *ServerConfig) {
	data, err := ioutil.ReadFile("../conf/server.json")
	if err != nil {
		fmt.Println("read config file failed, error: ", err.Error())
		os.Exit(-1)
	}

	err = json.Unmarshal(data, config)
	if err != nil {
		fmt.Println("parse config failed, error: ", err.Error())
		os.Exit(-1)
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var server_config ServerConfig
	initServerConfig(&server_config)

	initLog()
	config.Init

	http.HandleFunc("/pay/province", handle.ProvincePayHandle)

	addr := fmt.Sprintln(":%d", server_config.Port)
	http.ListenAndServe(addr, nil)
}
