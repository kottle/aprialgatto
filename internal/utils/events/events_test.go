package events

import (
	"os"
	"sync"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func init() {
	if os.Getenv("DEBUG") == "true" {
		logrus.SetLevel(logrus.DebugLevel)
	}
}

// MessageCounter is a simple subscriber for events that keeps a count of received events.
type MessageCounter struct {
	sync.Mutex
	count int
}

func (s *MessageCounter) Handler(event Event) {
	s.Lock()
	s.count++
	s.Unlock()
}

func TestSimplePub(t *testing.T) {
	assert := assert.New(t)

	topic := t.Name()

	var counter MessageCounter
	sub(topic, counter.Handler)
	pubEmpty(topic)

	assert.Equal(1, counter.count, "message was not received")
}

func TestPubWithoutSubscribers(t *testing.T) {
	topic := t.Name()
	pubEmpty(topic)
}

func TestPubMany(t *testing.T) {
	assert := assert.New(t)

	msgCount := 10000
	topic := t.Name()

	var counter MessageCounter
	sub(topic, counter.Handler)

	for i := 0; i < msgCount; i++ {
		pubEmpty(topic)
	}

	assert.Equal(msgCount, counter.count, "some messages were lost")
}

func TestSubManySameTopic(t *testing.T) {
	assert := assert.New(t)

	topic := t.Name()

	// generate a list of subscribers and sub them to the topic
	subCount := 10000
	subscribers := make([]*MessageCounter, subCount)
	for i := range subscribers {
		subscribers[i] = &MessageCounter{}
		sub(topic, subscribers[i].Handler)
	}

	// pub an event on the topic
	pubEmpty(topic)

	// sleep to ensure all the handlers are actually executed
	time.Sleep(100 * time.Millisecond)

	// check if everyone received it
	for _, s := range subscribers {
		assert.Equal(1, s.count, "one subscriber did not receive the message")
	}
}

func TestUnsubscribe(t *testing.T) {
	assert := assert.New(t)
	topic := t.Name()

	var counter MessageCounter
	unsubscribe := sub(topic, counter.Handler)
	unsubscribe()

	pubEmpty(topic)

	assert.Equal(0, counter.count, "message received after unsubscribe")
}

func pubEmpty(topic string) {
	Pub(Topic(topic), nil)

	// this sleep guarantees that handlers for this event have been executed
	time.Sleep(10 * time.Microsecond)
}

func sub(topic string, handler Handler) func() {
	return Sub(Topic(topic), handler)
}
