package uniq_test

import (
	"reflect"
	"testing"
	"uniq/uniq"
)

func TestReadFromFile(t *testing.T) {
	filename := "../unit_test.txt"
	want := []byte("Walter O'Brien\n")
	if got := uniq.ReadFromFile(filename); !reflect.DeepEqual(got, want) {
		t.Errorf("ReadFromFile() = %v, want: %v", got, want)
	}
}
