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
set PYTHONPATH="C:\...\demo-interview-question-merge-intervals\python_merge_intervals"
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

# Implementation Details (Algorithm)

The following algorithm is used:
1. Sort the intervals by their start time.
2. Loop over the sorted intervals.
    Invariants:
        current is an interval that can grow from the end (but not from the start).
    a. If current is not set, set it to the interval from the loop iteration
    b1. If current is set and the interval from the loop can be merged to current,
        then grow/update current to a merged interval
    b2. If current is set but the loop interval does not overlap with current,
        write current to output.
        Begin a new current interval
3. Make sure to output current interval

# Testing

I use table driven tests with pytest.


# Time to Completion

- 10 min for creating repository, setting up ssh keys, cloning
- 20 min for solving the merge intervals
- 20 min for code re-organization and pytest
- 40 min for code formatting, writing comments and README
- 20 min for Docker

# Extra: Using GoLang

Build:
```
cd demo-interview-question-merge-intervals/
docker build -f Dockerfile.go . -t stefansavev/go_merge_intervals
```

Then run:

```
docker run -it stefansavev/go_merge_intervals:latest
```

# Long story short

Simply run:

```
make
```

Requires Make and Docker.

## Output of make

```
docker build -f Dockerfile.python . -t stefansavev/python_merge_intervals
Sending build context to Docker daemon  200.2kB
Step 1/7 : FROM python:3.9.0-slim-buster
 ---> 8842fc1a4cb7
Step 2/7 : RUN mkdir /app && pip install pytest
 ---> Using cache
 ---> fdf9f40f4d24
Step 3/7 : WORKDIR /app
 ---> Using cache
 ---> ff1c024c5aee
Step 4/7 : COPY python_merge_intervals/ python_merge_intervals/
 ---> Using cache
 ---> 94f63828dd1f
Step 5/7 : ENV PYTHONPATH=/app/
 ---> Using cache
 ---> 09197b42e07a
Step 6/7 : RUN cd python_merge_intervals && pytest .
 ---> Using cache
 ---> 76e77af2bbc1
Step 7/7 : CMD python python_merge_intervals/demo.py
 ---> Using cache
 ---> 254f1406b325
Successfully built 254f1406b325
Successfully tagged stefansavev/python_merge_intervals:latest
docker run -it stefansavev/python_merge_intervals:latest
[(2, 23), (25, 30)]
docker build -f Dockerfile.go . -t stefansavev/go_merge_intervals
Sending build context to Docker daemon  200.2kB
Step 1/7 : FROM golang:1.15.3-buster
 ---> 4a581cd6feb1
Step 2/7 : RUN mkdir /app && go get github.com/google/go-cmp/cmp
 ---> Using cache
 ---> eec130f41ea5
Step 3/7 : WORKDIR /app
 ---> Using cache
 ---> 3407a1a23f10
Step 4/7 : COPY golang_merge_intervals/ golang_merge_intervals/
 ---> Using cache
 ---> a82c08a08444
Step 5/7 : RUN cd golang_merge_intervals/merge_intervals && go test
 ---> Using cache
 ---> ab4c38df4d21
Step 6/7 : RUN cd golang_merge_intervals/main && go build demo.go
 ---> Using cache
 ---> 2448d2d5ec66
Step 7/7 : CMD golang_merge_intervals/main/demo
 ---> Using cache
 ---> 5e0b95f089ea
Successfully built 5e0b95f089ea
Successfully tagged stefansavev/go_merge_intervals:latest
docker run -it stefansavev/go_merge_intervals:latest
Output: [{"Start":2,"End":23},{"Start":25,"End":30}]
```
