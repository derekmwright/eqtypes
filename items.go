package eqtypes

type Item struct {
	ID                    int32           `json:"id" db:"items.id"`
	MinStatus             int16           `json:"min_status,omitempty" db:"items.minstatus"`
	Name                  string          `json:"name,omitempty" db:"items.Name"`
	Agi                   int32           `json:"-" db:"items.aagi"`
	AC                    int32           `json:"-" db:"items.ac"`
	Accuracy              int32           `json:"accuracy,omitempty" db:"items.accuracy"`
	Cha                   int32           `json:"-" db:"items.acha"`
	Dex                   int32           `json:"-" db:"items.adex"`
	Int                   int32           `json:"-" db:"items.aint"`
	ArtifactFlag          *uint8          `json:"artifact_flag,omitempty" db:"items.artifactflag"`
	Sta                   int32           `json:"-" db:"items.asta"`
	Str                   int32           `json:"-" db:"items.astr"`
	Attack                int32           `json:"attack,omitempty" db:"items.attack"`
	AugRestrict           int32           `json:"aug_restrictions,omitempty" db:"items.augrestrict"`
	AugSlot1Type          int8            `json:"aug_slot_1_type,omitempty" db:"items.augslot1type"`
	AugSlot1Visible       int8            `json:"aug_slot_1_visible,omitempty" db:"items.augslot1visible"`
	AugSlot2Type          int8            `json:"aug_slot_2_type,omitempty" db:"items.augslot2type"`
	AugSlot2Visible       int8            `json:"aug_slot_2_visible,omitempty" db:"items.augslot2visible"`
	AugSlot3Type          int8            `json:"aug_slot_3_type,omitempty" db:"items.augslot3type"`
	AugSlot3Visible       int8            `json:"aug_slot_3_visible,omitempty" db:"items.augslot3visible"`
	AugSlot4Type          int8            `json:"aug_slot_4_type,omitempty" db:"items.augslot4type"`
	AugSlot4Visible       int8            `json:"aug_slot_4_visible,omitempty" db:"items.augslot4visible"`
	AugSlot5Type          int8            `json:"aug_slot_5_type,omitempty" db:"items.augslot5type"`
	AugSlot5Visible       int8            `json:"aug_slot_5_visible,omitempty" db:"items.augslot5visible"`
	AugSlot6Type          int8            `json:"aug_slot_6_type,omitempty" db:"items.augslot6type"`
	AugSlot6Visible       int8            `json:"aug_slot_6_visible,omitempty" db:"items.augslot6visible"`
	AugType               int32           `json:"augment_type,omitempty" db:"items.augtype"`
	Avoidance             int32           `json:"avoidance,omitempty" db:"items.avoidance"`
	Wis                   int32           `json:"-" db:"items.awis"`
	BagSize               Size            `json:"bag_size,omitempty" db:"items.bagsize"`
	BagSlots              int32           `json:"bag_slots,omitempty" db:"items.bagslots"`
	BagType               int32           `json:"bag_type,omitempty" db:"items.bagtype"`
	BagWeightReduction    int32           `json:"bag_weight_reduction,omitempty" db:"items.bagwr"`
	BaneDamageAmount      int32           `json:"bane_damage_amount,omitempty" db:"items.banedmgamt"`
	BaneDamageRaceAmount  int32           `json:"bane_damage_race_amount,omitempty" db:"items.banedmgraceamt"`
	BaneDamageBody        int32           `json:"bane_damage_body_amount,omitempty" db:"items.banedmgbody"`
	BaneDamageRace        NPCRace         `json:"bane_damage_race,omitempty" db:"items.banedmgrace"`
	BardType              int32           `json:"bard_type,omitempty" db:"items.bardtype"`
	BardValue             int32           `json:"bard_value,omitempty" db:"items.bardvalue"`
	Book                  int32           `json:"book,omitempty" db:"items.book"`
	CastTime              int32           `json:"cast_time,omitempty" db:"items.casttime"`
	_                     int32           `db:"items.casttime_"`
	CharmFile             string          `json:"charm_file,omitempty" db:"items.charmfile"`
	CharmFileID           string          `json:"charm_file_id,omitempty" db:"items.charmfileid"`
	Classes               ClassesBitmask  `json:"classes,omitempty" db:"items.classes"`
	Color                 Colors          `json:"color,omitempty" db:"items.color"`
	CombatEffects         string          `json:"combat_effects,omitempty" db:"items.combateffects"`
	ExtraDamageSkill      int32           `json:"extra_damage_skill,omitempty" db:"items.extradmgskill"`
	ExtraDamageAmount     int32           `json:"extra_damage_amount,omitempty" db:"items.extradmgamt"`
	Price                 int32           `json:"price,omitempty" db:"items.price"`
	CR                    int32           `json:"-" db:"items.cr"`
	Damage                int32           `json:"damage,omitempty" db:"items.damage"`
	DamageShield          int32           `json:"damage_shield,omitempty" db:"items.damageshield"`
	Deity                 int32           `json:"deity,omitempty" db:"items.deity"`
	Delay                 int32           `json:"delay,omitempty" db:"items.delay"`
	AugDistiller          uint32          `json:"augment_distiller,omitempty" db:"items.augdistiller"`
	DotShielding          int32           `json:"dot_shielding,omitempty" db:"items.dotshielding"`
	DR                    int32           `json:"-" db:"items.dr"`
	ClickType             int32           `json:"click_type,omitempty" db:"items.clicktype"`
	ClickLevel2           int32           `json:"click_level_2,omitempty" db:"items.clicklevel2"`
	ElementalDamageType   int32           `json:"elemental_damage_type,omitempty" db:"items.elemdmgtype"`
	ElementalDamageAmount int32           `json:"elemental_damage_amount,omitempty" db:"items.elemdmgamt"`
	Endurance             int32           `json:"endurance,omitempty" db:"items.endur"`
	FactionAmount1        int32           `json:"faction_amount_1,omitempty" db:"items.factionamt1"`
	FactionAmount2        int32           `json:"faction_amount_2,omitempty" db:"items.factionamt2"`
	FactionAmount3        int32           `json:"faction_amount_3,omitempty" db:"items.factionamt3"`
	FactionAmount4        int32           `json:"faction_amount_4,omitempty" db:"items.factionamt4"`
	FactionMod1           int32           `json:"faction_mod_1,omitempty" db:"items.factionmod1"`
	FactionMod2           int32           `json:"faction_mod_2,omitempty" db:"items.factionmod2"`
	FactionMod3           int32           `json:"faction_mod_3,omitempty" db:"items.factionmod3"`
	FactionMod4           int32           `json:"faction_mod_4,omitempty" db:"items.factionmod4"`
	Filename              string          `json:"filename,omitempty" db:"items.filename"`
	FocusEffect           *int32          `json:"focus_effect,omitempty" db:"items.focuseffect"`
	FR                    int32           `json:"-" db:"items.fr"`
	FVNoDrop              *int32          `json:"fv_no_drop,omitempty" db:"items.fvnodrop"`
	Haste                 int32           `json:"haste,omitempty" db:"items.haste"`
	ClickLevel            int32           `json:"click_level,omitempty" db:"items.clicklevel"`
	HP                    int32           `json:"-" db:"items.hp"`
	Regen                 int32           `json:"regen,omitempty" db:"items.regen"`
	Icon                  int32           `json:"icon,omitempty" db:"items.icon"`
	IDFile                string          `json:"id_file,omitempty" db:"items.idfile"`
	ItemClass             int32           `json:"item_class,omitempty" db:"items.itemclass"`
	ItemType              ItemType        `json:"item_type,omitempty" db:"items.itemtype"`
	LDONPrice             int32           `json:"ldon_price,omitempty" db:"items.ldonprice"`
	LDONTheme             int32           `json:"ldon_theme,omitempty" db:"items.ldontheme"`
	LDONSold              int32           `json:"ldon_sold,omitempty," db:"items.ldonsold"`
	Light                 int32           `json:"light,omitempty" db:"items.light"`
	Lore                  string          `json:"lore,omitempty" db:"items.lore"`
	LoreGroup             *int32          `json:"lore_group,omitempty" db:"items.loregroup"`
	Magic                 int32           `json:"magic,omitempty" db:"items.magic"`
	Mana                  int32           `json:"-" db:"items.mana"`
	ManaRegen             int32           `json:"mana_regen,omitempty" db:"items.manaregen"`
	EnduranceRegen        int32           `json:"endurance_regen,omitempty" db:"items.enduranceregen"`
	Material              int32           `json:"material,omitempty" db:"items.material"`
	HeroesForgeModel      int32           `json:"heroes_forge_model,omitempty" db:"items.herosforgemodel"`
	MaxCharges            int32           `json:"max_charges,omitempty" db:"items.maxcharges"`
	MR                    int32           `json:"-" db:"items.mr"`
	NoDrop                *int32          `json:"-" db:"items.nodrop"`
	NoRent                *int32          `json:"-" db:"items.norent"`
	PendingLoreflag       *uint8          `json:"pending_lore_flag,omitempty" db:"items.pendingloreflag"`
	PR                    int32           `json:"-" db:"items.pr"`
	ProcRate              int32           `json:"proc_rate,omitempty" db:"items.procrate"`
	Races                 RacesBitmask    `json:"races,omitempty" db:"items.races"`
	Range                 int32           `json:"range,omitempty" db:"items.range"`
	RecLevel              int32           `json:"rec_level,omitempty" db:"items.reclevel"`
	RecSkill              int32           `json:"rec_skill,omitempty" db:"items.recskill"`
	ReqLevel              int32           `json:"req_level,omitempty" db:"items.reqlevel"`
	SellRate              float64         `json:"sell_rate,omitempty" db:"items.sellrate"`
	Shielding             int32           `json:"shielding,omitempty" db:"items.shielding"`
	Size                  Size            `json:"size,omitempty" db:"items.size"`
	SkillModType          *int32          `json:"skill_mod_type,omitempty" db:"items.skillmodtype"`
	SkillModValue         int32           `json:"skill_mod_value,omitempty" db:"items.skillmodvalue"`
	Slots                 SlotsBitmask    `json:"slots,omitempty" db:"items.slots"`
	ClickEffect           *int32          `json:"click_effect,omitempty" db:"items.clickeffect"`
	SpellShield           int32           `json:"spellshield,omitempty" db:"items.spellshield"`
	Strikethrough         int32           `json:"strikethrough,omitempty" db:"items.strikethrough"`
	StunResist            int32           `json:"stunresist,omitempty" db:"items.stunresist"`
	SummonedFlag          uint8           `json:"summoned_flag,omitempty" db:"items.summonedflag"`
	Tradeskills           int32           `json:"tradeskills,omitempty" db:"items.tradeskills"`
	Favor                 int32           `json:"favor,omitempty" db:"items.favor"`
	Weight                Weight          `json:"weight,omitempty" db:"items.weight"`
	_                     int32           `db:"items.unk012"`
	_                     int32           `db:"items.unk013"`
	BenefitFlag           int32           `json:"benefit_flag,omitempty" db:"items.benefitflag"`
	_                     int32           `db:"items.unk054"`
	_                     int32           `db:"items.unk059"`
	BookType              int32           `json:"book_type,omitempty" db:"items.booktype"`
	RecastDelay           int32           `json:"recast_delay,omitempty" db:"items.recastdelay"`
	RecastType            int32           `json:"recast_type,omitempty" db:"items.recasttype"`
	GuildFavor            int32           `json:"guild_favor,omitempty" db:"items.guildfavor"`
	_                     int32           `db:"items.unk123"`
	_                     int32           `db:"items.unk124"`
	Attuneable            int32           `json:"attuneable,omitempty" db:"items.attuneable"`
	NoPet                 int32           `json:"no_pet,omitempty" db:"items.nopet"`
	Updated               *JsonNullTime   `json:"updated,omitempty" db:"items.updated"`
	Comment               string          `json:"comment,omitempty" db:"items.comment"`
	_                     int32           `db:"items.unk127"`
	PointType             int32           `json:"potion_type,omitempty" db:"items.pointtype"`
	PotionBelt            int32           `json:"potion_belt,omitempty" db:"items.potionbelt"`
	PotionBeltSlots       int32           `json:"potion_belt_slots,omitempty" db:"items.potionbeltslots"`
	StackSize             int32           `json:"stack_size,omitempty" db:"items.stacksize"`
	NoTransfer            int32           `json:"no_transfer,omitempty" db:"items.notransfer"`
	Stackable             int32           `json:"stackable,omitempty" db:"items.stackable"`
	_                     string          `db:"items.unk134"`
	_                     int32           `db:"items.unk137"`
	ProcEffect            *int32          `json:"proc_effect,omitempty" db:"items.proceffect"`
	ProcType              int32           `json:"proc_type,omitempty" db:"items.proctype"`
	ProcLevel2            int32           `json:"proc_level_2,omitempty" db:"items.proclevel2"`
	ProcLevel             int32           `json:"proc_level,omitempty" db:"items.proclevel"`
	_                     int32           `db:"items.unk142"`
	WornEffect            *int32          `json:"worn_effect,omitempty" db:"items.worneffect"`
	WornType              int32           `json:"worn_type,omitempty" db:"items.worntype"`
	WornLevel2            int32           `json:"worn_level_2,omitempty" db:"items.wornlevel2"`
	WornLevel             int32           `json:"worn_level,omitempty" db:"items.wornlevel"`
	_                     int32           `db:"items.unk147"`
	FocusType             int32           `json:"focus_type,omitempty" db:"items.focustype"`
	FocusLevel2           int32           `json:"focus_level_2,omitempty" db:"items.focuslevel2"`
	FocusLevel            int32           `json:"focus_level,omitempty" db:"items.focuslevel"`
	_                     int32           `db:"items.unk152"`
	ScrollEffect          *int32          `json:"scroll_effect,omitempty" db:"items.scrolleffect"`
	ScrollType            int32           `json:"scroll_type,omitempty" db:"items.scrolltype"`
	ScrollLevel2          int32           `json:"scroll_level_2,omitempty" db:"items.scrolllevel2"`
	ScrollLevel           int32           `json:"scroll_level,omitempty" db:"items.scrolllevel"`
	_                     int32           `db:"items.unk157"`
	Serialized            *JsonNullTime   `json:"serialized,omitempty" db:"items.serialized"`
	Verified              *JsonNullTime   `json:"verified,omitempty" db:"items.verified"`
	Serialization         *JsonNullString `json:"serialization,omitempty" db:"items.serialization"`
	Source                string          `json:"-" db:"items.source"`
	_                     int32           `db:"items.unk033"`
	LoreFile              string          `json:"lore_file,omitempty" db:"items.lorefile"`
	_                     int32           `db:"items.unk014"`
	SvCorruption          int32           `json:"-" db:"items.svcorruption"`
	SkillModMax           int32           `json:"skill_mod_max,omitempty" db:"items.skillmodmax"`
	_                     int32           `db:"items.unk060"`
	_                     int32           `db:"items.augslot1unk2"`
	_                     int32           `db:"items.augslot2unk2"`
	_                     int32           `db:"items.augslot3unk2"`
	_                     int32           `db:"items.augslot4unk2"`
	_                     int32           `db:"items.augslot5unk2"`
	_                     int32           `db:"items.augslot6unk2"`
	_                     int32           `db:"items.unk120"`
	_                     int32           `db:"items.unk121"`
	QuestItemFlag         int32           `json:"quest_item_flag,omitempty" db:"items.questitemflag"`
	_                     *JsonNullString `db:"items.unk132"`
	_                     int32           `db:"items.clickunk5"`
	_                     string          `db:"items.clickunk6"`
	_                     int32           `db:"items.clickunk7"`
	_                     int32           `db:"items.procunk1"`
	_                     int32           `db:"items.procunk2"`
	_                     int32           `db:"items.procunk3"`
	_                     int32           `db:"items.procunk4"`
	_                     string          `db:"items.procunk6"`
	_                     int32           `db:"items.procunk7"`
	_                     int32           `db:"items.wornunk1"`
	_                     int32           `db:"items.wornunk2"`
	_                     int32           `db:"items.wornunk3"`
	_                     int32           `db:"items.wornunk4"`
	_                     int32           `db:"items.wornunk5"`
	_                     string          `db:"items.wornunk6"`
	_                     int32           `db:"items.wornunk7"`
	_                     int32           `db:"items.focusunk1"`
	_                     int32           `db:"items.focusunk2"`
	_                     int32           `db:"items.focusunk3"`
	_                     int32           `db:"items.focusunk4"`
	_                     int32           `db:"items.focusunk5"`
	_                     string          `db:"items.focusunk6"`
	_                     int32           `db:"items.focusunk7"`
	_                     uint32          `db:"items.scrollunk1"`
	_                     int32           `db:"items.scrollunk2"`
	_                     int32           `db:"items.scrollunk3"`
	_                     int32           `db:"items.scrollunk4"`
	_                     int32           `db:"items.scrollunk5"`
	_                     string          `db:"items.scrollunk6"`
	_                     int32           `db:"items.scrollunk7"`
	_                     int32           `db:"items.unk193"`
	Purity                int32           `json:"purity,omitempty" db:"items.purity"`
	EvoItem               int32           `json:"evolving_item,omitempty" db:"items.evoitem"`
	EvoID                 int32           `json:"evolving_id,omitempty" db:"items.evoid"`
	EvolvingLevel         int32           `json:"evolving_level,omitempty" db:"items.evolvinglevel"`
	EvoMax                int32           `json:"evolving_max,omitempty" db:"items.evomax"`
	ClickName             string          `json:"click_name,omitempty" db:"items.clickname"`
	ProcName              string          `json:"proc_name,omitempty" db:"items.procname"`
	WornName              string          `json:"worn_name,omitempty" db:"items.wornname"`
	FocusName             string          `json:"focus_name,omitempty" db:"items.focusname"`
	ScrollName            string          `json:"scroll_name,omitempty" db:"items.scrollname"`
	DSMitigation          int16           `json:"damage_shield_mitigation,omitempty" db:"items.dsmitigation"`
	HeroicStr             int16           `json:"heroic_str,omitempty" db:"items.heroicstr"`
	HeroicInt             int16           `json:"heroic_int,omitempty" db:"items.heroicint"`
	HeroicWis             int16           `json:"heroic_wis,omitempty" db:"items.heroicwis"`
	HeroicAgi             int16           `json:"heroic_agi,omitempty" db:"items.heroicagi"`
	HeroicDex             int16           `json:"heroic_dex,omitempty" db:"items.heroicdex"`
	HeroicSta             int16           `json:"heroic_sta,omitempty" db:"items.heroicsta"`
	HeroicCha             int16           `json:"heroic_cha,omitempty" db:"items.heroiccha"`
	HeroicPr              int16           `json:"heroic_pr,omitempty" db:"items.heroicpr"`
	HeroicDr              int16           `json:"heroic_dr,omitempty" db:"items.heroicdr"`
	HeroicFr              int16           `json:"heroic_fr,omitempty" db:"items.heroicfr"`
	HeroicCr              int16           `json:"heroic_cr,omitempty" db:"items.heroiccr"`
	HeroicMr              int16           `json:"heroic_mr,omitempty" db:"items.heroicmr"`
	HeroicSvCorruption    int16           `json:"heroic_svcorruption,omitempty" db:"items.heroicsvcorrup"`
	HealAmt               int16           `json:"heal_amount,omitempty" db:"items.healamt"`
	SpellDamage           int16           `json:"spell_damage,omitempty" db:"items.spelldmg"`
	Clairvoyance          int16           `json:"clairvoyance,omitempty" db:"items.clairvoyance"`
	BackstabDamage        int16           `json:"backstab_damage,omitempty" db:"items.backstabdmg"`
	Created               string          `json:"created,omitempty" db:"items.created"`
	EliteMaterial         int16           `json:"elite_material,omitempty" db:"items.elitematerial"`
	LDONSellBackRate      int16           `json:"ldon_sell_back_rate,omitempty" db:"items.ldonsellbackrate"`
	ScriptFileID          int32           `json:"script_field_id,omitempty" db:"items.scriptfileid"`
	ExpendableArrow       int16           `json:"expendable_arrow,omitempty" db:"items.expendablearrow"`
	PowerSourceCapacity   int32           `json:"power_source_capacity,omitempty" db:"items.powersourcecapacity"`
	BardEffect            *int32          `json:"bard_effect,omitempty" db:"items.bardeffect"`
	BardEffectType        int16           `json:"bard_effect_type,omitempty" db:"items.bardeffecttype"`
	BardLevel2            int16           `json:"bard_level_2,omitempty" db:"items.bardlevel2"`
	BardLevel             int16           `json:"bard_level,omitempty" db:"items.bardlevel"`
	_                     int16           `db:"items.bardunk1"`
	_                     int16           `db:"items.bardunk2"`
	_                     int16           `db:"items.bardunk3"`
	_                     int16           `db:"items.bardunk4"`
	_                     int16           `db:"items.bardunk5"`
	_                     string          `db:"items.bardname"`
	_                     int16           `db:"items.bardunk7"`
	_                     int16           `db:"items.unk214"`
	SubType               int32           `json:"subtype,omitempty" db:"items.subtype"`
	_                     int32           `db:"items.unk220"`
	_                     int32           `db:"items.unk221"`
	Heirloom              int32           `json:"heirloom,omitempty" db:"items.heirloom"`
	_                     int32           `db:"items.unk223"`
	_                     int32           `db:"items.unk224"`
	_                     int32           `db:"items.unk225"`
	_                     int32           `db:"items.unk226"`
	_                     int32           `db:"items.unk227"`
	_                     int32           `db:"items.unk228"`
	_                     int32           `db:"items.unk229"`
	_                     int32           `db:"items.unk230"`
	_                     int32           `db:"items.unk231"`
	_                     int32           `db:"items.unk232"`
	_                     int32           `db:"items.unk233"`
	_                     int32           `db:"items.unk234"`
	Placeable             *int32          `json:"-" db:"items.placeable"`
	_                     int32           `db:"items.unk236"`
	_                     int32           `db:"items.unk237"`
	_                     int32           `db:"items.unk238"`
	_                     int32           `db:"items.unk239"`
	_                     int32           `db:"items.unk240"`
	_                     int32           `db:"items.unk241"`
	EpicItem              int32           `json:"epic_item,omitempty" db:"items.epicitem"`
}

func (i *Item) MarshalJSON() ([]byte, error) {
	if *i.BardEffect < 0 {
		i.BardEffect = nil
	}

	if *i.ClickEffect < 0 {
		i.ClickEffect = nil
	}

	if *i.FocusEffect < 0 {
		i.FocusEffect = nil
	}

	if *i.ProcEffect < 0 {
		i.ProcEffect = nil
	}

	if *i.ScrollEffect < 0 {
		i.ScrollEffect = nil
	}

	if *i.WornEffect < 0 {
		i.WornEffect = nil
	}

	if *i.SkillModType < 0 {
		i.SkillModType = nil
	}

	var tags []string

	if *i.LoreGroup != 0 {
		tags = append(tags, "Lore")
	}

	if *i.NoDrop == 0 {
		tags = append(tags, "No Trade")
	}

	if *i.NoRent == 0 {
		tags = append(tags, "Temporary")
	}

	if *i.Placeable == 0 {
		tags = append(tags, "Placeable")
	}

	type Alias Item
	m := struct {
		Alias
		Stats   Stats   `json:"stats"`
		Resists Resists `json:"resists"`

		ClassesString string   `json:"classes_string"`
		RacesString   string   `json:"races_string"`
		Tags          []string `json:"tags"`
	}{
		Alias: Alias(*i),
		Stats: Stats{
			AC:     i.AC,
			HP:     i.HP,
			Mana:   i.Mana,
			End:    i.Endurance,
			Purity: i.Purity,
			Str:    i.Str,
			Sta:    i.Sta,
			Int:    i.Int,
			Wis:    i.Wis,
			Agi:    i.Agi,
			Dex:    i.Dex,
		},
		Resists: Resists{
			Magic:   i.MR,
			Fire:    i.FR,
			Cold:    i.CR,
			Disease: i.DR,
			Poison:  i.PR,
			Corrupt: i.SvCorruption,
		},
		ClassesString: i.Classes.String(),
		RacesString:   i.Races.String(),
		Tags:          tags,
	}

	return json.Marshal(m)
}
