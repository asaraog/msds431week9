package main

import (
	"testing"
)

func TestProcessInput(t *testing.T) {
	userinput := ProcessInput("break")
	if string([]rune(userinput)[0]) != "'" {
		t.Errorf("Error in Processing input. No ' as expected")
	}
	if string([]rune(userinput)[len(userinput)-1]) != "'" {
		t.Errorf("Error in Processing input. No ' as expected")
	}

}
