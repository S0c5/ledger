package core

import "encoding/json"

type Script struct {
	Plain string                     `json:"plain"`
	Vars  map[string]json.RawMessage `json:"vars"`
}
