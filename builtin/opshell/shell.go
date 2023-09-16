package opshell

import (
	"bytes"
	"text/template"

	"github.com/heromicro/go-playbook/pkg/args"
	"github.com/heromicro/go-playbook/pkg/enumtipe"
	"github.com/heromicro/go-playbook/pkg/helper"
)

// https://docs.ansible.com/ansible/latest/collections/ansible/builtin/shell_module.html
var Tmplt_ansible_builtin_shell = `
{{ Indent " " 4}}- name: shell
{{ Indent " " 4}}  ansible.builtin.shell: 
{{- Indent " " 0}}   {{- if .Cmd }} 
{{ Indent " " 4}}    cmd: {{ .Cmd }} 
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .ChDir }}
{{ Indent " " 4}}    chdir: {{ .ChDir }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .Creates }}
{{ Indent " " 4}}    creates: {{ .Creates }}
{{- Indent " " 0}}   {{- end }}
{{ Indent " " 4}}  {{ with .Args }}
{{ Indent " " 4}}  args:
{{- Indent " " 0}} {{ if .ChDir }}
{{ Indent " " 4}}    chdir: {{ .ChDir }}
{{- Indent " " 0}} {{- end }}
{{- Indent " " 0}} {{ if .Creates }}
{{ Indent " " 4}}    creates: {{ .Creates }}
{{- Indent " " 0}} {{- end }}
{{- Indent " " 0}} {{ if .Executable }}
{{ Indent " " 4}}    executable: {{ .Executable }}
{{- Indent " " 0}} {{- end }}
{{- Indent " " 0}} {{ end }}
`

type AnsibleBuiltinShell struct {
	Args            *args.Args          `json:"args"`
	ChDir           string              `json:"ch_dir"`
	Cmd             string              `json:"cmd"`
	Creates         string              `json:"creates"`
	Executable      string              `json:"executable"`
	FreeForm        string              `json:"free_form"`
	Removes         string              `json:"removes"`
	Stdin           string              `json:"stdin"`
	StdinAddNewline enumtipe.CostomBool `json:"stdin_add_newline"`
}

func (a *AnsibleBuiltinShell) String() string {
	return helper.MarshalToString(a)
}

func (a *AnsibleBuiltinShell) MakeAnsibleTask() (string, error) {

	// var tmpl *template.Template
	funcMap := template.FuncMap{
		"Indent": strings.Repeat,
	}

	tmpl := template.Must(template.New("ansible_builtin_shell").Funcs(funcMap).Parse(Tmplt_ansible_builtin_shell))
	var buff bytes.Buffer

	err := tmpl.Execute(&buff, *a)
	if err != nil {
		return "", err
	}

	return buff.String(), nil
}
