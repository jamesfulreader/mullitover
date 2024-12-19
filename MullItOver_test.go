package mullitover

import (
	"reflect"
	"testing"
)

func TestPatternMatch(t *testing.T) {
	testCases := []struct {
		input    string
		expected [][]string
	}{
		{
			"xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))",
			[][]string{{"mul(2,4)", "2", "4"}, {"mul(5,5)", "5", "5"}, {"mul(11,8)", "11", "8"}, {"mul(8,5)", "8", "5"}},
		},
		{
			"mul(123,456)",
			[][]string{{"mul(123,456)", "123", "456"}},
		},
		{
			"mul[12,34]2_so_^mux(123,890)mul(12,34)",
			[][]string{{"mul(12,34)", "12", "34"}},
		},
		{
			"mul(1,2)mul(3,4)",
			[][]string{{"mul(1,2)", "1", "2"}, {"mul(3,4)", "3", "4"}},
		},
	}

	for i, tc := range testCases {
		result := PatternMatch(tc.input)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("Test case %d: Expected %v but got %v", i, tc.expected, result)
		}
	}
}
