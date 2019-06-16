FROM  golang:1.12-alpine

ENV GO111MODULE=on

ADD ./ /go/src/api

WORKDIR /go/src/api

CMD ["go", "run", "main.go"]

