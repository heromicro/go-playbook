package opfirewalld

import (
	"bytes"
	"strings"
	"text/template"

	"github.com/heromicro/go-playbook/pkg/enumtipe"
	"github.com/heromicro/go-playbook/pkg/helper"
)

var Tmplt_ansible_posix_firewalld = `
{{ Indent " " 4}}- name: {{ .Name }}  
{{ Indent " " 4}}  ansible.posix.firewalld:
{{- Indent " " 0}}   {{- if .IcmpBlock }}
{{ Indent " " 4}}    icmp_block: {{ .IcmpBlock }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .IcmpBlockInversion }}
{{ Indent " " 4}}    icmp_block_inversion: {{ .IcmpBlockInversion }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .Service }}
{{ Indent " " 4}}    service: {{ .Service }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .Immediate }}
{{ Indent " " 4}}    immediate: {{ .Immediate }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .Interface }}
{{ Indent " " 4}}    interface: {{ .Interface }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .Masquerade }}
{{ Indent " " 4}}    masquerade: {{ .Masquerade }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .Offline }}
{{ Indent " " 4}}    offline: {{ .Offline }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .Permanent }}
{{ Indent " " 4}}    permanent: {{ .Permanent }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .Port }}
{{ Indent " " 4}}    port: {{ .Port }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .Protocol }}
{{ Indent " " 4}}    protocol: {{ .Protocol }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .RichRule }}
{{ Indent " " 4}}    rich_rule: {{ .RichRule }}
{{- Indent " " 0}}   {{- end }} 
{{- Indent " " 0}}   {{- if .Source }}
{{ Indent " " 4}}    source: {{ .Source }}
{{- Indent " " 0}}   {{- end }}
{{ Indent " " 4}}    state: {{ .State }}
{{- Indent " " 0}}   {{- if .Target }}
{{ Indent " " 4}}    target: {{ .Target }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .Timeout }}
{{ Indent " " 4}}    timeout: {{ .Timeout }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .Zone }}
{{ Indent " " 4}}    zone: {{ .Zone }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .PortForward }}
{{ Indent " " 4}}    port_forward: {{ range .PortForward }}
{{ Indent " " 4}}      - port: {{ .Port }}
{{ Indent " " 4}}        proto: {{ .Proto }}
{{- Indent " " 0}}       {{- if .Toaddr }}
{{ Indent " " 4}}        toaddr: {{ .Toaddr }}
{{- Indent " " 0}}       {{- end }}
{{ Indent " " 4}}        toport: {{ .Toport }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- end }}
`

type AnsiblePosixFirewalld struct {
	Name               string              `json:"name"`
	IcmpBlock          string              `json:"icmp_block"`
	IcmpBlockInversion string              `json:"icmp_block_inversion"`
	Immediate          enumtipe.CostomBool `json:"immediate"`
	Interface          string              `json:"interface"`
	Masquerade         string              `json:"masquerade"`
	Offline            enumtipe.CostomBool `json:"offline"`
	Permanent          enumtipe.CostomBool `json:"permanent"`
	Port               string              `json:"port"`
	PortForward        []*PortForward      `json:"port_forward"`
	Protocol           string              `json:"protocol"`
	RichRule           string              `json:"rich_rule"`
	Service            string              `json:"service"`
	Source             string              `json:"source"`
	State              FirewalldState      `json:"state"`
	Target             FirewalldTarget     `json:"target"`
	Timeout            int32               `json:"timeout"`
	Zone               string              `json:"zone"`
}

func (a *AnsiblePosixFirewalld) String() string {
	return helper.MarshalToString(a)
}

func (a *AnsiblePosixFirewalld) MakeAnsibleTask() (string, error) {

	// var tmpl *template.Template
	funcMap := template.FuncMap{
		"Indent": strings.Repeat,
	}

	tmpl := template.Must(template.New("ansible_posix_firewalld").Funcs(funcMap).Parse(Tmplt_ansible_posix_firewalld))
	var buff bytes.Buffer

	err := tmpl.Execute(&buff, *a)
	if err != nil {
		return "", err
	}

	return buff.String(), nil
}
