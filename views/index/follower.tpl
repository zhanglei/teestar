<div class="row">
  <div class="col-md-9">
    <div class="panel panel-default">
      <div class="panel-heading">
        <ul class="nav nav-pills">
        <li><a href="/">主页</a></li>
        <li><a href="/repo">我的项目</a></li>
        <li><a href="/owe">欠我赞的人</a></li>
        <li><a href="/follow">互粉主页</a></li>
        <li class="active"><a href="/follower">我的粉丝</a></li>
        <li><a href="/follow_owe">欠我粉的人</a></li>
        {{template "components/follow_badge.tpl" .}}
        </ul>
      </div>
      <div class="panel-body paginate-bot">
        {{template "components/flash_error.tpl" .}}
        {{range .Followers}}
        <div class="media">
          <div class="media-left">
          </div>
          <div class="media-body">
            <div class="title">
              <a rel="noreferrer" target="_blank" href="https://github.com/{{.User}}">{{.User}}</a>
            </div>
            <p class="gray">
              <a target="_blank" href="/users/{{.User}}"><span class="label label-primary">{{.User}}</span></a>
              <span>•</span>
              {{if .Followed}}
              <span class="label label-success">我已经粉他</span>
              {{else}}
              <span class="label label-danger">我还没有粉他</span>
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
      {{template "components/follow_usage.tpl" .}}
      {{template "components/follow_button.tpl" .}}
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