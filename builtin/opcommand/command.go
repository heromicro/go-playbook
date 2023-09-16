package opcommand

import (
	"bytes"
	"text/template"

	"github.com/heromicro/go-playbook/pkg/args"
	"github.com/heromicro/go-playbook/pkg/enumtipe"
	"github.com/heromicro/go-playbook/pkg/helper"
)

// https://docs.ansible.com/ansible/latest/collections/ansible/builtin/command_module.html
var Tmplt_ansible_builtin_command = `
{{ Indent " " 4}}- name: command 
{{ Indent " " 4}}  ansible.builtin.command: 
{{- Indent " " 0}}   {{- if .Cmd }} 
{{ Indent " " 4}}    cmd: {{ .Cmd }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .Argv }}
{{ Indent " " 4}}    argv: {{- range $index, $word := .Argv }}
{{ Indent " " 4}}      - {{ $word }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .ChDir }} 
{{ Indent " " 4}}    chdir: {{ .ChDir }}
{{- Indent " " 0}}   {{- end }} 
{{- Indent " " 0}}   {{- if .Creates }}
{{ Indent " " 4}}    creates: {{ .Creates }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}} {{- with .Args }}
{{ Indent " " 4}}  args:
{{- Indent " " 0}} {{ if .ChDir }}
{{ Indent " " 4}}    chdir: {{ .ChDir }}
{{- Indent " " 0}} {{- end }}
{{- Indent " " 0}} {{ if .Creates }}
{{ Indent " " 4}}    creates: {{ .Creates }}
{{- Indent " " 0}} {{- end }}
{{- Indent " " 0}} {{ end }}
`

type AnsibleBuiltinCommand struct {
	Args            *args.Args          `json:"args"`
	Argv            []string            `json:"argv"`
	ChDir           string              `json:"ch_dir"`
	Cmd             string              `json:"cmd"`
	Creates         string              `json:"creates"`
	FreeForm        string              `json:"free_form"`
	Removes         string              `json:"removes"`
	Stdin           string              `json:"stdin"`
	StdinAddNewline enumtipe.CostomBool `json:"stdin_add_newline"`
	StripEmptyEnds  enumtipe.CostomBool `json:"strip_empty_ends"`
}

func (a *AnsibleBuiltinCommand) String() string {
	return helper.MarshalToString(a)
}

func (a *AnsibleBuiltinCommand) MakeAnsibleTask() (string, error) {

	// var tmpl *template.Template
	funcMap := template.FuncMap{
		"Indent": strings.Repeat,
	}

	tmpl := template.Must(template.New("ansible_builtin_command").Funcs(funcMap).Parse(Tmplt_ansible_builtin_command))
	var buff bytes.Buffer

	err := tmpl.Execute(&buff, *a)
	if err != nil {
		return "", err
	}

	return buff.String(), nil
}
