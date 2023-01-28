package model

import (
	"fmt"
	"strings"

	fs "github.com/Sam36502/funcscript"
)

type Consumable struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Effects []Effect `json:"effects"`
}

func NewConsumable(ID string) Consumable {
	for _, c := range g_DATA_FILE.Consumables {
		if strings.EqualFold(ID, c.ID) {
			return c
		}
	}
	return Consumable{}
}

func (c Consumable) TrigEvent(eID string) {
	for _, e := range c.Effects {
		if strings.EqualFold(e.TriggerEventID, eID) {
			if _, err := fs.Eval(e.FuncScript); err != nil {
				fmt.Printf("[Error]: Failed to execute script on event '%s':\n%v\n", eID, err)
			}
		}
	}
}
