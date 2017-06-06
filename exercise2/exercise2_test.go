package exercise2

import "testing"

func TestFindBob(t *testing.T) {
	cases := []struct {
		input    string
		expected int
	}{
		{"azcbobobegghakl", 2},
		{"zaxbffbpbgwgjmn", 0},
		{"bobwccifwdperkj", 1},
		{"xxkwhdpmkywcbob", 1},
		{"bobizfcpevqcbob", 2},
		{"lzpbobobobwxfoj", 3},
		{"bobobihiksbobob", 4},
	}

	for _, c := range cases {
		result := FindBob(c.input)
		if result != c.expected {
			t.Errorf("expect %d from FindBob(%q), got %d", c.expected, c.input, result)
		}
	}
}
