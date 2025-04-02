package counter

import (
	"testing"

	"github.com/Teerawat36167/PieFireDire/internal/util"
)

func TestCountMeats(t *testing.T) {
	mc := counter.NewMeatCounter()

	tests := []struct {
		name     string
		input    string
		expected map[string]int
	}{
		{
			name:  "Basic test",
			input: "Fatback t-bone t-bone, pastrami .. t-bone. pork, meatloaf jowl enim. Bresaola t-bone.",
			expected: map[string]int{
				"fatback":  1,
				"t-bone":   4,
				"pastrami": 1,
				"pork":     1,
				"meatloaf": 1,
				"jowl":     1,
				"enim":     1,
				"bresaola": 1,
			},
		},
		{
			name:     "Empty string",
			input:    "",
			expected: map[string]int{},
		},
		{
			name:  "Single word",
			input: "beef",
			expected: map[string]int{
				"beef": 1,
			},
		},
		{
			name:  "Mixed case",
			input: "Beef BEEF beef",
			expected: map[string]int{
				"beef": 3,
			},
		},
		{
			name:  "Special characters",
			input: "beef!@#$%^&*()_+{}|:<>?[]\\;',./~`",
			expected: map[string]int{
				"beef": 1,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := mc.CountMeats(test.input)

			if len(result) != len(test.expected) {
				t.Errorf("Expected %d items, got %d", len(test.expected), len(result))
			}

			for word, expectedCount := range test.expected {
				if count, ok := result[word]; !ok || count != expectedCount {
					t.Errorf("For word %q: expected count %d, got %d", word, expectedCount, count)
				}
			}
		})
	}
}
