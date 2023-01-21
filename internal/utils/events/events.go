// Package events provides a simple pub-sub interface for internal events.
package events

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

// PubWithPostfix generates a new event. Each subscriber for that baseTopic + extraData will receive it.
// it returns:
//  * number of subcribers, if there is no subscriber -1 is returned
//  * channel where consume result for each subscriber
func PubWithPostfix(baseTopic string, extraData string, payload Payload) (int, <-chan Result) {
	topic := fmt.Sprintf("%s-%s", baseTopic, extraData)
	return Pub(Topic(topic), payload)
}

// Pub generates a new event. Each subscriber for that topic will receive it.
// it returns:
//  * number of subcribers, if there is no subscriber -1 is returned
//  * channel where consume result for each subscriber
func Pub(topic Topic, payload Payload) (int, <-chan Result) {
	/*log.WithFields(logrus.Fields{
		"topic":   topic,
		"payload": fmt.Sprintf("%T", payload),
	}).Debugf("new message")
	*/
	event := Event{topic, payload}
	return defaultBus.pub(event)
}

// SubWithPostfix registers a handler to baseTopic + extraData. Everytime an event for that topic is
// published, the handler will be called asynchronously.
func SubWithPostfix(baseTopic string, extraData string, handler Handler) (unsubscribe func()) {
	topic := fmt.Sprintf("%s-%s", baseTopic, extraData)
	return Sub(Topic(topic), handler)
}

// Sub registers a handler to a topic. Everytime an event for that topic is
// published, the handler will be called asynchronously.
func Sub(topic Topic, handler Handler) (unsubscribe func()) {
	log.WithFields(logrus.Fields{
		"topic": topic,
	}).
		Debugf("new subscriber")

	return defaultBus.sub(topic, handler)
}

var defaultBus = newBus()

var log = logrus.WithFields(logrus.Fields{
	"module": "events",
})

//WaitResults is an helper to handle chan Result it returns after receiving all Result or
//after a timeout error
func WaitResults(nSub int, result <-chan Result, timeout time.Duration) ([]Result, error) {
	results := []Result{}
	for {
		if nSub < 1 {
			break
		}
		select {
		case res := <-result:
			results = append(results, res)
			nSub--
		case <-time.After(timeout):
			return results, fmt.Errorf("Timeout error after %s", timeout)
		}
	}
	return results, nil
}
