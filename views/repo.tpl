<div class="row">
  <div class="col-md-9">
    <div class="panel panel-default">
      <div class="panel-heading">
        <ul class="nav nav-pills">
        <li><a href="/">主页</a></li>
        <li class="active"><a href="/repo">我的项目</a></li>
        <li><a href="/owe">欠我赞的人</a></li>
        <div class="pull-right">我的点赞: <span class="label label-primary">{{.UserStarringCount}}</span> &nbsp;&nbsp; 我被点赞: <span class="label label-primary">{{.UserStarredCount}}</span> &nbsp;&nbsp; 欠赞: <span class="label label-primary">{{.UserOweCount}}</span></div>
        </ul>
      </div>
      <div class="panel-body paginate-bot">
        {{template "components/flash_error.tpl" .}}
        {{range .Repos}}
        <div class="media">
          <div class="media-left">
          </div>
          <div class="media-body">
            <div class="title">
              <a rel="noreferrer" target="_blank" href="https://github.com/{{.Name}}">{{.Name}}：获得 <span class="label label-default">{{len .Stargazers}}</span> 个赞</a>
            </div>
            <p class="gray">
              {{if ne (len .Stargazers) 0}}
                被 {{range .Stargazers}}<a target="_blank" href="/users/{{.}}"><span class="label label-primary">{{.}}</span></a> {{end}}
                等{{len .Stargazers}}人点赞
              {{else}}
                0人点赞
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