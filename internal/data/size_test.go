package data

import "testing"

func TestIsNonNull(t *testing.T) {

	testCases := []struct {
		name     string
		width    int
		height   int
		expected bool
	}{
		{"has width and heigth", 100, 100, true},
		{"has width", 100, 0, true},
		{"has height", 0, 100, true},
		{"is null", 0, 0, false},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			s := Size{
				Width:  testCase.width,
				Height: testCase.height,
			}

			if s.IsNonNull() != testCase.expected {
				t.Errorf("expected %v got %v; s.IsNonNull()", testCase.expected, s.IsNonNull())
			}
		})
	}
}
