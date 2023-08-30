package opfile

import (
	"bytes"
	"text/template"

	"github.com/heromicro/go-playbook/pkg/enumtipe"
	"github.com/heromicro/go-playbook/pkg/helper"
)

// https://docs.ansible.com/ansible/latest/collections/ansible/builtin/copy_module.html
var Tmplt_ansible_copy_file = `
    - name: copy {{- if .Content }} " {{- range $index, $word := .Content }} {{- $word }}  {{- end }}"{{- end }} {{- if .Src }} {{- .Src }} {{- end }} to {{ .Dest }} 
      ansible.builtin.copy:
        {{- if .Content }}
        content: | {{- range $index, $word := .Content }}
          {{ $word }}
        {{- end }}
        {{- end }}
        {{- if .Src }} 
        src: {{ .Src }}
        {{- end }}
        dest: {{ .Dest }}
        {{- if .Owner }}
        owner: {{ .Owner }}
        {{- end }}
        {{- if .Group }}
        group: {{ .Group }} 
        {{- end }}
	      {{- if .Mode }}
        mode: {{ .Mode }}
        {{- end }}
        {{- if .RemoteSrc }}
  	    remote_src: {{ .RemoteSrc }}
        {{- end }}
        {{- if .Follow }}
        follow: {{ .Follow }}
        {{- end }}
        {{- if .Backup }}
        backup: {{ .Backup }}
        {{- end }}
        validate: /usr/sbin/visudo -csf %s
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
}

func (a *AnsibleCopyFile) String() string {
	return helper.MarshalToString(a)
}

func (a *AnsibleCopyFile) MakeAnsibleTask() (string, error) {

	tmpl := template.Must(template.New("ansible_copy_file").Parse(Tmplt_ansible_copy_file))
	var buff bytes.Buffer

	err := tmpl.Execute(&buff, *a)
	if err != nil {
		return "", nil
	}

	return buff.String(), nil
}
