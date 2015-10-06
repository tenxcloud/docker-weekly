# 一周 Docker 镜像精选
# 使用时速云 [代码构建] 和 [持续集成]
# Version:1.0.0

FROM index.tenxcloud.com/edcapding/beegoimage
MAINTAINER Qi Gong Huang <huangqg@tenxcloud.com>

ADD src/. $GOPATH/src

EXPOSE 8080

WORKDIR $GOPATH/src/docker-weekly

RUN go build -o docker-weekly main.go

RUN chmod +x ./docker-weekly

ENV DB_HOST mysql:3306
ENV DB_USER admin
ENV DB_PASS admin

CMD ["./docker-weekly"]
