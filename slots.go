package eqtypes

import (
	"encoding/json"
	"strings"
)

type (
	Slot int

	SlotsBitmask int32
)

const (
	SlotCharm Slot = 1 << iota
	SlotEarL
	SlotHead
	SlotFace
	SlotEarR
	SlotNeck
	SlotShoulders
	SlotArms
	SlotBack
	SlotWristL
	SlotWristR
	SlotRange
	SlotHands
	SlotSecondary
	SlotPrimary
	SlotFingerL
	SlotFingerR
	SlotChest
	SlotLegs
	SlotFeet
	SlotWaist
	SlotPowerSource
	SlotAmmo
)

var slotToString = map[Slot]string{
	SlotCharm:       "Charm",
	SlotEarL:        "Ear",
	SlotHead:        "Head",
	SlotFace:        "Face",
	SlotEarR:        "Ear",
	SlotNeck:        "Neck",
	SlotShoulders:   "Shoulders",
	SlotArms:        "Arms",
	SlotBack:        "Back",
	SlotWristL:      "Wrist",
	SlotWristR:      "Wrist",
	SlotRange:       "Range",
	SlotHands:       "Hands",
	SlotSecondary:   "Secondary",
	SlotPrimary:     "Primary",
	SlotFingerL:     "Finger",
	SlotFingerR:     "Finger",
	SlotChest:       "Chest",
	SlotLegs:        "Legs",
	SlotFeet:        "Feet",
	SlotWaist:       "Waist",
	SlotPowerSource: "Power Source",
	SlotAmmo:        "Ammo",
}

func (s Slot) String() string {
	return slotToString[s]
}

// list parses the SlotsBitmask and returns an array of strings with duplicate slots removed.
func (s SlotsBitmask) list() []string {
	var slots []Slot

	var i int32
	for i = 1; i <= int32(s); i <<= 1 {
		if i&int32(s) != 0 {
			slots = append(slots, Slot(i))
		}
	}

	type unique map[string]struct{}

	// There are duplicate entries for things like ears, wrists that don't matter for item display.
	// Create a unique map of slots by name, then concatenate them.
	tmp := make(unique)
	for _, slot := range slots {
		tmp[slotToString[slot]] = struct{}{}
	}

	var sl []string
	for slot := range tmp {
		sl = append(sl, slot)
	}

	return sl
}

func (s SlotsBitmask) String() string {
	return strings.Join(s.list(), " ")
}

func (s SlotsBitmask) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.list())
}
