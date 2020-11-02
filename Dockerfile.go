FROM golang:1.15.3-buster

RUN mkdir /app && go get github.com/google/go-cmp/cmp
WORKDIR /app

COPY golang_merge_intervals/ golang_merge_intervals/

RUN cd golang_merge_intervals/merge_intervals && go test
RUN cd golang_merge_intervals/main && go build demo.go

CMD golang_merge_intervals/main/demo



