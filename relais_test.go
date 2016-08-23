package Relay

import (
	"fmt"
	"testing"
	"time"
)

func TestRelay(t *testing.T) {

	r, err := New()

	if err != nil {
		t.Error("Error Start relay ", err)
	}

	fmt.Println("Off the relay")
	r.POff()

	time.Sleep(2 * time.Second)

	fmt.Println("On the relay")
	r.POn()

	time.Sleep(2 * time.Second)

	fmt.Println("Off the relay")
	r.POff()

	time.Sleep(2 * time.Second)

	fmt.Println("Off the relay")
	r.POff()

	time.Sleep(2 * time.Second)

	fmt.Println("On the relay")
	r.POn()

	time.Sleep(2 * time.Second)

}
