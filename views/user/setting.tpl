<div class="row">
  <div class="col-md-9">
    <div class="panel panel-default">
      <div class="panel-heading">
        <a href="/">主页</a> / 个人设置
      </div>
      <div class="panel-body">
        {{if .flash.success}}
        <div class="alert alert-success alert-dismissible" role="alert">
          <button type="button" class="close" data-dismiss="alert" aria-label="Close"><span
            aria-hidden="true">&times;</span></button>
          {{.flash.success}}
        </div>
        {{end}}
        {{template "../components/flash_error.tpl" .}}
        <form action="/user/setting" method="post">
          <div class="form-group">
            <label for="username">用户名</label>
            <input type="text" disabled="" class="form-control" id="username" value="{{.Username}}">
          </div>
          <div class="form-group">
            <label for="email">GitHub点赞账号</label>
            <input type="text" class="form-control" name="hitter" id="hitter" value="{{.Hitter}}">
          </div>
          <button type="submit" class="btn btn-default">保存设置</button>
        </form>
      </div>
    </div>

    <div class="panel panel-default">
      <div class="panel-heading">
        项目管理（靠前的项目会被优先展示）
        <a href="/repo/add" class="pull-right">添加项目</a>
      </div>
      <div class="table-responsive">
        <table class="table table-striped table-responsive">
          <tbody>
          {{range $i, $v := .Repos}}
          <tr>
            <td>{{$v}}</td>
            <td>
              <a href="javascript:if(confirm('确认删除吗?')) location.href='/repo/delete/{{$v}}'" class="btn btn-xs btn-danger">删除</a>
            </td>
          </tr>
          {{end}}
          </tbody>
        </table>
      </div>
    </div>
  </div>
  <div class="col-md-3 hidden-sm hidden-xs">
    {{template "../components/user_info.tpl" .}}
  </div>
</div>