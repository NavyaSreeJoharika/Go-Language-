package main

import "testing"

func TestAdd(t *testing.T) {

	//arrange
	l, r := 1, 2
	expect := 3

	//act
	got := Add(l, r)

	//assert
	if expect != got {
		t.Errorf("failed to add %v and %v. Got %v,expected %v \n", l, r, got, expect)
	}
}
