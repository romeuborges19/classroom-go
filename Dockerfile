FROM golang:latest

WORKDIR /go-classroom

ADD . /go-classroom

RUN go build -o /server

EXPOSE 8080

CMD ["/server"]
