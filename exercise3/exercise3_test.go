package exercise3

import "testing"

func TestLongestAlphabeticalOrder(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{"azcbobobegghakl", "beggh"},
		{"abcbcd", "abc"},
	}

	for _, c := range cases {
		result := LongestAlphabeticalOrder(c.input)
		if result != c.expected {
			t.Errorf("expect %q from LongestAlphabeticalOrder(%q), got %q", c.expected, c.input, result)
		}
	}
}
