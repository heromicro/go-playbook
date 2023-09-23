package oppackage

import (
	"bytes"
	"strings"
	"text/template"

	"github.com/heromicro/go-playbook/pkg/enumtipe"
	"github.com/heromicro/go-playbook/pkg/helper"
)

// https://docs.ansible.com/ansible/latest/collections/ansible/builtin/package_module.html
var Tmplt_ansible_builtin_package = `
{{ Indent " " 4}}- name: package
{{ Indent " " 4}}  ansible.builtin.package:
{{ Indent " " 4}}    name: {{- range $index, $word := .Name }}
{{ Indent " " 4}}      - {{ $word }} 
{{- Indent " " 0}}     {{- end }}
{{ Indent " " 4}}    state: {{ .State }}
{{- Indent " " 0}}   {{- if .Use }}
{{ Indent " " 4}}    use: {{ .Use }}
{{- Indent " " 0}}   {{- end }}
`

type AnsibleBuiltinPackage struct {
	Name  []string             `json:"name"`
	State enumtipe.CommonState `json:"state"`
	Use   string               `json:"use"`
}

func (a *AnsibleBuiltinPackage) String() string {
	return helper.MarshalToString(a)
}

func (a *AnsibleBuiltinPackage) MakeAnsibleTask() (string, error) {

	// var tmpl *template.Template
	funcMap := template.FuncMap{
		"Indent": strings.Repeat,
	}

	tmpl := template.Must(template.New("ansible_builtin_package").Funcs(funcMap).Parse(Tmplt_ansible_builtin_package))
	var buff bytes.Buffer

	err := tmpl.Execute(&buff, *a)
	if err != nil {
		return "", err
	}

	return buff.String(), nil
}
