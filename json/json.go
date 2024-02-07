package json

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// GenericUnmarshaller will unmarshall []byte to the provided dataType
// and return errors if any
func GenericUnmarshaller(in []byte, out interface{}) error {
	val := reflect.ValueOf(out)

	if val.Kind() != reflect.Pointer {
		return fmt.Errorf("destination is not pointer")
	}

	if in == nil {
		val.Elem().Set(reflect.Zero(val.Elem().Type()))
		return nil
	}

	return json.Unmarshal(in, out)
}
