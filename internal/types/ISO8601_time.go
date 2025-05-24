package types

import (
	"database/sql"
	"time"
)

const (
	ISO8601 = "2006-01-02T15:04:05-0700"
)

type ISO8601Time struct {
	sql.NullTime
}

// return time in ISO8601 format
func (t ISO8601Time) String() string {
	if !t.Valid {
		return ""
	}
	return t.Time.Format(ISO8601)
}
// UnmarshalJSON implements the json.Unmarshaler interface for ISO8601Time
func (t *ISO8601Time) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		t.Valid = false
		return nil
	}
	var err error
	t.Time, err = time.Parse(`"`+ISO8601+`"`, string(data))
	if err != nil {
		return err
	}
	t.Valid = true
	return nil
}
// MarshalJSON implements the json.Marshaler interface for ISO8601Time
func (t ISO8601Time) MarshalJSON() ([]byte, error) {
	if !t.Valid {
		return []byte("null"), nil
	}
	return []byte(`"` + t.Time.Format(ISO8601) + `"`), nil
}