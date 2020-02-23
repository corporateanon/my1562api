package client

import (
	"testing"
)

func TestGetStatus(t *testing.T) {
	res, err := GetStatus(2334, "4")
	if err != nil {
		t.Fatal(err)
	}
	if res.HasMessage {

		t.Logf("Title:       %s\n", res.Title)
		t.Logf("Description: %s\n", res.Description)
	} else {
		t.Log("No incidents\n")
	}
}
