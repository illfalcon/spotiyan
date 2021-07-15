FROM golang:1.16

WORKDIR /go/src/github.com/illfalcon/spotiyan
COPY . .
RUN go build -o /go/bin/spotiyan cmd/main.go

ENTRYPOINT ["/go/bin/spotiyan"]