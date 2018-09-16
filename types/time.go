package types

import (
	"encoding/json"
	"time"

	"github.com/volatiletech/null"
)

type NullTime struct {
	null.Time
}

func (t NullTime) MarshalJSON() ([]byte, error) {
	var ts int64
	if !t.Time.Valid {
		return json.Marshal(nil)
	}
	ts = t.Time.Time.Unix()
	return json.Marshal(ts)
}

func NullTimeFrom(t time.Time) NullTime {
	return NullTime{
		Time: null.TimeFrom(t),
	}
}
