# docker-weekly
时速云技术周刊《一周 Docker 镜像精选》  [http://weekly.tenxcloud.com/]( http://weekly.tenxcloud.com/)

## 如何部署到时速云平台

1）创建 mysql 容器实例

2）创建 docker-weekly 实例，并指定 mysql 配置信息，在高级设置里配置如下环境变量

    DB_HOST <mysql:3306>
    DB_USER <admin>
    DB_PASS <admin>

## 如何使用

### Step 1: fork ```git@github.com:tenxcloud/docker-weekly.git```

### Step 2: start app
```
cd ~/git && git clone git@github.com:Samurais/docker-weekly.git
cd docker-weekly
export GOPATH=~/git/docker-weekly
cd src/docker-weekly
go get github.com/astaxie/beego
go run main.go
```

### Step 3: get your index

```
open http://localhost:8080/
```
