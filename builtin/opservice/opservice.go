package opservice

import (
	"bytes"
	"strings"
	"text/template"

	"github.com/heromicro/go-playbook/pkg/enumtipe"
	"github.com/heromicro/go-playbook/pkg/helper"
)

// https://docs.ansible.com/ansible/latest/collections/ansible/builtin/service_module.html
var Tmplt_ansible_builtin_service = `
{{ Indent " " 4}}- name: service {{ .Name }}  
{{ Indent " " 4}}  ansible.builtin.service:
{{ Indent " " 4}}    name: {{ .Name }}
{{- Indent " " 0}}   {{- if .State }}
{{ Indent " " 4}}    state: {{ .State }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .Enabled }}
{{ Indent " " 4}}    enabled: {{ .Enabled }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .Sleep }}
{{ Indent " " 4}}    sleep: {{ .Sleep }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .Pattern }}
{{ Indent " " 4}}    pattern: {{ .Pattern }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .Arguments }}
{{ Indent " " 4}}    args: {{ .Arguments }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .Runlevel }}
{{ Indent " " 4}}    runlevel: {{ .Runlevel }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .Use }}
{{ Indent " " 4}}    use: {{ .Use }}
{{- Indent " " 0}}   {{- end }}
`

type AnsibleBuiltinService struct {
	Name      string              `json:"name"`
	Arguments string              `json:"arguments"`
	Enabled   enumtipe.CostomBool `json:"enabled"`
	Pattern   string              `json:"pattern"`
	RunLevel  string              `json:"run_level"`
	Sleep     int32               `json:"sleep"`
	State     ServiceState        `json:"state"`
	Use       string              `json:"use"`
}

func (a *AnsibleBuiltinService) String() string {
	return helper.MarshalToString(a)
}

func (a *AnsibleBuiltinService) MakeAnsibleTask() (string, error) {

	// var tmpl *template.Template
	funcMap := template.FuncMap{
		"Indent": strings.Repeat,
	}

	tmpl := template.Must(template.New("ansible_builtin_service").Funcs(funcMap).Parse(Tmplt_ansible_builtin_service))
	var buff bytes.Buffer

	err := tmpl.Execute(&buff, *a)
	if err != nil {
		return "", err
	}

	return buff.String(), nil
}
