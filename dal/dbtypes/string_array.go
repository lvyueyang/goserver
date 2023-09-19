package dbtypes

import (
	"database/sql/driver"
	"strings"
)

type StringArray []string

func (m *StringArray) Scan(val interface{}) error {
	s := val.([]byte)
	ss := strings.Split(string(s), ",")
	if len(ss) == 1 && ss[0] == "" {
		ss = []string{}
	}
	*m = ss
	return nil
}
func (m StringArray) Value() (driver.Value, error) {
	str := strings.Join(m, ",")
	return str, nil
}
