package jflat

import (
	"encoding/json"
	"reflect"
)

var _ Flatten = (*mustFlattenList)(nil)

func MustFlatten(i ...interface{}) Flatten {
	list := mustFlattenList(i)
	return &list
}

type mustFlattenList []interface{}

func (list mustFlattenList) MarshalJSON() ([]byte, error) {
	var result []byte
	var errSet ErrorSet
	for i := range list {
		src := list[i]
		data, err := json.Marshal(src)
		if err != nil {
			errSet.Err = append(errSet.Err, &MustFlattenError{
				Type:  reflect.TypeOf(src),
				Index: i,
				Cause: err,
			})
		}

		if len(result) > 0 {
			data[0] = ','
			result = append(result[:len(result)-1], data...)
		} else {
			result = data
		}
	}

	var err error
	if len(errSet.Err) > 0 {
		err = &errSet
	}

	return result, err
}

func (list *mustFlattenList) UnmarshalJSON(data []byte) error {
	var errSet ErrorSet
	for i := range *list {
		dst := (*list)[i]
		err := json.Unmarshal(data, dst)
		if err != nil {
			errSet.Err = append(errSet.Err, &MustFlattenError{
				Type:  reflect.TypeOf(dst),
				Index: i,
				Cause: err,
			})
		}
	}

	if len(errSet.Err) > 0 {
		return &errSet
	}

	return nil
}

func (list *mustFlattenList) Add(i interface{}) error {
	*list = append(*list, i)
	return nil
}
