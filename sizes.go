package eqtypes

import "encoding/json"

type Size int32

var sizeMap = map[Size]string{
	0: "TINY",
	1: "SMALL",
	2: "MEDIUM",
	3: "LARGE",
	4: "GIANT",
}

func (s Size) String() string {
	if size, ok := sizeMap[s]; ok {
		return size
	}

	return "UNKNOWN"
}

func (s Size) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}
