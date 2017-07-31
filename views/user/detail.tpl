<div class="row">
  <div class="col-md-9">
    {{if .TargetInfo}}
    <div class="panel panel-default">
      <div class="panel-body">
        <div class="media">
          <div class="media-body">
            <h3 style="margin-top: 0">{{.TargetInfo.User}}</h3>
            <p><i class="gray"><a href="https://github.com/{{.TargetInfo.User}}" target="_blank">https://github.com/{{.TargetInfo.User}}</a></i></p>
            <div>GitHub小号: <a href="https://github.com/{{.TargetInfo.Hitter}}" target="_blank">{{.TargetInfo.Hitter}}</a></div>
            <div>QQ号: {{.TargetInfo.QQ}}</div>
            <div>QQ昵称: {{.TargetInfo.Nickname}}</div>
            <div>入驻时间: {{.TargetInfo.CreatedAt}}</div>
            <div>是否为管理员: {{.TargetInfo.IsAdmin}}</div>
            <div>账户状态: {{if .TargetInfo.IsDisabled}}<font color="#FF0000">已被管理员禁用</font>{{else}}正常{{end}}</div>
            <div>项目个数: {{.TargetRepoCount}}</div>
            <b><div>点赞次数: {{.TargetStarringCount}}</div></b>
            <b><div>被点赞次数: {{.TargetStarredCount}}</div></b>
            <b><div>欠赞次数: {{.TargetOweCount}}</div></b>
          </div>
        </div>
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