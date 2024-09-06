package eqtypes

type DiscoveredItem struct {
	ItemID         uint32        `json:"item_id,omitempty" db:"discovered_items:item_id"`
	CharName       string        `json:"character_name,omitempty" db:"discovered_items:char_name"`
	DiscoveredDate UnixTimestamp `json:"discovered_date,omitempty" db:"discovered_items:discovered_date"`
	AccountStatus  int32         `json:"account_status,omitempty" db:"discovered_items:account_status"`
}
