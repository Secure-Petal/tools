package tools

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// GenericUnmarshaller will unmarshall []byte to the provided dataType
// and return errors if any
func GenericUnmarshaller(in []byte, out interface{}, errs []error) []error {
	val := reflect.ValueOf(out)

	if val.Kind() != reflect.Pointer {
		errs = append(errs, fmt.Errorf("destination is not pointer"))
		return errs
	}

	if in == nil {
		val.Elem().Set(reflect.Zero(val.Elem().Type()))
		return errs
	}

	err := json.Unmarshal(in, out)
	if err != nil {
		errs = append(errs, err)
	}

	return errs
}
