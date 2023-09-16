package opscript

import (
	"bytes"
	"text/template"

	"github.com/heromicro/go-playbook/pkg/args"
	"github.com/heromicro/go-playbook/pkg/enumtipe"
	"github.com/heromicro/go-playbook/pkg/helper"
)

// https://docs.ansible.com/ansible/latest/collections/ansible/builtin/script_module.html
var Tmplt_ansible_builtin_script = `
{{ Indent " " 4}}- name: script 
{{ Indent " " 4}}  ansible.builtin.script: 
{{- Indent " " 0}}   {{- if .Cmd }} 
{{ Indent " " 4}}    cmd: {{ .Cmd }} 
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .ChDir }}
{{ Indent " " 4}}    chdir: {{ .ChDir }}
{{- Indent " " 0}}   {{- end }}
{{ Indent " " 4}}  {{ with .Args }}
{{ Indent " " 4}}  args:
{{- Indent " " 0}} {{ if .ChDir }}
{{ Indent " " 4}}    chdir: {{ .ChDir }}
{{- Indent " " 0}} {{- end }}
{{- Indent " " 0}} {{ if .Creates }}
{{ Indent " " 4}}    creates: {{ .Creates }}
{{- Indent " " 0}} {{- end }}
{{- Indent " " 0}} {{ if .Removes }}
{{ Indent " " 4}}    removes: {{ .Removes }}
{{- Indent " " 0}} {{- end }}
{{- Indent " " 0}} {{ end }}
`

type AnsibleBuiltinScript struct {
	Args       *args.Args          `json:"args"`
	ChDir      string              `json:"ch_dir"`
	Cmd        string              `json:"cmd"`
	Creates    string              `json:"creates"`
	Decrypt    enumtipe.CostomBool `json:"decrypt"`
	Executable string              `json:"executable"`
	FreeForm   string              `json:"free_form"`
	Removes    string              `json:"removes"`
}

func (a *AnsibleBuiltinScript) String() string {
	return helper.MarshalToString(a)
}

func (a *AnsibleBuiltinScript) MakeAnsibleTask() (string, error) {

	funcMap := template.FuncMap{
		"Indent": strings.Repeat,
	}

	tmpl := template.Must(template.New("ansible_builtin_script").Funcs(funcMap).Parse(Tmplt_ansible_builtin_script))
	var buff bytes.Buffer

	err := tmpl.Execute(&buff, *a)
	if err != nil {
		return "", err
	}

	return buff.String(), nil
}
