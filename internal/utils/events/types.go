package events

import (
	"fmt"
)

// Event represents something that happened in the app.
type Event struct {
	Topic   Topic
	Payload Payload
}

// Topic is a unique identifier for the kind of an event.
type Topic string

// Payload is a generic value attached to an event. Can be nil.
type Payload interface{}

// Result is a generic response to an handler.
type Result struct {
	Error error
	Value interface{}
}

// Handler is a callback executed everytime a message arrives on one of the
// subscribed topics.
// The handler MUST BE THREAD SAFE.
type Handler func(event Event) Result

//OK empty response without error
func OK() Result {
	return Result{nil, nil}
}

//Errorf base error Response
func Errorf(format string, args ...interface{}) Result {
	return Error(fmt.Errorf(format, args...))
}

//Error base error Response
func Error(err error) Result {
	return Result{err, nil}
}
