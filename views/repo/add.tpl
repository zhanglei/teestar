<div class="row">
  <div class="col-md-9">
    <div class="panel panel-default">
      <div class="panel-heading">添加项目</div>
      <div class="panel-body">
        {{template "../components/flash_error.tpl" .}}
        <form action="/repo/add" method="post">
          <div class="form-group">
            <label for="name">项目地址</label>
            <input type="text" id="name" name="name" class="form-control" placeholder="格式为：user_name/repo_name或者https://github.com/user_name/repo_name都可以">
          </div>
          <button type="submit" class="btn btn-sm btn-default">添加</button>
        </form>
      </div>
    </div>
  </div>
</div>