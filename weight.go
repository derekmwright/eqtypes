package eqtypes

import "fmt"

type Weight int32

func (w Weight) String() string {
	return fmt.Sprintf("%.1f", float64(w)/10)
}

func (w Weight) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", w.String())), nil
}
