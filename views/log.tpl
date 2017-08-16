<div class="row">
  <div class="col-md-9">
    <div class="panel panel-default">
      <div class="panel-heading">
      系统日志
      </div>
      {{if .UserInfo.IsAdmin}}
      <div class="table-responsive">
        <table class="table table-striped">
          <tbody>
          {{range .Log}}
          <tr><td>{{.}}</td></tr>
          {{end}}
          </tbody>
        </table>
      </div>
      {{else}}
        <div class="panel panel-default">
          <div class="panel-body">该页仅允许管理员查看</div>
        </div>
      {{end}}
      <div class="panel-body" style="padding: 0 15px;">
        <ul id="page"></ul>
      </div>
    </div>
  </div>
</div>