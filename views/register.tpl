<div class="row">
  <div class="col-md-6">
    <div class="panel panel-default">
      <div class="panel-heading">注册</div>
      <div class="panel-body">
        {{template "components/flash_error.tpl" .}}
        <form action="/register" method="post">
          <div class="form-group">
            <label for="username">用户名</label>
            <input type="text" id="username" name="username" class="form-control" placeholder="与GitHub用户名一致，不是邮箱">
          </div>
          <div class="form-group">
            <label for="password">密码</label>
            <input type="password" id="password" name="password" class="form-control" placeholder="可以为任意密码，不是GitHub密码">
          </div>
          <div class="form-group">
            <label for="qq">QQ号</label>
            <input type="text" id="qq" name="qq" class="form-control" placeholder="你的QQ号码">
          </div>
          <input type="submit" class="btn btn-default" value="注册"> <a href="/login">去登录</a>
        </form>
      </div>
    </div>
  </div>
</div>