package events

import (
	"sync"
)

func newBus() *bus {
	var busInstance = new(bus)
	busInstance.mutex = &sync.RWMutex{}
	busInstance.subscriberChannels = make(map[Topic][]Handler)
	return busInstance
}

func (b *bus) pub(event Event) (int, <-chan Result) {
	if b.topicHasNoSubscribers(event.Topic) {
		return -1, nil
	}

	return b.notifyAll(event)
}

func (b *bus) sub(topic Topic, handler Handler) (unsubscribe func()) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	b.subscriberChannels[topic] = append(b.subscriberChannels[topic], handler)
	return func() {
		b.mutex.Lock()
		defer b.mutex.Unlock()
		delete(b.subscriberChannels, topic)
	}
}

type bus struct {
	subscriberChannels map[Topic][]Handler
	mutex              *sync.RWMutex
}

func (b *bus) notifyAll(event Event) (int, <-chan Result) {
	b.mutex.RLock()
	defer b.mutex.RUnlock()
	subs := b.subscriberChannels[event.Topic]
	c := make(chan Result, len(subs))
	for _, handler := range subs {
		go func(h Handler) {
			if h == nil {
				return
			}
			c <- h(event)
		}(handler)
	}
	return len(subs), c
}

func (b *bus) topicHasNoSubscribers(topic Topic) bool {
	b.mutex.RLock()
	defer b.mutex.RUnlock()
	return len(b.subscriberChannels[topic]) == 0
}
