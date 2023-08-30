package opfile

import (
	"bytes"
	"text/template"

	"github.com/heromicro/go-playbook/pkg/enumtipe"
	"github.com/heromicro/go-playbook/pkg/helper"
)

// https://docs.ansible.com/ansible/latest/collections/ansible/builtin/file_module.html
var Tmplt_ansible_file = `
    - name: {{ .Name }} 
      ansible.builtin.file:
        {{.Path}}
        {{- if .State }}
        state: {{ .State }}
        {{- end }}
        {{- if .Securse }}
        recurse: {{ .Securse }}
        {{- end }}
        {{- if .Owner }}
        owner: {{ .Owner }}
        {{- end }}
        {{- if .Group }}
        group: {{ .Group }} 
        {{- end }}
        {{- if .Mode }}
        mode: {{ .Mode }}
        {{- end }}
        {{- if .Src }}
        src: {{ .Src }}
        {{- end }}
        {{- if .Dest }}
        dest: {{ .Dest }}
        {{- end }}
        {{- if .AccessTime }}
        dest: {{ .AccessTime }}
        {{- end }}
        {{- if .AccessTimeFormat }}
        dest: {{ .AccessTimeFormat }}
        {{- end }}
        {{- if .ModificationTime }}
        dest: {{ .ModificationTime }}
        {{- end }}
        {{- if .ModificationTimeFormat }}
        dest: {{ .ModificationTimeFormat }}
        {{- end }}
        {{- if .Attributes }}
        dest: {{ .Attributes }}
        {{- end }}
        {{- if .Follow }}
        dest: {{ .Follow }}
        {{- end }}
        {{- if .Force }}
        dest: {{ .Force }}
        {{- end }}
        {{- if .UnsafeWrites }}
        dest: {{ .UnsafeWrites }}
        {{- end }}
        {{- if .SeLevel }}
        dest: {{ .SeLevel }}
        {{- end }}
        {{- if .SeRole }}
        dest: {{ .SeRole }}
        {{- end }}
        {{- if .SeUser }}
        dest: {{ .SeUser }}
        {{- end }}
`

type AnsibleBuiltinFile struct {
	Name    string              `json:"name"`
	Path    string              `json:"path"`
	Recurse enumtipe.CostomBool `json:"recurse"`
	State   FileState           `json:"state"`
	Owner   string              `json:"owner"`
	Group   string              `json:"group"`
	Mode    string              `json:"mode"` // '0644', u+rw,g-wx,o-rwx
	Src     string              `json:"src"`
	Dest    string              `json:"dest"`

	AccessTime             string `json:"access_time"`
	AccessTimeFormat       string `json:"access_time_format"`
	ModificationTime       string `json:"modification_time"`
	ModificationTimeFormat string `json:"modification_time_format"`

	Attributes   string              `json:"attributes"`
	Follow       enumtipe.CostomBool `json:"follow"`
	Force        enumtipe.CostomBool `json:"force"`
	UnsafeWrites enumtipe.CostomBool `json:"unsafe_writes"`

	SeLevel string `json:"selevel"`
	SeRole  string `json:"serole"`
	SeUser  string `json:"seuser"`
}

func (a *AnsibleBuiltinFile) String() string {
	return helper.MarshalToString(a)
}

func (a *AnsibleBuiltinFile) MakeAnsibleTask() (string, error) {

	tmpl := template.Must(template.New("ansible_builtin_gather_facts").Parse(Tmplt_ansible_file))
	var buff bytes.Buffer

	err := tmpl.Execute(&buff, *a)
	if err != nil {
		return "", err
	}

	return buff.String(), nil
}
