package util

type EventTriggerable interface {
	TrigEvent(eventID string)
}

const DATA_FILE_NAME = "data.json"
const TEAM_SIZE = 5

// All available events
const (
	E_RND_START = "on_rnd_start" // Start of the round
	E_TRN_START = "on_trn_start" // Start of the turn
	E_TRN_END   = "on_trn_end"   // End of the turn
	E_CONSUME   = "on_consume"   // Consumable consumed
	E_HURT      = "on_hurt"      // Consumable consumed
	E_FAINT     = "on_faint"     // Consumable consumed
)
