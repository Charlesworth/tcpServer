package portManager

import "testing"

func TestNew(t *testing.T) {
	firstPort, lastPort := 10, 12
	testPM := New(firstPort, lastPort)

	if len(testPM.portChan) != 3 {
		t.Error("New gave an incorrect output after New")
	}
}

func TestTakePort(t *testing.T) {
	firstPort, lastPort := 10, 11
	testPM := New(firstPort, lastPort)

	testPort, _ := testPM.TakePort()
	if testPort != firstPort {
		t.Error("TakePort did not return the first assigned port")
	}

	testPort, _ = testPM.TakePort()
	if testPort != firstPort+1 {
		t.Error("TakePort did not return the second assigned port")
	}

	testPort, testErr := testPM.TakePort()
	if testErr == nil { //testPort != 0 && testErr == nil {
		t.Error("TakePort did not return error when out of ports")
	}
}

func TestReturnPort(t *testing.T) {
	firstPort, lastPort := 10, 10
	testPM := New(firstPort, lastPort)

	testPM.TakePort()
	testPM.TakePort()
	testPM.ReturnPort(10)

	testReturnedPort, _ := testPM.TakePort()
	if testReturnedPort != 10 {
		t.Error("ReturnedPort did not replace the port into the port manager")
	}
}
