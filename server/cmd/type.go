package cmd

import "encoding/json"

type stringArray []string

func (s stringArray) Contain(c string) bool {
	for _, a := range s {
		if a == c {
			return true
		}
	}
	return false
}

func (s stringArray) String() string {
	marshal, err := json.Marshal(s)
	if err != nil {
		return ""
	}
	return string(marshal)
}
