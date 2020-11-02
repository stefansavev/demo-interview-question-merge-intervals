# About

This is a simple interview problem. Demonstrates
problem-solving, code-formatting, testing and Docker.

It comes in two flavors:

- Python
- GoLang (to be committed later)


# Getting Started

Clone the repository:

```
git clone https://github.com/stefansavev/demo-interview-question-merge-intervals.git
```

Linux/Mac users simply run

```
cd demo-interview-question-merge-intervals/python_merge_intervals/
export PYTHONPATH="$(cd .. && pwd)"
python ./demo.py
```

Windows users open a `cmd.exe` window and set the `PYTHONPATH` with
```
set PYTHONPATH="C:\...\demo-interview-question-merge-intervals/python_merge_intervals/"
```
Then call `python ./demo.py`

Expected output:

```
[(2, 23), (25, 30)]
```

# Running the tests

```
cd python_merge_intervals/
pytest .
```

to get an experience how the code works.

Expected output:

```
platform linux -- Python 3.8.3, pytest-6.0.2, py-1.9.0, pluggy-0.13.1
rootdir: /home/stefan/demo-interview-question-merge-intervals/python_merge_intervals
collected 5 items                                                                                                                                     

test_merge_intervals.py .....                                                                                                                   [100%]

================================================================== 5 passed in 0.01s ==================================================================
```

# Using Docker
```
cd demo-interview-question-merge-intervals/
docker build -f Dockerfile.python . -t stefansavev/python_merge_intervals
```

To run:

```
docker run -it stefansavev/python_merge_intervals:latest
```

# Time to Completion

- 10 min for creating repository, setting up ssh keys, cloning
- 20 min for solving the merge intervals
- 20 min for code re-organization and pytest
- 40 min for code formatting, writing comments and README
- 20 min for Docker
- 20 min for setting up Go env. (upcoming)
- 20 min for Go (upcoming)
