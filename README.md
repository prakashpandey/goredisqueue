# Redis-Queue

A lightweight Go library to use Redis as a message queue. Supports structured and primitive data types using customizable serialization.

## Features

- Blocking Enqueue / Dequeue
- Generic Payload interface
- Built-in support for structs and primitives
- Minimal API surface with helper constructors

## Install 

```sh
go get github.com/prakashpandey/goredisqueue
```

## Examples

Queue Initialization

```go
queue := goredisqueue.NewWithOptions(goredisqueue.Options{
    RedisAddr: "localhost:6379",
    QueueName: "queue-1",
    Timeout:   3 * time.Second,
})
```

Enqueue a Struct:

```go
type Event struct {
    ID string `json:"id"`
}
e := Event{ID: "abc123"}
_ = queue.Enqueue(ctx, goredisqueue.NewPayloadFromValue(e))
```

Dequeue a Struct

```go
var out Event
_ = queue.Dequeue(ctx, goredisqueue.NewPayloadFromPtr(&out))
```

Enqueue/Dequeue a String

```go
_ = queue.Enqueue(ctx, goredisqueue.NewPayloadFromValue("hello"))

var msg string
_ = queue.Dequeue(ctx, goredisqueue.NewPayloadFromPtr(&msg))
```

## Testing 

Run: `go test ./...`

## Helpers

Interfaces

```go
type Payload interface {
    Marshal() ([]byte, error)
    Unmarshal([]byte) error
}
```

Convert any supported type to `Payload` type

```go
NewPayloadFromValue(value T) *PrimitivePayload[T]
NewPayloadFromPtr(ptr *T) *PrimitivePayload[T]
```

Supported types: Any type tha can be supported by `json.Marshal/Unmarshal` can be used.