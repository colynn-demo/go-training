package room01

import "testing"

func TestCanUseItToSleep(t *testing.T) {
	if ans := CanUseItToSleep(1); ans != true {
		t.Error("func error, when int small than 3")
	}
	if ans := CanUseItToSleep(19); ans != false {
		t.Error("func error")
	}
}
