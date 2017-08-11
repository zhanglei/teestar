<div class="row">
  <div class="col-md-9">
    <div class="panel panel-default">
      <div class="panel-heading">
        <ul class="nav nav-pills">
        <li class="active"><a href="/">主页</a></li>
        <li><a href="/repo">我的项目</a></li>
        <li><a href="/owe">欠我赞的人</a></li>
        <li><a href="/follow">互粉主页</a></li>
        <li><a href="/follower">我的粉丝</a></li>
        <li><a href="/follow_owe">欠我粉的人</a></li>
        <div class="pull-right">我的点赞: <span class="label label-primary">{{.UserInfo.StarringCount}}</span> &nbsp;&nbsp; 我被点赞: <span class="label label-primary">{{.UserInfo.StarredCount}}</span> &nbsp;&nbsp; 欠赞: <span class="label label-primary">{{.UserInfo.OweCount}}</span></div>
        </ul>
      </div>
      <div class="panel-body paginate-bot">
        {{template "components/flash_error.tpl" .}}
        {{range .Recommend}}
        <div class="media">
          <div class="media-left">
          </div>
          <div class="media-body">
            <div class="title">
              <a rel="noreferrer" target="_blank" href="https://github.com/{{.Repo}}">{{.Repo}}</a>
            </div>
            <p class="gray">
              <a target="_blank" href="/users/{{.Target}}"><span class="label label-primary">{{.Target}}</span></a>
              <span>•</span>
              {{if ge .Score 0}}
              <span class="hidden-sm hidden-xs">{{.Target}}还欠我{{.Score}}个赞</span>
              {{else}}
              <b><span class="hidden-sm hidden-xs">我还欠{{.Target}} {{.ScoreR}}个赞</span></b>
              {{end}}
            </p>
          </div>
        </div>
        <div class="divide mar-top-5"></div>
        {{end}}
        <ul id="page"></ul>
      </div>
    </div>
  </div>
  <div class="col-md-3 hidden-sm hidden-xs">
    {{if .IsLogin}}
      {{template "components/user_info.tpl" .}}
      {{template "components/topic_create.tpl" .}}
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