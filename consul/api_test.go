package consulapi

import (
	"testing"
)

func TestPutKey(t *testing.T) {
	value := []byte("testingvalue")
	key := "consulfstest/testingkey"
	err := Upload(key, value)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log("Put success")
	}
}
