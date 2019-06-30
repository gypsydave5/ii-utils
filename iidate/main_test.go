package main

import "testing"

func TestFormatDate(t *testing.T) {
	line := "1558738697 <yitz> Yeah. But annoying when you are working in one of those and got it a lot"
	want := "[22:58:17] <yitz> Yeah. But annoying when you are working in one of those and got it a lot"

	got := formatDate(line)
	if got != want {
		t.Errorf("\nWanted\t%#v\ngot\t%#v\n", want, got)
	}
}
