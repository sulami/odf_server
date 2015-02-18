package server

import "testing"

func TestOnlineOffline(t *testing.T) {
	s := GameServer{1339, false, make(chan bool)}

	err := s.Listen()
	if err != nil {
		t.Errorf("Server could not come online: %s", err.Error())
	}

	err = s.Listen()
	if err == nil {
		t.Errorf("Server did not return an error on duplicated start")
	}

	err = s.StopListening()
	if err != nil {
		t.Errorf("Server could go offline: %s", err.Error())
	}

	err = s.StopListening()
	if err == nil {
		t.Errorf("Server did not return an error on duplicated stop")
	}
}

