<div class="row">
  <div class="col-md-9">
    {{if .TargetInfo}}
    <div class="panel panel-default">
      <div class="panel-body">
        <div class="media">
          <div class="media-body">
            <h3 style="margin-top: 0">{{.TargetInfo.User}}</h3>
            <p><i class="gray"><a href="https://github.com/{{.TargetInfo.User}}" rel="noreferrer" target="_blank">https://github.com/{{.TargetInfo.User}}</a></i></p>
            <div>GitHub小号: <a href="https://github.com/{{.TargetInfo.Hitter}}" rel="noreferrer" target="_blank">{{.TargetInfo.Hitter}}</a></div>
            <div>QQ号: {{.TargetInfo.QQ}}</div>
            <div>QQ昵称: {{.TargetInfo.Nickname}}</div>
            <div>入驻时间: {{.TargetInfo.CreatedAt}}</div>
            {{if .UserInfo.IsAdmin}}<div>是否为管理员: {{.TargetInfo.IsAdmin}}</div>{{end}}
            <div>账户状态: {{if .TargetInfo.IsDisabled}}<font color="#FF0000">已被管理员禁用</font>{{else}}正常{{end}}</div>
            <div>项目个数: {{.TargetInfo.RepoCount}}</div>
            <b><div>点赞次数: {{.TargetInfo.StarringCount}}</div></b>
            <b><div>被点赞次数: {{.TargetInfo.StarredCount}}</div></b>
            <b><div>欠赞次数: {{.TargetInfo.OweCount}}</div></b>
          </div>
        </div>
      </div>
    </div>
    <div class="panel panel-default">
      <div class="panel-heading">
        他的项目
      </div>
      <div class="table-responsive">
        <table class="table table-striped table-responsive">
          <tbody>
          {{range .TargetRepos}}
          <tr>
            <td><a rel="noreferrer" target="_blank" href="https://github.com/{{.Name}}">{{.Name}}：获得 <span class="label label-default">{{len .Stargazers}}</span> 个赞</a></td>
          </tr>
          {{end}}
          </tbody>
        </table>
      </div>
    </div>
    {{else}}
    <div class="panel panel-default">
      <div class="panel-body">用户不存在</div>
    </div>
    {{end}}
  </div>
  <div class="col-md-3 hidden-sm hidden-xs">

  </div>
</div>