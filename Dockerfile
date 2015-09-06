FROM golang:1.4.2

ADD src/. $GOPATH/src

EXPOSE 8080

WORKDIR $GOPATH/src/docker-weekly

RUN go build -o docker-weekly main.go

RUN chmod +x ./docker-weekly

CMD ["./docker-weekly"]
