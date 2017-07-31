<div class="row">
  <div class="col-md-9">
    <div class="panel panel-default">
      <div class="panel-heading">
        用户列表
        <span class="pull-right">{{.TotalCount}}个用户</span>
      </div>
      {{if .UserInfo.IsAdmin}}
      <div class="table-responsive">
        <table class="table table-striped">
          <tbody>
          <tr>
            <td>用户名</td>
            <td>点赞账号</td>
            <td>QQ号</td>
            <td>QQ昵称</td>
            <td>入驻时间</td>
            <td>是否为管理员</td>
            <td>是否已禁用</td>
            <td>Star缓存</td>
          </tr>
          {{range .UserInfos}}
          <tr>
            <td><a href="/users/{{.User}}" target="_blank">{{.User}}</a></td>
            <td><a href="https://github.com/{{.Hitter}}" target="_blank">{{.Hitter}}</a></td>
            <td>{{.QQ}}</td>
            <td>{{.Nickname}}</td>
            <td>{{.CreatedAt}}</td>
            <td>{{.IsAdmin}}</td>
            <td>{{.IsDisabled}}</td>
            <td><a target="_blank" href="api/users/{{.User}}/starring-repos/update">更新</a></td>
          </tr>
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
  <div class="col-md-3 hidden-sm hidden-xs">

  </div>
</div>