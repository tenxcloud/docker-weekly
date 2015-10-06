FROM index.tenxcloud.com/edcapding/beegoimage

ADD src/. $GOPATH/src

EXPOSE 8080

WORKDIR $GOPATH/src/docker-weekly

RUN go build -o docker-weekly main.go

RUN chmod +x ./docker-weekly

ENV DB_HOST mysql:3306
ENV DB_USER admin
ENV DB_PASS admin

CMD ["./docker-weekly"]
