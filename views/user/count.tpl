<div class="row">
  <div class="col-md-9">
    <div class="panel panel-default">
      <div class="panel-heading">
        用户统计数据
        <span class="pull-right">{{len .UserInfos}}个用户</span>
      </div>
      {{if .UserInfo.IsAdmin}}
      <div class="table-responsive">
        <table class="table table-striped">
          <tbody>
          <tr>
            <td>用户名</td>
            <td>点赞账号</td>
            <td>昵称</td>
            <td>项目数</td>
            <td>点赞数</td>
            <td>被赞数</td>
            <td>欠赞数</td>
            <td>加粉数</td>
            <td>被粉数</td>
            <td>欠粉数</td>
            <td>Star缓存</td>
            <td>Follow缓存</td>
          </tr>
          {{range .UserInfos}}
          <tr>
            <td><a href="/users/{{.User}}" target="_blank">
            {{if .IsFlagged}}
              <font style="background:#D9534F" color="white">{{.User}}</font>
            {{else if .IsDisabled}}
              <font style="background:#3E3E3E" color="white">{{.User}}</font>
            {{else}}
              {{.User}}
            {{end}}
            </a></td>
            <td><a href="https://github.com/{{.Hitter}}" rel="noreferrer" target="_blank">{{.Hitter}}</a></td>
            <td>{{.Nickname}}</td>
            <td>{{.RepoCount}}</td>
            <td>{{.StarringCount}}</td>
            <td>{{.StarredCount}}</td>
            <td>{{.OweCount}}</td>
            <td>{{.FollowingCount}}</td>
            <td>{{.FollowedCount}}</td>
            <td>{{.FollowOweCount}}</td>
            <td><a target="_blank" href="api/users/{{.User}}/starring-repos/update">刷</a></td>
            <td><a target="_blank" href="api/users/{{.User}}/following-users/update">刷</a></td>
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
</div>