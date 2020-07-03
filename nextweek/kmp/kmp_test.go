package kmp

import (
	"testing"
)

func TestKmp(t *testing.T) {
	checkMatch(t)
	checkNotMatch(t)
}

func checkMatch(t *testing.T) {
	test := [][2]string{
		{"abc", "abc"},
		{"aa=", "afafaaa=fafaaff"},
		{"aa", "bbbafaa"},
		{"afafafa", "fafafafafafaafa"},
	}
	for _, m := range test {
		p, s := m[0], m[1]
		kmp := NewKmp(p)
		index := kmp.Search(s)
		if s[index:index+len(p)] != p {
			t.Error("kmp no correct")
		}
	}

}

func checkNotMatch(t *testing.T) {
	p := "abc="
	s := "aaaaaabc"

	kmp := NewKmp(p)

	if kmp.Search(s) != -1 {
		t.Error("kmp no correct")
	}
}
