package opyum

import (
	"bytes"
	"strings"
	"text/template"

	"github.com/heromicro/go-playbook/pkg/enumtipe"
	"github.com/heromicro/go-playbook/pkg/helper"
)

var Tmplt_ansible_builtin_yum = `
{{ Indent " " 4}}- name: yum operation
{{ Indent " " 4}}  ansible.builtin.yum:
{{ Indent " " 4}}    name: {{- range $index, $word := .Name }}
{{ Indent " " 4}}      - {{ $word }} 
{{- Indent " " 0}}     {{- end }}
{{- Indent " " 0}}   {{- if .State }}
{{ Indent " " 4}}    state: {{ .State }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .Enablerepo }}
{{ Indent " " 4}}    enablerepo: {{- range $index, $word := .Enablerepo }}
{{ Indent " " 4}}      - {{ $word }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .Disablerepo }}
{{ Indent " " 4}}    disablerepo: {{- range $index, $word := .Disablerepo }}
{{ Indent " " 4}}      - {{ $word }} 
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .Exclude }}
{{ Indent " " 4}}    exclude: {{- range $index, $word := .Exclude }}
{{ Indent " " 4}}      - {{ $word }} 
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .AllowDowngrade }}
{{ Indent " " 4}}    allow_downgrade: {{ .AllowDowngrade }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .Autoremove }}
{{ Indent " " 4}}    autoremove: {{ .Autoremove }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .Bugfix }}
{{ Indent " " 4}}    bugfix: {{ .Bugfix }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .Cacheonly }}
{{ Indent " " 4}}    cacheonly: {{ .Cacheonly }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .ConfFile }}
{{ Indent " " 4}}    conf_file: {{ .ConfFile }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .DisableExcludes }}
{{ Indent " " 4}}    disable_excludes: {{ .DisableExcludes }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .DisableGpgCheck }}
{{ Indent " " 4}}    disable_gpg_check: {{ .DisableGpgCheck }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .DisablePlugin }}
{{ Indent " " 4}}    disable_plugin: {{- range $index, $word := .DisablePlugin }}
{{ Indent " " 4}}      - {{ $word }} 
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- end }} 
{{- Indent " " 0}}   {{- if .DownloadDir }}
{{ Indent " " 4}}    download_dir: {{ .DownloadDir }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .DownloadOnly }}
{{ Indent " " 4}}    download_only: {{ .DownloadOnly }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .EnablePlugin }}
{{ Indent " " 4}}    enable_plugin: {{- range $index, $word := .EnablePlugin }}
{{ Indent " " 4}}      - {{ $word }} 
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .InstallRepoquery }}
{{ Indent " " 4}}    install_repoquery: {{ .InstallRepoquery }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .InstallWeakDeps }}
{{ Indent " " 4}}    install_weak_deps: {{ .InstallWeakDeps }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .Installroot }}
{{ Indent " " 4}}    installroot: {{ .Installroot }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .Installroot }}
{{ Indent " " 4}}    installroot: {{ .Installroot }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .List }}
{{ Indent " " 4}}    list: {{ .List }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .LockTimeout }}
{{ Indent " " 4}}    lock_timeout: {{ .LockTimeout }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .Security }}
{{ Indent " " 4}}    security: {{ .Security }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .Releasever }}
{{ Indent " " 4}}    releasever: {{ .Releasever }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .SkipBroken }}
{{ Indent " " 4}}    skip_broken: {{ .SkipBroken }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .Sslverify }}
{{ Indent " " 4}}    sslverify: {{ .Sslverify }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .UpdateCache }}
{{ Indent " " 4}}    update_cache: {{ .UpdateCache }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .UpdateOnly }}
{{ Indent " " 4}}    update_only: {{ .UpdateOnly }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .UseBackend }}
{{ Indent " " 4}}    use_backend: {{ .UseBackend }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .ValidateCerts }}
{{ Indent " " 4}}    validate_certs: {{ .ValidateCerts }}
{{- Indent " " 0}}   {{- end }}
`

type AnsibleBuiltinYum struct {
	Name []string `json:"name"`
	// Pkg              []string         `json:"pkg"`

	AllowDowngrade   enumtipe.CostomBool       `json:"allow_downgrade"`
	Autoremove       enumtipe.CostomBool       `json:"autoremove"`
	Bugfix           enumtipe.CostomBool       `json:"bugfix"`
	Cacheonly        enumtipe.CostomBool       `json:"cacheonly"`
	ConfFile         string                    `json:"conf_file"`
	DisableExcludes  string                    `json:"disable_excludes"`
	DisableGpgCheck  enumtipe.CostomBool       `json:"disable_gpg_check"`
	DisablePlugin    []string                  `json:"disable_plugin"`
	EnablePlugin     []string                  `json:"enable_plugin"`
	Disablerepo      []string                  `json:"disablerepo"`
	Enablerepo       []string                  `json:"enablerepo"`
	DownloadDir      string                    `json:"download_dir"`
	DownloadOnly     enumtipe.CostomBool       `json:"download_only"`
	Exclude          []string                  `json:"exclude"`
	InstallRepoquery enumtipe.CostomBool       `json:"install_repoquery"`
	InstallWeakDeps  enumtipe.CostomBool       `json:"install_weak_deps"`
	Installroot      string                    `json:"installroot"`
	List             string                    `json:"list"`
	LockTimeout      int32                     `json:"lock_timeout"`
	Releasever       string                    `json:"releasever"`
	Security         enumtipe.CostomBool       `json:"security"`
	SkipBroken       enumtipe.CostomBool       `json:"skip_broken"`
	Sslverify        enumtipe.CostomBool       `json:"sslverify"`
	State            enumtipe.YumDnfState      `json:"state"`
	UpdateCache      enumtipe.CostomBool       `json:"update_cache"`
	UpdateOnly       enumtipe.CostomBool       `json:"update_only"`
	UseBackend       enumtipe.YumDnfUseBackend `json:"use_backend"`
	ValidateCerts    enumtipe.CostomBool       `json:"validate_certs"`
}

func (a *AnsibleBuiltinYum) String() string {
	return jsonhelper.MarshalToString(a)
}

func (a *AnsibleBuiltinYum) MakeAnsibleTask() (string, error) {

	// var tmpl *template.Template
	funcMap := template.FuncMap{
		"Indent": strings.Repeat,
	}

	tmpl := template.Must(template.New("ansible_builtin_yum").Funcs(funcMap).Parse(Tmplt_ansible_builtin_yum))
	var buff bytes.Buffer

	err := tmpl.Execute(&buff, *a)
	if err != nil {
		return "", err
	}

	return buff.String(), nil
}
