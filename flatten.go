package jflat

import (
	"encoding/json"
)

type Flatten interface {
	json.Marshaler
	json.Unmarshaler
	Add(interface{}) error
}
