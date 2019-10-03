FROM golang:1.9


ENV GOPATH=/go
ENV COMPOSE_HTTP_TIMEOUT=3600


COPY . /go/src/
WORKDIR /go/src/


RUN go get github.com/lib/pq && go get github.com/golang/protobuf/proto && go get github.com/jinzhu/gorm && go get github.com/satori/go.uuid