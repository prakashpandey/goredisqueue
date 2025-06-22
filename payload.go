package redisqueue

import "encoding/json"

type Payload interface {
	Marshal() ([]byte, error)
	Unmarshal(data []byte) error
}

// PrimitivePayload wraps a primitive type or struct to make it a Payload.
type PrimitivePayload[T any] struct {
	Value *T
}

// NewPayloadFromPtr wraps an existing pointer in a PrimitivePayload.
func NewPayloadFromPtr[T any](value *T) *PrimitivePayload[T] {
	return &PrimitivePayload[T]{Value: value}
}

// NewPayloadFromValue creates a new PrimitivePayload with a value.
// The value will be wrapped in a pointer.
func NewPayloadFromValue[T any](value T) *PrimitivePayload[T] {
	return &PrimitivePayload[T]{Value: &value}
}

func (p *PrimitivePayload[T]) Marshal() ([]byte, error) {
	return json.Marshal(p.Value)
}

func (p *PrimitivePayload[T]) Unmarshal(data []byte) error {
	return json.Unmarshal(data, p.Value)
}
