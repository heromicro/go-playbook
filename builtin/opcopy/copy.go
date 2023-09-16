package opcopy

import (
	"bytes"
	"text/template"

	"github.com/heromicro/go-playbook/pkg/enumtipe"
	"github.com/heromicro/go-playbook/pkg/helper"
)

// https://docs.ansible.com/ansible/latest/collections/ansible/builtin/copy_module.html
var Tmplt_ansible_copy_file = `
{{ Indent " " 4}}- name: copy {{- if .Content }} content {{- end }} {{- if .Src }} {{- .Src }} {{- end }} to {{ .Dest }} 
{{ Indent " " 4}}  ansible.builtin.copy:
{{- Indent " " 0}}   {{- if .Content }}
{{ Indent " " 4}}    content: | {{- range $index, $word := .Content }}
{{ Indent " " 4}}      {{ $word }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .Src }} 
{{ Indent " " 4}}    src: {{ .Src }}
{{- Indent " " 0}}   {{- end }}
{{ Indent " " 4}}    dest: {{ .Dest }}
{{- Indent " " 0}}   {{- if .Owner }}
{{ Indent " " 4}}    owner: {{ .Owner }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .Group }}
{{ Indent " " 4}}    group: {{ .Group }} 
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .Mode }}
{{ Indent " " 4}}    mode: {{ .Mode }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .RemoteSrc }}
{{ Indent " " 4}}    remote_src: {{ .RemoteSrc }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .Follow }}
{{ Indent " " 4}}    follow: {{ .Follow }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .Backup }}
{{ Indent " " 4}}    backup: {{ .Backup }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .Validate }}
{{ Indent " " 4}}    validate: {{ .Validate }}
{{- Indent " " 0}}   {{- end }}
`

type AnsibleCopyFile struct {
	Content []string `json:"content"`
	Src     string   `json:"src"`
	Dest    string   `json:"dest"`

	Owner string `json:"owner"`
	Group string `json:"group"`
	Mode  string `json:"mode"` // '0644', u+rw,g-wx,o-rwx

	RemoteSrc enumtipe.CostomBool `json:"remote_src"` // true, false
	Follow    enumtipe.CostomBool `json:"follow"`     // true, false
	Backup    enumtipe.CostomBool `json:"backup"`     // true, false
	Validate  string              `json:"validate"`   // true, false
}

func (a *AnsibleCopyFile) String() string {
	return helper.MarshalToString(a)
}

func (a *AnsibleCopyFile) MakeAnsibleTask() (string, error) {

	funcMap := template.FuncMap{
		"Indent": strings.Repeat,
	}

	tmpl := template.Must(template.New("ansible_builtin_copy").Funcs(funcMap).Parse(Tmplt_ansible_copy))
	var buff bytes.Buffer

	err := tmpl.Execute(&buff, *a)
	if err != nil {
		return "", nil
	}

	return buff.String(), nil
}
