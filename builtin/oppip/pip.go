package oppip

import (
	"bytes"
	"strings"
	"text/template"

	"github.com/heromicro/go-playbook/pkg/enumtipe"
	"github.com/heromicro/go-playbook/pkg/helper"
)

var Tmplt_ansible_builtin_pip = `
{{ Indent " " 4}}- name: pip install
{{ Indent " " 4}}  ansible.builtin.pip:
{{ Indent " " 4}}    name: {{- range $index, $word := .Name }}
{{ Indent " " 4}}      - {{ $word }} 
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .Chdir }}
{{ Indent " " 4}}    chdir: {{ .Chdir }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .Editable }}
{{ Indent " " 4}}    editable: {{ .Editable }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .Executable }}
{{ Indent " " 4}}    executable: {{ .Executable }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .ExtraArgs }}
{{ Indent " " 4}}    extra_args: {{ .ExtraArgs }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .State }}
{{ Indent " " 4}}    state: {{ .State }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .Umask }}
{{ Indent " " 4}}    umask: {{ .Umask }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .Version }}
{{ Indent " " 4}}    version: {{ .Version }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .Virtualenv }}
{{ Indent " " 4}}    virtualenv: {{ .Virtualenv }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .VirtualenvCommand }}
{{ Indent " " 4}}    virtualenv_command: {{ .VirtualenvCommand }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .VirtualenvPython }}
{{ Indent " " 4}}    virtualenv_python: {{ .VirtualenvPython }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .VirtualenvSitePackages }}
{{ Indent " " 4}}    virtualenv_site_packages: {{ .VirtualenvSitePackages }}
{{- Indent " " 0}}   {{- end }}
`

type AnsibleBuiltinPip struct {
	Name                   []string            `json:"name"`
	Chdir                  string              `json:"chdir"`
	Editable               enumtipe.CostomBool `json:"editable"`
	Executable             string              `json:"executable"`
	ExtraArgs              string              `json:"extra_args"`
	Requirements           string              `json:"requirements"`
	State                  PipState            `json:"state"`
	Umask                  string              `json:"umask"`
	Version                string              `json:"version"`
	Virtualenv             string              `json:"virtualenv"`
	VirtualenvCommand      string              `json:"virtualenv_command"`
	VirtualenvPython       string              `json:"virtualenv_python"`
	VirtualenvSitePackages string              `json:"virtualenv_site_packages"`
}

func (a *AnsibleBuiltinPip) String() string {
	return jsonhelper.MarshalToString(a)
}

func (a *AnsibleBuiltinPip) MakeAnsibleTask() (string, error) {

	// var tmpl *template.Template
	funcMap := template.FuncMap{
		"Indent": strings.Repeat,
	}

	tmpl := template.Must(template.New("ansible_builtin_pip").Funcs(funcMap).Parse(Tmplt_ansible_builtin_pip))
	var buff bytes.Buffer

	err := tmpl.Execute(&buff, *a)
	if err != nil {
		return "", err
	}

	return buff.String(), nil
}
