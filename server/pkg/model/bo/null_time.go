package bo

import (
	"database/sql"
	"encoding/json"
)

type NullTime struct {
	sql.NullTime
}

func (n *NullTime) MarshalJSON() ([]byte, error) {
	if n.Valid {
		res, err := n.Time.MarshalJSON()
		return res, err
	} else {
		return json.Marshal("")
	}
}

func (n *NullTime) UnmarshalJSON(b []byte) error {
	if string(b) == "\"\"" {
		return nil
	}
	err := n.Time.UnmarshalJSON(b)
	if err != nil {
		return err
	}
	if !n.Time.IsZero() {
		n.Valid = true
	}
	return nil
}
