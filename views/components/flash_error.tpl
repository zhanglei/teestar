{{if .flash.error}}
<div class="alert alert-danger alert-dismissible" role="alert">
  <button type="button" class="close" data-dismiss="alert" aria-label="Close"><span aria-hidden="true">&times;</span></button>
  {{.flash.error}}
  {{.flash_html_error}}
</div>
{{end}}

{{if .flash.warning}}
<div class="alert alert-warning alert-dismissible" role="alert">
  <button type="button" class="close" data-dismiss="alert" aria-label="Close"><span aria-hidden="true">&times;</span></button>
  {{.flash.warning}}
  {{.flash_html_warning}}
</div>
{{end}}

{{if .flash.notice}}
<div class="alert alert-info alert-dismissible" role="alert">
  <button type="button" class="close" data-dismiss="alert" aria-label="Close"><span aria-hidden="true">&times;</span></button>
  {{.flash.notice}}
  {{.flash_html_notice}}
</div>
{{end}}

{{if .flash.success}}
<div class="alert alert-success alert-dismissible" role="alert">
  <button type="button" class="close" data-dismiss="alert" aria-label="Close"><span aria-hidden="true">&times;</span></button>
  {{.flash.success}}
  {{.flash_html_success}}
</div>
{{end}}