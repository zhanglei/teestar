<div class="panel panel-default">
  <div class="panel-heading">
    使用说明
  </div>
  <div class="panel-body">
    <div class="media">
      <div class="media-left">
        <a href="/user/{{.UserInfo.Username}}">
          <img src="{{.UserInfo.Avatar}}" title="{{.UserInfo.Username}}" class="avatar">
        </a>
      </div>
      1. 如果GitHub点赞的账号与用户名不同，则填写后，保存设置；<br>
      2. 点击“我需要被别人点赞的项目”右侧的“添加项目”，把自己需要点赞的项目添加上去。
      <div class="media-body">
        <div class="media-heading">
          <a href="/user/{{.UserInfo.Username}}">{{.UserInfo.Username}}</a>
        </div>
        {{if .UserInfo.Url}}<a href="{{.UserInfo.Url}}" target="_blank">{{.UserInfo.Url}}</a>{{end}}
      </div>
      {{if .UserInfo.Signature}}
      <div style="color: #7A7A7A; font-size: 12px; margin-top:5px;">
        <i>“ {{.UserInfo.Signature}} ” </i>
      </div>
      {{end}}
    </div>
  </div>
</div>