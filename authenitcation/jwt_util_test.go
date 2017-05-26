package authentication

import (
	"fmt"
	"testing"
)

func TestGetJWT(t *testing.T) {
	userID := 2
	userGroups := []int{2, 3}
	userEntities := []int{2, 3}
	jwt, err := GetJWT(userID, userGroups, userEntities)

	if err != nil {
		t.Error(err)
	}

	payload, err := ParseJWT(jwt)

	if err != nil {
		t.Fatal(err)
	}

	if payload.UserID != userID {
		t.Error("user_id does not match")
	}

	groups := payload.UserGroups
	if len(groups) != len(userGroups) {
		t.Error("groups do not match")
	}
	for i, g := range groups {
		if userGroups[i] != g {
			t.Error("groups do not match")
		}
	}

	entitys := payload.EntityIDs
	if len(entitys) != len(userEntities) {
		t.Error("entites do not match")
	}
	for i, e := range entitys {
		if userEntities[i] != e {
			t.Error("entites do not match")
		}
	}

	fmt.Println(jwt, payload)
}
