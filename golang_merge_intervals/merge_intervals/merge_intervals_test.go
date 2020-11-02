package merge_intervals

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

type mergeIntervalsTest struct {
	input  Intervals
	output Intervals
}

var testCases = []mergeIntervalsTest{
	{
		input:  Intervals{},
		output: Intervals{},
	},
	{
		input:  Intervals{{Start: 1, End: 2}, {Start: 2, End: 4}},
		output: Intervals{{Start: 1, End: 4}},
	},
	{
		input:  Intervals{{Start: 1, End: 2}},
		output: Intervals{{Start: 1, End: 2}},
	},
	{
		input:  Intervals{{Start: 1, End: 2}, {Start: 3, End: 4}},
		output: Intervals{{Start: 1, End: 2}, {Start: 3, End: 4}},
	},
	{
		input: Intervals{
			{Start: 25, End: 30},
			{Start: 2, End: 19},
			{Start: 14, End: 23},
			{Start: 4, End: 8}},
		output: Intervals{
			{Start: 2, End: 23},
			{Start: 25, End: 30}},
	},
}

func TestIntervals(t *testing.T) {
	for _, testCase := range testCases {
		result := MergeIntervals(testCase.input)
		expected := testCase.output
		if !cmp.Equal(result, expected) {
			t.Errorf("expected: %v, got: %v", expected, result)
		}
	}
}
