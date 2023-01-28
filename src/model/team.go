package model

import (
	"autohaipeaghr7/src/util"
	"fmt"
	"strings"
)

type Team struct {
	Name    string
	Members [util.TEAM_SIZE]*Creature // 0 is front
}

func NewTeamOfIDs(Name string, IDs []string) *Team {
	t := Team{
		Name:    Name,
		Members: [util.TEAM_SIZE]*Creature{},
	}
	for i, ID := range IDs {
		if i > util.TEAM_SIZE {
			break
		}
		t.Members[i] = NewCreature(ID)
	}
	return &t
}

func NewTeamClone(tm *Team) *Team {
	nt := *tm
	return &nt
}

func (t *Team) IsAlive() bool {
	isAlive := false
	for _, m := range t.Members {
		isAlive = isAlive && m.IsAlive()
	}
	return isAlive
}

func (t *Team) CheckFainted() {
	nextVacancy := -1
	for i := 0; i < util.TEAM_SIZE; i++ {
		m := t.Members[i]
		if (m == nil || !m.IsAlive()) && nextVacancy < 0 {
			nextVacancy = i
			fmt.Printf("---> Next Vacant position: %d\n", nextVacancy)
		} else if m != nil && m.IsAlive() && nextVacancy > -1 {
			fmt.Printf("---> '%s' moved into vacancy.\n", m.Name)
			t.Members[nextVacancy] = m
			nextVacancy = i
			fmt.Printf("---> Now %d is vacant\n", i)
		}
	}
	fmt.Println(t.ToString())
}

func (t *Team) TrigEvent(eID string) {
	for _, m := range t.Members {
		if strings.EqualFold(eID, util.E_FAINT) {
			t.CheckFainted()
		}
		m.TrigEvent(eID)
	}
}

func (t *Team) ToString() string {
	sb := strings.Builder{}
	for i, m := range t.Members {
		if m == nil {
			sb.WriteString("()")
		} else {
			sb.WriteString(m.ToString())
		}
		if i < util.TEAM_SIZE {
			sb.WriteRune(' ')
		}
	}
	return sb.String()
}
