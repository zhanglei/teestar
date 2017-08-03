<!doctype html>
<html lang="en">
<head>
  <meta name="referrer" content="never">
  <meta charset="UTF-8">
  <meta name="viewport"
        content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
  <meta http-equiv="X-UA-Compatible" content="ie=edge">
  <title>{{.PageTitle}}</title>
  <link rel="stylesheet" href="/static/bootstrap/css/bootstrap.min.css"/>
  <link rel="stylesheet" href="/static/css/pybbs.css">
  <script src="//cdn.bootcss.com/jquery/2.2.2/jquery.min.js"></script>
  <script src="//cdn.bootcss.com/bootstrap/3.3.6/js/bootstrap.min.js"></script>
  <script type="text/javascript" src="/static/js/prototype.js"></script>
  <script type="text/javascript" src="/static/js/noreferrer.js"></script>
</head>
<body>
<div class="wrapper">
  <nav class="navbar navbar-inverse">
    <div class="container">
      <div class="navbar-header">
        <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#navbar"
                aria-expanded="false" aria-controls="navbar">
          <span class="sr-only">Toggle navigation</span>
          <span class="icon-bar"></span>
          <span class="icon-bar"></span>
          <span class="icon-bar"></span>
        </button>
        <img class="navbar-brand" src="/static/imgs/star.png" alt="Logo"><a class="navbar-brand" style="color:#fff;" href="/">GitStar - GitHub项目点赞</a>
      </div>
      <div id="navbar" class="navbar-collapse collapse header-navbar">
        <ul class="nav navbar-nav navbar-right">
          {{if .IsLogin}}
          {{if .UserInfo.IsAdmin}}
          <li>
            <a href="/users">用户列表</a>
          </li>
          <li>
            <a href="/count">用户统计数据</a>
          </li>
          <li>
            <a href="/log">系统日志</a>
          </li>
          {{end}}
          <li>
            <a href="/swagger/index.html">API文档</a>
          </li>
          <li>
          <a href="/referrer" rel="noreferrer">Referrer测试</a>
          </li>
          <li>
          <a href="/owes">欠赞排行</a>
          </li>
          {{end}}
          <li>
            <a href="/about">关于</a>
          </li>
          {{if .IsLogin}}
          <li>
            <a target="_blank" href="/users/{{.UserInfo.User}}">
              欢迎您，{{.UserInfo.User}}
            </a>
          </li>
          <li>
            <a href="javascript:;" class="dropdown-toggle" data-toggle="dropdown"
               data-hover="dropdown">
              设置
              <span class="caret"></span>
            </a>
            <span class="dropdown-arrow"></span>
            <ul class="dropdown-menu">
              <li><a href="/user/setting">个人资料</a></li>
              <li><a href="/logout">退出</a></li>
            </ul>
          </li>
          {{else}}
          <li><a href="/login">登录</a></li>
          <li><a href="/register">注册</a></li>
          {{end}}
        </ul>
      </div>
    </div>
  </nav>
  <div class="container">
    {{.LayoutContent}}
  </div>
</div>
<div class="container">
  <br>
  <div class="text-center">
    <p>©2017 GitStar.cn. All Rights Reserved.</p>
    {{if .IsLogin}}<a href="/referrer" rel="noreferrer">点击进入Referrer测试</a>{{end}}
  </div>
  <br>
</div>
</body>
</html>