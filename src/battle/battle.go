package battle

import (
	"autohaipeaghr7/src/model"
	"autohaipeaghr7/src/util"
)

type Battle struct {
	Teams []model.Team
	Turn  int
}

func NewBattle(teams []model.Team) Battle {
	btl := Battle{
		Teams: teams,
	}
	return btl
}

func (b Battle) Step() {

}

func (b Battle) FrontFight() {
	b.TrigEvent(util.E_TRN_START)

	for _, at := range b.Teams {
		for _, bt := range b.Teams {
			a := at.Members[0]
			b := bt.Members[0]
			if a == nil ||
				b == nil ||
				a == b ||
				!a.IsAlive() {
				continue
			}

		}
	}

	b.TrigEvent(util.E_TRN_END)
}

func (b Battle) TrigEvent(eID string) {
	for _, t := range b.Teams {
		t.TrigEvent(eID)
	}
}
