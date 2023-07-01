package madtskey

import (
	"log"
	"testing"
)

func TestCreateAPIKey(t *testing.T) {
	description := "my test key"
	tags := []string{"tag:mad-ts-key"}
	key, err := CreateAPIKey(300, description, tags)
	if err != nil {
		log.Fatal(err)
	}
	if key.Description != description {
		t.Errorf("wanted %s got %s", description, key.Description)
	}
	var selectedIndex int
	for i, tag := range key.Capabilities.Devices.Create.Tags {
		if tag == tags[0] {
			selectedIndex = i
		}
	}
	if selectedIndex == -1 {
		t.Error("wanted non negative got -1")
	}
}
