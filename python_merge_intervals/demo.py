from python_merge_intervals.merge_intervals import Interval, merge_intervals

"""
A demo program for merging intervals
Please run `python ./demo.py`
"""

sample_input = [Interval(25, 30),
                Interval(2, 19),
                Interval(14, 23),
                Interval(4, 8)]


def main():
    print(merge_intervals(sample_input))


if __name__ == '__main__':
    main()
