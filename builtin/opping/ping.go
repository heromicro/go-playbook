package opping

import (
	"bytes"
	"text/template"

	"github.com/heromicro/go-playbook/pkg/helper"
)

var Tmplt_ansible_ping = `
{{ Indent " " 4}}- name: {{ .Name }}  
{{ Indent " " 4}}    ansible.builtin.ping:
`

type AnsibleBuiltinPing struct {
	Name string `json:"name"`
}

func (a *AnsibleBuiltinPing) String() string {
	return helper.MarshalToString(a)
}

func (a *AnsibleBuiltinPing) MakeAnsibleTask() (string, error) {

	funcMap := template.FuncMap{
		"Indent": strings.Repeat,
	}

	tmpl := template.Must(template.New("ansible_builtin_ping").Funcs(funcMap).Parse(Tmplt_ansible_ping))
	var buff bytes.Buffer

	err := tmpl.Execute(&buff, *a)
	if err != nil {
		return "", err
	}

	return buff.String(), nil
}
