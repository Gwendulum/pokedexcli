package main
import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{	
			input:    "My name Jeff  ",
			expected: []string{"my", "name", "jeff"},
		},
				{	
			input:    "Punctuation: Commas like this one, are not separated  ",
			expected: []string{"punctuation:", "commas", "like", "this", "one,", "are", "not", "separated"},
		},
}	

	for _, c := range cases {
		actual := CleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("FAILED: expected slice length: %v Actual: %v", len(c.expected), len(actual))
			t.Fail()
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("FAILED: expected: %v, actual: %v", expectedWord, word)
				t.Fail()
			}

		}
}

}