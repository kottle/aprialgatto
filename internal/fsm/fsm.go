package fsm

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type State string
type Event string

type Action struct {
	Destination State
}

type Callback func(Event, State)

type Transition map[Event]Action
type StateMap map[State]Transition
type CallbackMap map[State]Callback

type FiniteStateMachine struct {
	Initial     State
	current     State
	StateMap    StateMap
	CallbackMap CallbackMap
}

type IFiniteStateMachine interface {
	Current() State
	Transition(event Event) error
}

func (fsm *FiniteStateMachine) Current() State {
	if fsm.current == "" {
		return fsm.Initial
	}
	return fsm.current
}

func (fsm *FiniteStateMachine) Transition(event Event) error {
	action := fsm.StateMap[fsm.Current()][event]
	if fmt.Sprint(action) != fmt.Sprint(Action{}) {
		prevState := fsm.current
		fsm.current = action.Destination
		logrus.Infof("STATE[%s->%s] with event %s ", prevState, action.Destination, event)
		cbk := fsm.CallbackMap[fsm.current]
		if cbk != nil {
			cbk(event, prevState)

		}
		return nil
	}
	return fmt.Errorf("transition invalid")
}

func (fsm *FiniteStateMachine) Start() {
	fsm.CallbackMap[fsm.Current()]("", fsm.Initial)
}
