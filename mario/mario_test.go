package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO update this function to match the function name or add addtional
// test cases.
func TestFn(t *testing.T) {
	testCases := map[string]struct {
		input    int
		expected string
	}{
		"zero": {
			input:    0,
			expected: "",
		},
		"one": {
			input:    1,
			expected: "#",
		},
		"two": {
			input:    2,
			expected: "##",
		},
		"ten": {
			input:    10,
			expected: "##########",
		},
		"negative one": {
			input:    -1,
			expected: "",
		},
		"too big": {
			input:    100,
			expected: "",
		},
	}

	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			rt := fn(testCase.input)
			assert.Equal(t, testCase.expected, rt)
		})
	}
}
