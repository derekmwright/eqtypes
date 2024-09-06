package eqtypes

import (
	"strings"

	"github.com/goccy/go-json"
)

type Race int32

type RacesBitmask int32

const (
	RaceHuman Race = 1 << iota
	RaceBarbarian
	RaceErudite
	RaceWoodElf
	RaceHighElf
	RaceDarkElf
	RaceHalfElf
	RaceDwarf
	RaceTroll
	RaceOgre
	RaceHalfling
	RaceGnome
	RaceIksar
	RaceVahShir
	RaceFroglok
	RaceDrakkin
)

var raceToString = map[Race]string{
	RaceHuman:     "Human",
	RaceBarbarian: "Barbarian",
	RaceErudite:   "Erudite",
	RaceWoodElf:   "Wood Elf",
	RaceHighElf:   "High Elf",
	RaceDarkElf:   "Dark Elf",
	RaceHalfElf:   "Half Elf",
	RaceDwarf:     "Dwarf",
	RaceTroll:     "Troll",
	RaceOgre:      "Ogre",
	RaceHalfling:  "Halfling",
	RaceGnome:     "Gnome",
	RaceIksar:     "Iksar",
	RaceVahShir:   "Vah Shir",
	RaceFroglok:   "Froglok",
	RaceDrakkin:   "Drakkin",
}

var raceToShortString = map[Race]string{
	RaceHuman:     "HUM",
	RaceBarbarian: "BAR",
	RaceErudite:   "ERU",
	RaceWoodElf:   "ELF",
	RaceHighElf:   "HIE",
	RaceDarkElf:   "DEF",
	RaceHalfElf:   "HEF",
	RaceDwarf:     "DWF",
	RaceTroll:     "TRL",
	RaceOgre:      "OGR",
	RaceHalfling:  "HFL",
	RaceGnome:     "GNM",
	RaceIksar:     "IKS",
	RaceVahShir:   "VAH",
	RaceFroglok:   "FRG",
	RaceDrakkin:   "DRK",
}

func (r RacesBitmask) list(typ string) []string {
	races := make([]string, 0)

	var i int32
	for i = 1; i <= int32(r); i <<= 1 {
		if i&int32(r) != 0 {
			switch typ {
			case "short":
				races = append(races, raceToShortString[Race(i)])
			default:
				races = append(races, raceToString[Race(i)])
			}
		}
	}

	return races
}

func (r RacesBitmask) String() string {
	list := r.list("short")

	if len(list) == 16 {
		return "ALL"
	}

	return strings.Join(list, " ")
}

// NamesList returns a list of race names in long format.
func (r RacesBitmask) NamesList() []string {
	return r.list("long")
}

// ShortNamesList returns a list of race short names.
func (r RacesBitmask) ShortNamesList() []string {
	return r.list("short")
}

func (r RacesBitmask) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.list("short"))
}
