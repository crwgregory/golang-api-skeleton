package utils

import (
	"testing"
	"fmt"
)

func TestGetJWT(t *testing.T) {
	expectedGroups := []int{2}
	jwt, err := GetJWT(1, expectedGroups)

	if err != nil {
		t.Error(err)
	}

	payload, err := ParseJWT(jwt)

	if err != nil {
		t.Fatal(err)
	}

	if payload.UserID != 1 {
		t.Error("user_id does not match")
	}

	groups := payload.UserGroups
	if len(groups) != len(expectedGroups) {
		t.Error("groups do not match")
	}

	for i, g := range groups {
		if expectedGroups[i] != g {
			t.Error("groups do not match")
		}
	}
	fmt.Println(jwt, payload)
}
