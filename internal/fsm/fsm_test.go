package fsm

import (
	"fmt"
	"testing"
)

var ifsm FiniteStateMachine

func init() {
	ifsm = FiniteStateMachine{
		Initial: "wait",
		StateMap: StateMap{
			"wait": Transition{
				"proximity_obj": Action{
					Destination: "open_camera",
				},
			},
			"open_camera": Transition{
				"detected": Action{
					Destination: "open_gate",
				},
				"undetected": Action{
					Destination: "wait",
				},
			},
			"open_gate": Transition{
				"cat_gone": Action{
					Destination: "wait",
				},
			},
		},
	}
}

func TestFsm(t *testing.T) {
	for _, event := range []Event{"proximity_obj", "undetected", "proximity_obj", "detected", "cat_gone"} {
		curr := ifsm.Current()
		err := ifsm.Transition(event)
		fmt.Printf("%s -> %s -> %s\n", curr, event, ifsm.Current())
		if err != nil {
			fmt.Println(err)
		}
	}
}
