package eqtypes

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"reflect"
)

type Flags struct {
	Enabled bool
	value   int32
}

//goland:noinspection GoMixedReceiverTypes
func (f Flags) Scan(value any) error {
	if v, ok := value.(int32); ok {
		f.Enabled = v != 1
		f.value = v
		return nil
	}

	typ := reflect.TypeOf(value)

	return &ScanError{
		got:      typ.String(),
		expected: reflect.Int32.String(),
	}
}

//goland:noinspection GoMixedReceiverTypes
func (f Flags) Value() (driver.Value, error) {
	return f.Enabled, nil
}

//goland:noinspection GoMixedReceiverTypes
func (f Flags) String() string {
	return fmt.Sprintf("%t", f.Enabled)
}

//goland:noinspection GoMixedReceiverTypes
func (f *Flags) MarshalJSON() ([]byte, error) {
	return json.Marshal(f.value)
}
