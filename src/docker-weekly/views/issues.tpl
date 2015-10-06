<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>一周Docker镜像精选-时速云</title>
  <meta charset="UTF-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="renderer" content="webkit">
  <meta name="viewport" content="width=device-width,initial-scale=1">
  <meta property="wb:webmaster" content="dd9a209e374eab5c">
  <meta name="keywords" content="Docker镜像,镜像市场,Docker Hub,时速云,CaaS,Mesos,Swarm,Kubernetes">
  <meta name="description" content="时速云技术期刊：一周Docker镜像精选">
  <link rel="icon" type="image/x-icon" href="/static/img/favicon.ico">
  <script src="https://dn-tenxcloud-static.qbox.me/jquery-2.1.0.min.js"></script>
  <link rel="stylesheet" href="https://dn-tenxcloud-static.qbox.me/bootstrap.min.css">
  <link rel="stylesheet" type="text/css" href="https://dn-tenxcloud-static.qbox.me/index-top.css?v=0.1.0">
  <link rel="stylesheet" href="/static/css/index.css">
</head>
<body>
  <header>
    <div class="container">
      <a href="/"><span class="white-logo"></span></a>
      <div class="pull-right">
        <a href='/'><span class="btn-hotlists todayhotlists">本期推荐</span></a>
        <a href='/issues'><span class="btn-hotlists historyhotlists active">往期回顾</span></a>
      </div>
    </div>
  </header>
  <div id="home">
    <div class="h3 text-center">一周 Docker 镜像精选</div>
    <div class="h5 text-center">往期回顾</div>
    <br>
    <br>
    <div class="container">
      {{range $k, $v := .articles}}
        <div class="col-xs-12 col-md-6 week">
          <a href="/issue/{{$v.Id }}">
            <div class="mirror-body">
              <div class="mirror-date pull-left">
                <div class="h3">第{{$v.Number}}期</div>
                <hr class="hrs">
                <div class="h5">{{$v.Created}}</div>
              </div>
              <div class="mirror-article pull-left">
                <div class="h5">{{$v.Title}}</div>
              </div>
            </div>
          </a>
        </div>
      {{end}}
    </div>
  </div>
  <footer>
    <br>
    <div class="h5 text-center">Copyright © 2015 时速云    京ICP备14045471号</div>
  </footer>
</body>
</html>