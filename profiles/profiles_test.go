package profiles

import (
	"reflect"
	"testing"
)

func Test_NewProfiles(t *testing.T) {
	credFile := "../fixtures/credentials"
	prof, err := NewProfiles(credFile)
	if err != nil {
		t.Fatalf("Fail to read file: %v", err)
	}
	list := prof.makeList()
	expected := []string{"profile-1", "profile-2", "profile-3", "profile-4"}

	if !reflect.DeepEqual(expected, list) {
		t.Errorf("\n%v\n%v", expected, list)
	}
}
