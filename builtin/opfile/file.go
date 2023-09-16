package opfile

import (
	"bytes"
	"text/template"

	"github.com/heromicro/go-playbook/pkg/enumtipe"
	"github.com/heromicro/go-playbook/pkg/helper"
)

// https://docs.ansible.com/ansible/latest/collections/ansible/builtin/file_module.html
var Tmplt_ansible_file = `
{{ Indent " " 4}}- name: {{ .Name }} 
{{ Indent " " 4}}  ansible.builtin.file:
{{ Indent " " 4}}    path: {{.Path}}
{{- Indent " " 0}}   {{- if .State }}
{{ Indent " " 4}}    state: {{ .State }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .Recurse }}
{{ Indent " " 4}}    recurse: {{ .Recurse }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .Owner }}
{{ Indent " " 4}}    owner: {{ .Owner }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .Group }}
{{ Indent " " 4}}    group: {{ .Group }} 
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .Mode }}
{{ Indent " " 4}}    mode: {{ .Mode }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .Src }}
{{ Indent " " 4}}    src: {{ .Src }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .Dest }}
{{ Indent " " 4}}    dest: {{ .Dest }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .AccessTime }}
{{ Indent " " 4}}    access_time: {{ .AccessTime }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .AccessTimeFormat }}
{{ Indent " " 4}}    access_time_format: {{ .AccessTimeFormat }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .ModificationTime }}
{{ Indent " " 4}}    modification_time: {{ .ModificationTime }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .ModificationTimeFormat }}
{{ Indent " " 4}}    modification_time_format: {{ .ModificationTimeFormat }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .Attributes }}
{{ Indent " " 4}}    attributes: {{ .Attributes }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .Follow }}
{{ Indent " " 4}}    follow: {{ .Follow }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .Force }}
{{ Indent " " 4}}    force: {{ .Force }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .UnsafeWrites }}
{{ Indent " " 4}}    unsafe_writes: {{ .UnsafeWrites }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .SeLevel }}
{{ Indent " " 4}}    selevel: {{ .SeLevel }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .SeRole }}
{{ Indent " " 4}}    serole: {{ .SeRole }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .SeType }}
{{ Indent " " 4}}    seuser: {{ .SeType }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .SeUser }}
{{ Indent " " 4}}    seuser: {{ .SeUser }}
{{- Indent " " 0}}   {{- end }}
`

type AnsibleBuiltinFile struct {
	Name    string              `json:"name"`
	Path    string              `json:"path"`
	Recurse enumtipe.CostomBool `json:"recurse"`
	State   FileState           `json:"state"`
	Owner   string              `json:"owner"`
	Group   string              `json:"group"`
	Mode    string              `json:"mode"` // '0644', u+rw,g-wx,o-rwx
	Src     string              `json:"src"`
	Dest    string              `json:"dest"`

	AccessTime             string `json:"access_time"`
	AccessTimeFormat       string `json:"access_time_format"`
	ModificationTime       string `json:"modification_time"`
	ModificationTimeFormat string `json:"modification_time_format"`

	Attributes   string              `json:"attributes"`
	Follow       enumtipe.CostomBool `json:"follow"`
	Force        enumtipe.CostomBool `json:"force"`
	UnsafeWrites enumtipe.CostomBool `json:"unsafe_writes"`

	SeLevel string `json:"selevel"`
	SeRole  string `json:"serole"`
	SeType  string `json:"setype"`
	SeUser  string `json:"seuser"`
}

func (a *AnsibleBuiltinFile) String() string {
	return helper.MarshalToString(a)
}

func (a *AnsibleBuiltinFile) MakeAnsibleTask() (string, error) {

	funcMap := template.FuncMap{
		"Indent": strings.Repeat,
	}

	tmpl := template.Must(template.New("ansible_builtin_file").Funcs(funcMap).Parse(Tmplt_ansible_file))
	var buff bytes.Buffer

	err := tmpl.Execute(&buff, *a)
	if err != nil {
		return "", err
	}

	return buff.String(), nil
}
