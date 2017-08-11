<div class="row">
  <div class="col-md-9">
    <div class="panel panel-default">
      <div class="panel-heading">
        <a href="/">主页</a> / 个人设置
      </div>
      <div class="panel-body">
        {{template "../components/flash_error.tpl" .}}
        <form action="/user/setting" method="post">
          <div class="form-group">
            <label for="username">用户名{{if .UserInfo.IsDisabled}}<font color="#FF0000">（该账户已被管理员禁用，请联系管理员，QQ群：646373152）</font>{{end}}</label>
            <input type="text" disabled="" class="form-control" id="username" value="{{.UserInfo.User}}">
          </div>
          <div class="form-group">
            <label for="email">GitHub点赞账号（即给别人点赞用的小号，如果与上面用户名（大号）相同，则留空即可）</label>
            <input type="text" {{if .UserInfo.IsDisabled}}disabled=""{{end}} class="form-control" name="hitter" id="hitter" value="{{.UserInfo.Hitter}}">
            <label for="email">QQ号（必填，当有新点赞需要处理时，管理员适时会在QQ群里进行@通知）</label>
            <input type="text" {{if .UserInfo.IsDisabled}}disabled=""{{end}} class="form-control" name="qq" id="qq" value="{{.UserInfo.QQ}}">
            <label>QQ昵称（推荐填写）</label>
            <input type="text" {{if .UserInfo.IsDisabled}}disabled=""{{end}} class="form-control" name="nickname" id="nickname" value="{{.UserInfo.Nickname}}">
            <label>Email邮箱（可选填，将来每日Star进展会推送至该邮箱）</label>
            <input type="text" {{if .UserInfo.IsDisabled}}disabled=""{{end}} class="form-control" name="email" id="email" value="{{.UserInfo.Email}}">
            <label>是否参与账户互粉(Follow)（必须设置QQ号后才能参与，不参与则不会被圈粉主页展示）</label>
            <div class="switch switch-large"><input type="checkbox" name="followable" id="followable" {{if .UserInfo.IsFollowable}} checked {{else}} unchecked {{end}}></div>
          </div>
          <button type="submit" class="btn btn-default">保存设置</button> <a href="/" class="pull-right">已经设置完毕，带我去点赞</a>
        </form>
      </div>
    </div>

    <div class="panel panel-default">
      <div class="panel-heading">
        我需要被别人点赞的项目（靠前的项目会被优先展示）
        <a href="/repo/add" class="pull-right">添加项目</a>
      </div>
      <div class="table-responsive">
        <table class="table table-striped table-responsive">
          <tbody>
          {{range .EscapedRepos}}
          <tr>
            <td><a rel="noreferrer" target="_blank" href="https://github.com/{{.Repo}}">{{.Repo}}</a></td>
            <td>
              <a id="delete_repo" href="javascript:if(confirm('确认删除项目{{.Repo}}吗?')) location.href='/repo/delete/{{.RepoEscaped}}'" class="btn btn-xs btn-danger">删除</a>
            </td>
          </tr>
          {{end}}
          </tbody>
        </table>
      </div>
    </div>
    <div class="panel panel-default">
      <div class="panel-heading">
        修改密码
      </div>
      <div class="panel-body">
        <form action="/user/changepwd" method="post">
          <div class="form-group">
            <label for="oldpassword">旧密码</label>
            <input type="password" class="form-control" name="oldpassword" id="oldpassword" value="">
          </div>
          <div class="form-group">
            <label for="newpassword">新密码</label>
            <input type="password" class="form-control" name="newpassword" id="newpassword" value="">
          </div>
          <button type="submit" class="btn btn-default">修改密码</button>
        </form>
      </div>
    </div>
  </div>
  <div class="col-md-3 hidden-sm hidden-xs">
    {{template "../components/setting_help_info.tpl" .}}
  </div>
</div>