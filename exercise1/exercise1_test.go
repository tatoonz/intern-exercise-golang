package exercise1

import "testing"

func TestVowelCount(t *testing.T) {
	cases := []struct {
		input    string
		expected int
	}{
		{"azcbobobegghakl", 5},
		{"kupfoczkgwgdpwq", 2},
		{"uwpaqbzwghdqxsu", 3},
		{"jkoszrzcfhjlqya", 2},
		{"svdreyqueauztqo", 6},
		{"qwbmhkkgapxaerb", 3},
		{"hzuobtptfixdnwr", 3},
		{"djhfuxjwslcywcp", 1},
		{"consvkqzaayjvgn", 3},
		{"ysjjoscfuwhjkdr", 2},
		{"pqpbpemtbdyhydn", 1},
	}

	for _, c := range cases {
		result := VowelCount(c.input)
		if result != c.expected {
			t.Errorf("expect %d from VowelCount(%q), got %d", c.expected, c.input, result)
		}
	}
}
