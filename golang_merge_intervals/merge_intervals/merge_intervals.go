package merge_intervals

import (
	"sort"
)

type Interval struct {
	Start int
	End   int
}

type Intervals []Interval

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func copyIntervals(intervals Intervals) []Interval {
	target := make(Intervals, len(intervals))
	copy(target, intervals)
	return target
}

func tryMerge(firstInterval Interval, secondInterval Interval) (bool, Interval) {
	if secondInterval.Start <= firstInterval.End {
		return true, Interval{
			Start: min(firstInterval.Start, secondInterval.Start),
			End:   max(firstInterval.End, secondInterval.End),
		}
	}
	return false, Interval{}
}

type byIntervalStart []Interval

func (a byIntervalStart) Len() int           { return len(a) }
func (a byIntervalStart) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byIntervalStart) Less(i, j int) bool { return a[i].Start < a[j].Start }

func MergeIntervals(intervals Intervals) Intervals {
	var output []Interval = make(Intervals, 0)
	if len(intervals) == 0 {
		return output
	}

	sortedIntervals := copyIntervals(intervals)
	sort.Sort(byIntervalStart(sortedIntervals))

	var current = sortedIntervals[0]

	for _, interval := range sortedIntervals {
		ok, mergedInterval := tryMerge(current, interval)
		if ok {
			current = mergedInterval
		} else {
			output = append(output, current)
			current = interval
		}
	}

	output = append(output, current)

	return output
}
