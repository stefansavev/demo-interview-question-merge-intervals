from typing import List


class Interval:
    def __init__(self, start: int, end: int):
        self.start = start
        self.end = end

    def __repr__(self):
        return repr((self.start, self.end))

    def __eq__(self, other: 'Interval'):
        if not isinstance(other, Interval):
            return False
        return self.start == other.start and \
               self.end == self.end


def sorted_intervals_overlap(first_interval: Interval, second_interval: Interval):
    """
    Tests whether two sorted intervals overlap.
    The first_interval must have its start before the second one.
    :param first_interval: the first interval. Must have its start before the second one
    :param second_interval: the second interval. Must have its start after the second one
    :return: true if the second interval start is before the first end
    """
    assert first_interval.start <= second_interval.start
    return second_interval.start <= first_interval.end


def merge_sorted_intervals(first_interval: Interval, second_interval: Interval):
    """
    Merges two sorted intervals which overlap. The intervals must overlap as
    detected via `sorted_intervals_overlap` function. Otherwise an assertion is raised.
    :param first_interval: the first interval.
    :param second_interval: the second interval.
    :return: A new interval which represents the `union` of the first and second interval
    """

    assert sorted_intervals_overlap(first_interval, second_interval)

    start = min(first_interval.start, second_interval.start)
    end = max(first_interval.end, second_interval.end)
    return Interval(start, end)


def merge_intervals(intervals: List[Interval]) -> List[Interval]:
    """
    Merge a list of intervals as per the 'merge' specification.
    Please see test_merge_intervals for examples.
    The input is not mutated.
    :param intervals:
    :return: a list of merged intervals
    """
    if len(intervals) == 0:
        return []

    sorted_intervals = sorted(intervals, key=lambda interval: interval.start)

    current = sorted_intervals[0]
    merged_intervals = []
    for interval in sorted_intervals[1:]:
        if sorted_intervals_overlap(current, interval):
            current = merge_sorted_intervals(current, interval)
        else:
            merged_intervals.append(current)
            current = interval
    merged_intervals.append(current)

    return merged_intervals
