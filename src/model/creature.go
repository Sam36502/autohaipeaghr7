package model

import (
	"autohaipeaghr7/src/util"
	"fmt"
	"strings"

	fs "github.com/Sam36502/funcscript"
)

type Creature struct {
	Name        string
	Damage      int
	Health      int
	Effects     []Effect
	CurrentItem Consumable
}

type CreatureEgg struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Damage  int      `json:"damage"`
	Health  int      `json:"health"`
	Effects []Effect `json:"effects"`
}

func NewCreature(eggID string) *Creature {
	if g_DATA_FILE.Creatures == nil {
		return nil
	}
	for _, c := range g_DATA_FILE.Creatures {
		if strings.EqualFold(c.ID, eggID) {
			return &Creature{
				Name:        c.Name,
				Damage:      c.Damage,
				Health:      c.Health,
				Effects:     c.Effects,
				CurrentItem: Consumable{},
			}
		}
	}
	return nil
}

func (c *Creature) IsAlive() bool {
	return c.Health > 0
}

func (c *Creature) TakeDamage(dmg int) {
	c.Health -= dmg

	if c.IsAlive() {
		c.TrigEvent(util.E_HURT)
	} else {
		c.TrigEvent(util.E_FAINT)
	}
}

func (c *Creature) Consume(item Consumable) {
	item.TrigEvent(util.E_CONSUME)
	c.TrigEvent(util.E_CONSUME)
	c.CurrentItem = item
}

func (c *Creature) TrigEvent(eID string) {
	for _, e := range c.Effects {
		if strings.EqualFold(e.TriggerEventID, eID) {
			if _, err := fs.Eval(e.FuncScript); err != nil {
				fmt.Printf("[Error]: Failed to execute script on event '%s':\n%v\n", eID, err)
			}
		}
	}
}

func (c *Creature) ToString() string {
	return fmt.Sprintf("(%s: % 2d/% 2d)", c.Name, c.Damage, c.Health)
}
