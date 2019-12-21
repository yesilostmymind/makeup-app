FROM golang:1.8

WORKDIR /go/src/dbwebpage
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD["dbwebpage"]