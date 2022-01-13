package slice

import (
	"errors"
	"reflect"
)

func Compact(slice interface{}) error {

	if reflect.TypeOf(slice).Kind() != reflect.Slice {
		return errors.New("not a slice")
	}

	return nil

}
