package eqtypes

import (
	"encoding/binary"

	"github.com/goccy/go-json"
)

type Colors uint32

func (c Colors) MarshalJSON() ([]byte, error) {
	values := struct {
		Red   uint8 `json:"red"`
		Green uint8 `json:"green"`
		Blue  uint8 `json:"blue"`
		Alpha uint8 `json:"alpha"`
	}{}

	buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, uint32(c))

	values.Blue = buf[0]
	values.Green = buf[1]
	values.Red = buf[2]
	values.Alpha = buf[3]
	return json.Marshal(values)
}
