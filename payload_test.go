package redisqueue

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPrimitivePayload_String(t *testing.T) {
	original := "hello"
	p := NewPayloadFromValue(original)

	data, err := p.Marshal()
	require.NoError(t, err)

	result := ""
	up := NewPayloadFromPtr(&result)
	require.NoError(t, up.Unmarshal(data))
	assert.Equal(t, original, result)
}

func TestPrimitivePayload_Int(t *testing.T) {
	original := 123
	p := NewPayloadFromValue(original)

	data, err := p.Marshal()
	require.NoError(t, err)

	result := 0
	up := NewPayloadFromPtr(&result)
	require.NoError(t, up.Unmarshal(data))
	assert.Equal(t, original, result)
}

func TestPrimitivePayload_Bool(t *testing.T) {
	original := true
	p := NewPayloadFromValue(original)

	data, err := p.Marshal()
	require.NoError(t, err)

	result := false
	up := NewPayloadFromPtr(&result)
	require.NoError(t, up.Unmarshal(data))
	assert.Equal(t, original, result)
}

func TestPrimitivePayload_InvalidJSON(t *testing.T) {
	result := 0
	p := NewPayloadFromPtr(&result)
	err := p.Unmarshal([]byte(`"not a number"`))
	assert.Error(t, err)
}

func TestPrimitivePayload_CustomStruct(t *testing.T) {
	type MyStruct struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	original := MyStruct{Name: "Alice", Age: 30}
	p := NewPayloadFromValue(original)

	data, err := p.Marshal()
	require.NoError(t, err)

	result := MyStruct{}
	up := NewPayloadFromPtr(&result)
	require.NoError(t, up.Unmarshal(data))
	assert.Equal(t, original, result)
}

func TestPrimitivePayload_JSONRoundTrip(t *testing.T) {
	value := map[string]any{"key": "value", "n": 42}
	p := NewPayloadFromValue(value)

	data, err := p.Marshal()
	require.NoError(t, err)

	var result map[string]any
	up := NewPayloadFromPtr(&result)
	require.NoError(t, up.Unmarshal(data))

	// Normalize both maps using JSON.Marshal to compare them.
	expectedJSON, err := json.Marshal(value)
	require.NoError(t, err)
	resultJSON, err := json.Marshal(result)
	require.NoError(t, err)

	assert.JSONEq(t, string(expectedJSON), string(resultJSON))
}
