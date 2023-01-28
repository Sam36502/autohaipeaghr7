package model

import (
	"encoding/json"
	"io/ioutil"
)

type DataFile struct {
	Creatures   []CreatureEgg `json:"creatures"`
	Consumables []Consumable  `json:"consumables"`
}

var g_DATA_FILE DataFile

func LoadDataFile(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	g_DATA_FILE = DataFile{}
	return json.Unmarshal(data, &g_DATA_FILE)
}
