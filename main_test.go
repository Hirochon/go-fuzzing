package main

import (
	"fmt"
	"testing"
)

func FuzzReverse(f *testing.F) {
	testcase := []string{"Hiroto", "Hello", "Check it out"}
	for _, tc := range testcase {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, s string) {
		fmt.Println(s)
		t.Log(s)
		rs := reverse(s)
		if s != reverse(rs) {
			t.Errorf("%s: %s", rs, s)
		}
	})
}

func FuzzInt(f *testing.F) {
	f.Fuzz(func(t *testing.T, i int) {
		fmt.Println(i)
		t.Log(i)
	})
}
