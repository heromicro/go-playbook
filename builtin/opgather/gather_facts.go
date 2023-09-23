package opgather

import (
	"bytes"
	"text/template"

	"github.com/heromicro/go-playbook/pkg/enumtipe"
	"github.com/heromicro/go-playbook/pkg/helper"
)

var Tmplt_ansible_gather_facts = `
{{ Indent " " 4}}- name: {{ .Name }} 
{{ Indent " " 4}}  ansible.builtin.gather_facts:
{{- Indent " " 0}}   {{- if .Parallel }}
{{ Indent " " 4}}    parallel: {{ .Parallel }}
{{- Indent " " 0}}   {{- else }}
{{ Indent " " 4}}    parallel: false
{{- Indent " " 0}}   {{- end }}
`

type AnsibleBuiltinGatherFacts struct {
	Name     string              `json:"name"`
	Parallel enumtipe.CostomBool `json:"parallel"`
}

func (a *AnsibleBuiltinGatherFacts) String() string {
	return helper.MarshalToString(a)
}

func (a *AnsibleBuiltinGatherFacts) MakeAnsibleTask() (string, error) {

	funcMap := template.FuncMap{
		"Indent": strings.Repeat,
	}

	tmpl := template.Must(template.New("ansible_builtin_gather_facts").Funcs(funcMap).Parse(Tmplt_ansible_gather_facts))
	var buff bytes.Buffer

	err := tmpl.Execute(&buff, *a)
	if err != nil {
		return "", err
	}

	return buff.String(), nil
}
