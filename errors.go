package jflat

import (
	"fmt"
	"reflect"
	"strings"
)

var _ error = (*MustFlattenError)(nil)

type MustFlattenError struct {
	Type  reflect.Type
	Index int
	Cause error
}

func (m *MustFlattenError) Error() string {
	return fmt.Sprintf("type: %s / index: %d / cause: %s", m.Type.String(), m.Index, m.Cause.Error())
}

var _ error = (*ErrorSet)(nil)

type ErrorSet struct {
	Err []error
}

func (e *ErrorSet) Error() string {
	arr := make([]string, len(e.Err))
	for i := range arr {
		arr[i] = e.Err[i].Error()
	}

	return strings.Join(arr, "\n")
}
