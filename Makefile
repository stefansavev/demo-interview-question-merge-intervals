all: python_run golang_run

python_build:
	docker build -f Dockerfile.python . -t stefansavev/python_merge_intervals

python_run: python_build
	docker run -it stefansavev/python_merge_intervals:latest

golang_build:
	docker build -f Dockerfile.go . -t stefansavev/go_merge_intervals

golang_run: golang_build
	docker run -it stefansavev/go_merge_intervals:latest
