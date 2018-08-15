package endpoint

import "testing"

func TestPhase2State(t *testing.T) {
	a := phase2state(PhaseADD)
	if a != "1" {
		t.Error("phaseadd err")
	}
	a = phase2state(PhaseUPDATE)
	if a != "1" {
		t.Error("phaseupdate err")
	}
	a = phase2state(PhaseDEL)
	if a != "0" {
		t.Error("phasedel err")
	}
}
