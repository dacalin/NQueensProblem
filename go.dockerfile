FROM golang:1.14

RUN go get -u github.com/wk8/go-ordered-map

WORKDIR /go/src/app

CMD ["app"]

