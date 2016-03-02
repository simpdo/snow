package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
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

func UpdateProvincePay(channel string, isp int, province_pay map[string][]PayPriority) error {
	data, err := ioutil.ReadFile(FilePathConfig.ProvincePayConfigFile)
	if err != nil {
		return err
	}

	var config ProvincePayPriorityConfig
	err = json.Unmarshal(data, &config)
	if err != nil {
		return err
	}

	for i := 0; i < len(config.Config); i++ {
		value := config.Config[i]
		if channel != value.Channel || isp != value.Isp {
			continue
		}

		for j := 0; j < len(config.Config[i].Config); j++ {
			province := config.Config[i].Config[j].Province
			pay_priority, ok := province_pay[province]
			if !ok {
				continue
			}

			config.Config[i].Config[j].PayPrioritys = pay_priority
		}
	}

	data, err = json.Marshal(config)
	if err != nil {
		return err
	}

	ioutil.WriteFile(province_pay_file, data, os.ModePerm)
	return nil

}

func GetProvincePay(channel string, isp int) ([]byte, error) {
	data, err := ioutil.ReadFile(province_pay_file)
	if err != nil {
		return nil, err
	}

	var config ProvincePayPriorityConfig
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	for _, val := range config.Config {
		if channel != val.Channel || isp != val.Isp {
			continue
		}

		buff, err := json.Marshal(val)
		if err != nil {
			return nil, err
		}
		return buff, nil
	}

	msg := fmt.Sprintf("channel[%s], isp[%d] config is not exist!")
	return nil, errors.New(msg)
}
