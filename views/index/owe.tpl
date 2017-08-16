<div class="row">
  <div class="col-md-9">
    <div class="panel panel-default">
      <div class="panel-heading">
        <ul class="nav nav-pills">
        <li><a href="/">主页</a></li>
        <li><a href="/repo">我的项目</a></li>
        <li class="active"><a href="/owe">欠我赞的人</a></li>
        <li><a href="/follow">互粉主页</a></li>
        <li><a href="/follower">我的粉丝</a></li>
        <li><a href="/follow_owe">欠我粉的人</a></li>
        {{template "components/star_badge.tpl" .}}
        </ul>
      </div>
      <div class="panel-body paginate-bot">
        {{template "components/flash_error.tpl" .}}
        {{range .Owe}}
        <div class="media">
          <div class="media-left">
          </div>
          <div class="media-body">
            <div class="title">
              <a target="_blank" href="/users/{{.Target}}">{{.Target}}还欠我 <span class="label label-default">{{.Score}}</span> 个赞</a>
            </div>
            <p class="gray">
              <a target="_blank" href="/users/{{.Target}}"><span class="label label-primary">{{.Target}}</span></a>
              <span>•</span>
              <span class="label label-warning">QQ：{{.QQ}}</span>
              <span>•</span>
              <span class="label label-danger">昵称：{{.Nickname}}</span>
            </p>
            <p class="gray">
               <span>我已赞他{{len .StarringRepos}}个项目：{{range .StarringRepos}} <a rel="noreferrer" target="_blank" href="https://github.com/{{.}}"><span class="label label-success">{{.}}</span></a> {{end}}</span>
            </p>
            <p class="gray">
              <span>他已赞我{{len .StarredRepos}}个项目：{{range .StarredRepos}} <a rel="noreferrer" target="_blank" href="https://github.com/{{.}}"><span class="label label-success">{{.}}</span></a> {{end}}</span>
            </p>
            <p class="gray">
              <span>我还可以赞他{{len .CanStarRepos}}个项目：{{range .CanStarRepos}} <a rel="noreferrer" target="_blank" href="https://github.com/{{.}}"><span class="label label-default">{{.}}</span></a> {{end}}</span>
            </p>
            <p class="gray">
              <span>他还可以赞我{{len .CanBeStarredRepos}}个项目：{{range .CanBeStarredRepos}} <a rel="noreferrer" target="_blank" href="https://github.com/{{.}}"><span class="label label-default">{{.}}</span></a> {{end}}</span>
            </p>
          </div>
        </div>
        <div class="divide mar-top-5"></div>
        {{end}}
        <ul id="page"></ul>
      </div>
    </div>
  </div>
  <div class="col-md-3">
    {{if .IsLogin}}
      {{template "components/star_usage.tpl" .}}
      {{template "components/star_button.tpl" .}}
    {{else}}
      {{template "components/welcome.tpl" .}}
    {{end}}
    {{template "components/otherbbs.tpl" .}}
  </div>
</div>
<script type="text/javascript" src="/static/js/bootstrap-paginator.min.js"></script>
<script type="text/javascript">
  $(function () {
    $("#tab_{{.S}}").addClass("active");
    $("#page").bootstrapPaginator({
      currentPage: '{{.Page.PageNo}}',
      totalPages: '{{.Page.TotalPage}}',
      bootstrapMajorVersion: 3,
      size: "small",
      onPageClicked: function(e,originalEvent,type,page){
        var s = {{.S}};
        if (s > 0) {
          window.location.href = "/?p=" + page + "&s={{.S}}"
        } else {
          window.location.href = "/?p=" + page
        }
      }
    });
  });
</script>