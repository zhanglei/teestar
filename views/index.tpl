<div class="row">
  <div class="col-md-9">
    <div class="panel panel-default">
      <div class="panel-heading">
        <ul class="nav nav-pills">
          <li id="tab_0"><a href="/?tab=all">全部</a></li>
          {{range .Sections}}
          <li id="tab_{{.Id}}"><a href="/?s={{.Id}}">{{.Name}}</a></li>
          {{end}}
        </ul>
      </div>
      <div class="panel-body paginate-bot">
        {{range .Recommend}}
        <div class="media">
          <div class="media-left">
          </div>
          <div class="media-body">
            <div class="title">
              <a target="_blank" href="https://github.com/{{.Repo}}">{{.Repo}}</a>
            </div>
            <p class="gray">
              <a target="_blank" href="https://github.com/{{.Target}}"><span class="label label-primary">{{.Target}}</span></a>
              <span>•</span>
              <span class="hidden-sm hidden-xs">{{.Target}}还欠我{{.Score}}个赞</span>
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