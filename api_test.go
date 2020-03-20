package client

import (
	"testing"
)

func TestGetStatus(t *testing.T) {
	res, err := GetStatus(2562, "14")
	if err != nil {
		t.Fatal(err)
	}
	if res.HasMessage {
		for _, message := range res.Messages {
			t.Logf("Title:       %s\n", message.Title)
			t.Logf("Description: %s\n", message.Description)
		}
	} else {
		t.Log("No incidents\n")
	}
}
