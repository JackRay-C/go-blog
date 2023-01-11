package bo

import (
	"database/sql/driver"
	"strconv"
	"time"
)

type LocalTime struct {
	int64
}

func (t *LocalTime) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (t *LocalTime) UnmarshalJSON(b []byte) error {
	return nil
}

func (t *LocalTime) Value() (driver.Value, error) {
	return t, nil
}

func (t *LocalTime) Scan(v interface{}) error {
	_ = time.Unix(t.int64, 0)

	return nil
}
func (t *LocalTime) String() string {
	return strconv.FormatInt(t.int64, 10)
}
