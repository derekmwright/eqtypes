package eqtypes

import (
	"database/sql"
	"time"

	"github.com/goccy/go-json"
)

type JsonNullTime struct {
	sql.NullTime
}

func (v JsonNullTime) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Time)
	}
	return json.Marshal(nil)

}

func (v *JsonNullTime) UnmarshalJSON(data []byte) error {
	// Unmarshalling into a pointer will let us detect null
	var x *time.Time

	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}

	if x != nil {
		v.Valid = true
		v.Time = *x
	} else {
		v.Valid = false
	}

	return nil
}

type JsonNullString struct {
	sql.NullString
}

func (v JsonNullString) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.String)
	}
	return json.Marshal("")
}

func (v *JsonNullString) UnmarshalJSON(data []byte) error {
	var x *string

	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}

	if x != nil {
		v.Valid = true
		v.String = *x
	} else {
		v.Valid = false
	}

	return nil
}

type UnixTimestamp struct {
	time.Time
}

func (v UnixTimestamp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.Unix())
}

func (v *UnixTimestamp) UnmarshalJSON(data []byte) error {
	var t time.Time
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*v = UnixTimestamp{t}
	return nil
}
