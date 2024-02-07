package json_test

import (
	"testing"

	"github.com/Secure-Petal/tools/json"
	"github.com/stretchr/testify/assert"
)

func TestGenericUnmarshaller(t *testing.T) {
	type User struct {
		Name       string `json:"name"`
		Age        int    `json:"age"`
		IsVerified bool   `json:"is_verified"`
	}

	t.Run("valid json to user struct", func(t *testing.T) {
		input := []byte(`{"name":"John Doe", "age":20, "is_verified":true}`)
		output := User{}

		err := json.GenericUnmarshaller(input, &output)
		assert.Nil(t, err)

		assert.Equal(t, "John Doe", output.Name)
		assert.Equal(t, 20, output.Age)
		assert.Equal(t, true, output.IsVerified)
	})

	t.Run("empty slice of byte", func(t *testing.T) {
		input := make([]byte, 0)
		output := User{}

		err := json.GenericUnmarshaller(input, &output)
		assert.Equal(t, "unexpected end of JSON input", err.Error())
	})

	t.Run("destination is not pointer", func(t *testing.T) {
		input := []byte(`{"first_name":"John", "age_num":20, "is_active":true}`)
		output := User{}

		err := json.GenericUnmarshaller(input, output)
		assert.Equal(t, "destination is not pointer", err.Error())
	})

	t.Run("different fields in JSON, no error", func(t *testing.T) {
		input := []byte(`{"first_name":"John", "age_num":20, "is_active":true}`)
		output := User{}

		err := json.GenericUnmarshaller(input, &output)
		assert.Nil(t, err)
	})
}
