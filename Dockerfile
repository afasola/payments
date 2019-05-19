FROM golang:latest
RUN go get -u google.golang.org/grpc
ADD . /go/src/github.com/afasola/payments
RUN go install github.com/afasola/payments/server
ENTRYPOINT ["/go/bin/server"]
EXPOSE 20000
