package core

import "testing"

func TestAbs(t *testing.T) {
	GetCore().Init()
	GetCore().SendMessage("detectObject", "detect")
	//GetCore().SendMessage("detectObject", "detect_ok")

}
