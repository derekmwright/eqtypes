package eqtypes

type Stats struct {
	AC     int32 `json:"ac,omitempty"`
	HP     int32 `json:"hp,omitempty"`
	Mana   int32 `json:"mana,omitempty"`
	End    int32 `json:"end,omitempty"`
	Purity int32 `json:"purity,omitempty"`
	Str    int32 `json:"strength,omitempty"`
	Sta    int32 `json:"stamina,omitempty"`
	Int    int32 `json:"intelligence,omitempty"`
	Wis    int32 `json:"wisdom,omitempty"`
	Agi    int32 `json:"agility,omitempty"`
	Dex    int32 `json:"dexterity,omitempty"`
	Cha    int32 `json:"charisma,omitempty"`
}

type Resists struct {
	Magic   int32 `json:"magic,omitempty"`
	Fire    int32 `json:"fire,omitempty"`
	Cold    int32 `json:"cold,omitempty"`
	Disease int32 `json:"disease,omitempty"`
	Poison  int32 `json:"poison,omitempty"`
	Corrupt int32 `json:"corrupt,omitempty"`
}
