package main

import (
	mi "../merge_intervals"
	"encoding/json"
	"fmt"
)

func main() {
	intervals := mi.Intervals{
			{Start: 25, End: 30},
			{Start: 2, End: 19},
			{Start: 14, End: 23},
			{Start: 4, End: 8},
	}
	output := mi.MergeIntervals(intervals)
	jsonOutput, _ := json.Marshal(output)
	fmt.Printf("Output: %s", string(jsonOutput))
}

