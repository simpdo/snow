package config

import (
	"encoding/json"
	"fmt"
	"github.com/tealeg/xlsx"
	"io/ioutil"
	"os"
)

type RoundBoxConfig struct {
	Boxes []RoomBoxConfig `json:"room_box"`
}

type RoomBoxConfig struct {
	RoomId     int            `json:"room"`
	GoodsId    int            `json:"goods"`
	LevelBoxes []LevelBox     `json:"round_box"`
	LevelRule  []LevelBoxRule `json:"-"`
}

type LevelBoxRule struct {
	BoxId         int `json:"-"`
	LowLimitAdd   int `json:"-"`
	UpperLimitAdd int `json:"-"`
}

type LevelBox struct {
	MeetRounds int       `json:"meet_rounds"`
	Boxes      []BoxItem `json:"boxes"`
}

type BoxItem struct {
	BoxId        int              `json:"box_id"`
	BoxType      int              `json:"box_type"`
	PropId       int              `json:"prop_id"`
	PropName     string           `json:"prop_name"`
	RewardIcon   string           `json:"reward_icon"`
	RewardMin    int              `json:"reward_min"`
	RewardMax    int              `json:"reward_max"`
	FirstBoxRate float64          `json:"first_box_rate"`
	SecBoxRate   float64          `json:"sec_box_rate"`
	ReduceRule   []ReduceRateRule `json:"reduce_rule"`
}

type ReduceRateRule struct {
	PropId   int     `json:"prop_id"`
	LimitNum int     `json:"limit_num"`
	Rate     float64 `json:"rate"`
}

func WriteRoundBox2Json(input string, output string) error {
	file, err := xlsx.OpenFile(input)
	if err != nil {
		return err
	}

	config := RoundBoxConfig{}
	room_config := RoomBoxConfig{}

	levels := []int{}
	base_boxes := []BoxItem{}

	sheet := file.Sheets[0]
	fmt.Println(len(sheet.Rows))
	for i := 1; i < len(sheet.Rows); i++ {
		cells := sheet.Rows[i].Cells

		item := BoxItem{}
		level_rule := LevelBoxRule{}

		reduce_rule := ReduceRateRule{}
		for j := 0; j < len(cells); j++ {
			if j == 14 || j == 17 {
				reduce_rule.PropId, _ = cells[j].Int()
				if -1 == reduce_rule.PropId {
					break
				}
			}

			switch j {
			case 0:
				val, _ := cells[j].Int()
				if val > 0 {
					if room_config.RoomId > 0 {
						BuildRoomBox(&room_config, levels, base_boxes)
						config.Boxes = append(config.Boxes, room_config)
					}
					room_config = RoomBoxConfig{}
					room_config.RoomId = val
					levels = []int{}
					base_boxes = []BoxItem{}
				}
			case 1:
				val, _ := cells[j].Int()
				if val > 0 {
					goods := val
					room_config.GoodsId = goods
				}
			case 2:
				val, _ := cells[j].Int()
				levels = append(levels, val)
			case 3:
				item.BoxId, _ = cells[j].Int()
				level_rule.BoxId = item.BoxId
			case 4:
				item.PropName = cells[j].String()
			case 5:
				item.BoxType, _ = cells[j].Int()
			case 6:
				item.PropId, _ = cells[j].Int()
			case 7:
				item.RewardIcon = cells[j].String()
			case 8:
				item.RewardMin, _ = cells[j].Int()
			case 9:
				item.RewardMax, _ = cells[j].Int()
			case 10:
				item.FirstBoxRate, _ = cells[j].Float()
			case 11:
				item.SecBoxRate, _ = cells[j].Float()
			case 12:
				level_rule.LowLimitAdd, _ = cells[j].Int()
			case 13:
				level_rule.UpperLimitAdd, _ = cells[j].Int()
			case 15:
				reduce_rule.LimitNum, _ = cells[j].Int()
			case 16:
				reduce_rule.Rate, _ = cells[j].Float()
				item.ReduceRule = append(item.ReduceRule, reduce_rule)
			case 18:
				reduce_rule.LimitNum, _ = cells[j].Int()
			case 19:
				reduce_rule.Rate, _ = cells[j].Float()
				item.ReduceRule = append(item.ReduceRule, reduce_rule)
			}
		}

		room_config.LevelRule = append(room_config.LevelRule, level_rule)
		if item.BoxId > 0 {
			base_boxes = append(base_boxes, item)
		}
		fmt.Println(base_boxes)
	}

	BuildRoomBox(&room_config, levels, base_boxes)
	config.Boxes = append(config.Boxes, room_config)
	fmt.Println(config)

	data, err := json.Marshal(config)
	if err != nil {
		return err
	}

	fmt.Println(string(data))
	ioutil.WriteFile(output, data, os.ModePerm)

	return nil
}

func BuildRoomBox(room_box *RoomBoxConfig, levels []int, base_boxes []BoxItem) {
	level_box := LevelBox{}
	level_box.MeetRounds = levels[0]
	level_box.Boxes = base_boxes[:]
	room_box.LevelBoxes = append(room_box.LevelBoxes, level_box)

	for i := 1; i < len(levels); i++ {
		level_box = LevelBox{}
		level_box.MeetRounds = levels[i]

		for j, base_box := range base_boxes {
			rule := room_box.LevelRule[j]
			box := base_box
			box.RewardMin = base_box.RewardMin + rule.LowLimitAdd*i
			box.RewardMax = base_box.RewardMax + rule.UpperLimitAdd*i
			level_box.Boxes = append(level_box.Boxes, box)
		}
		room_box.LevelBoxes = append(room_box.LevelBoxes, level_box)
	}
}
