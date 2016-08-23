package Relay

import (
	"testing"
	"time"
)

func TestRelay(t *testing.T) {
	t.Log("Start Test Relay")

	r, err := New()

	if err != nil {
		t.Error("Error Start relay ", err)
	}

	t.Log("Off the relay")
	r.pOff()

	time.Sleep(2 * time.Second)

	t.Log("On the relay")
	r.pOn()

	time.Sleep(2 * time.Second)

	t.Log("Off the relay")
	r.pOff()

	time.Sleep(2 * time.Second)

	t.Log("On the relay")
	r.pOn()

	time.Sleep(2 * time.Second)

}
