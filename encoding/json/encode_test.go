package json_test

import (
	"testing"

	"github.com/dploop/gox/encoding/json"
	"github.com/stretchr/testify/assert"
)

func TestUnsafeMarshal(t *testing.T) {
	v := map[string]interface{}{"name": "Alice", "age": 9}
	data := string(json.UnsafeMarshal(v))
	assert.JSONEq(t, data, `{"name": "Alice", "age": 9}`)
	v = map[string]interface{}{"func": json.UnsafeMarshal}
	data = string(json.UnsafeMarshal(v))
	assert.Empty(t, data)
}
