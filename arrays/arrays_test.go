package arrays

import (
	"reflect"
	"testing"
)

var errorString = "Expected %v, got %v"

func TestContains(t *testing.T) {
	if c := Contains([]interface{}{1, 2, 3}, 2); !c {
		t.Errorf(errorString, true, c)
	}

	if c := Contains([]interface{}{1, 2, 3}, 5); c {
		t.Errorf(errorString, false, c)
	}
}

func TestIndex(t *testing.T) {
	if i := Index([]interface{}{1, 2, 3}, 2); i != 1 {
		t.Errorf(errorString, 1, i)
	}

	if i := Index([]interface{}{1, 2, 3}, 5); i != -1 {
		t.Errorf(errorString, -1, i)
	}
}

func TestRemove(t *testing.T) {
	a := []interface{}{1, 2, 3}
	b := Remove(a, 1, 1)
	c := []interface{}{1, 3}

	if !reflect.DeepEqual(b, c) {
		t.Errorf(errorString, c, b)
	}

	a = []interface{}{1, 2, 3}
	b = Remove(a, 1, -1)
	c = []interface{}{1}

	if !reflect.DeepEqual(b, c) {
		t.Errorf(errorString, c, b)
	}
}
