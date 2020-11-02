import pytest

from typing import List

from python_merge_intervals.merge_intervals import Interval, merge_intervals


class MergeIntervalsTestCase:
    def __init__(self, input_intervals: List[Interval], output_intervals: List[Interval]):
        self.input_intervals = input_intervals
        self.output_intervals = output_intervals


test_cases = [
    MergeIntervalsTestCase([], []),
    MergeIntervalsTestCase([Interval(1, 2)], [Interval(1, 2)]),
    MergeIntervalsTestCase([Interval(1, 2), Interval(3, 4)], [Interval(1, 2), Interval(3, 4)]),
    MergeIntervalsTestCase([Interval(1, 2), Interval(2, 3)], [Interval(1, 3)]),
    MergeIntervalsTestCase([Interval(25, 30), Interval(2, 19),
                            Interval(14, 23), Interval(4, 8)],
                           [Interval(2, 23), Interval(25, 30)])
]


@pytest.mark.parametrize("test_case", test_cases)
def test_eval(test_case):
    actual_output = merge_intervals(test_case.input_intervals)
    assert actual_output == test_case.output_intervals
