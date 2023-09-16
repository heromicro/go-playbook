package opuser

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/heromicro/go-playbook/pkg/enumtipe"
	"github.com/heromicro/go-playbook/pkg/helper"
)

// https://docs.ansible.com/ansible/latest/collections/ansible/builtin/user_module.html
var Tmplt_ansible_builtin_user = `
{{ Indent " " 4}}- name: Add Systemuser {{ .Username }} {{ if .UID }}with uid {{ .UID }} {{ end }}
{{ Indent " " 4}}  ansible.builtin.user:
{{ Indent " " 4}}    name: {{ .Username }}
{{- Indent " " 0}}   {{- if .Password }}
{{ Indent " " 4}}    password: {{ .Password }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .UpdatePassword }}
{{ Indent " " 4}}    update_password: {{ .UpdatePassword }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .Home }}
{{ Indent " " 4}}    home: {{ .Home }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .CreateHome }}
{{ Indent " " 4}}    create_home: {{ .CreateHome }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .MoveHome }}
{{ Indent " " 4}}    move_home: {{ .MoveHome }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .Group }}
{{ Indent " " 4}}    group: {{ .Group }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .Comment }}
{{ Indent " " 4}}    comment: {{ .Comment }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .UID }}
{{ Indent " " 4}}    uid: {{ .UID }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .Shell }}
{{ Indent " " 4}}    shell: {{ .Shell }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .Groups }}
{{ Indent " " 4}}    groups: {{ .Groups }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .Append }}
{{ Indent " " 4}}    append: {{ .Append }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .Authorization }}
{{ Indent " " 4}}    authorization: {{ .Authorization }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .Expires }}
{{ Indent " " 4}}    expires: {{ .Expires }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .Force }}
{{ Indent " " 4}}    force: {{ .Force }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .GenerateSshKey }}
{{ Indent " " 4}}    generate_ssh_key: {{ .GenerateSshKey }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .GenerateSshKey }}
{{ Indent " " 4}}    generate_ssh_key: {{ .GenerateSshKey }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .Hidden }}
{{ Indent " " 4}}    hidden: {{ .Hidden }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .Local }}
{{ Indent " " 4}}    local: {{ .Local }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .LoginClass }}
{{ Indent " " 4}}    login_class: {{ .LoginClass }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .NonUnique }}
{{ Indent " " 4}}    non_unique: {{ .NonUnique }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .PasswordExpireMax }}
{{ Indent " " 4}}    password_expire_max: {{ .PasswordExpireMax }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .PasswordExpireMin }}
{{ Indent " " 4}}    password_expire_min: {{ .PasswordExpireMin }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .PasswordLock }}
{{ Indent " " 4}}    password_lock: {{ .PasswordLock }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .Profile }}
{{ Indent " " 4}}    profile: {{ .Profile }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .Remove }}
{{ Indent " " 4}}    remove: {{ .Remove }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .Role }}
{{ Indent " " 4}}    role: {{ .Role }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .SeUser }}
{{ Indent " " 4}}    seuser: {{ .SeUser }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .Skeleton }}
{{ Indent " " 4}}    skeleton: {{ .Skeleton }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .SshKeyBits }}
{{ Indent " " 4}}    ssh_key_bits: {{ .SshKeyBits }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .SshKeyComment }}
{{ Indent " " 4}}    ssh_key_comment: {{ .SshKeyComment }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .SshKeyFile }}
{{ Indent " " 4}}    ssh_key_file: {{ .SshKeyFile }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .SshKeyPassphrase }}
{{ Indent " " 4}}    ssh_key_passphrase: {{ .SshKeyPassphrase }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .SshKeyType }}
{{ Indent " " 4}}    ssh_key_type: {{ .SshKeyType }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .State }}
{{ Indent " " 4}}    state: {{ .State }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .System }}
{{ Indent " " 4}}    system: {{ .System }}
{{- Indent " " 0}}   {{- end }}
{{- Indent " " 0}}   {{- if .Umask }}
{{ Indent " " 4}}    umask: {{ .Umask }}
{{- Indent " " 0}}   {{- end }}
`

type AnsibleBuiltinUser struct {
	Username          string                        `json:"username"`
	Password          string                        `json:"password"`
	UpdatePassword    enumtipe.UpdatePasswordChoice `json:"update_password"`
	PasswordExpireMax *int64                        `json:"password_expire_max"`
	PasswordExpireMin *int64                        `json:"password_expire_min"`
	CreateHome        enumtipe.CostomBool           `json:"create_home"`
	MoveHome          enumtipe.CostomBool           `json:"move_home"`
	PasswordLock      enumtipe.CostomBool           `json:"password_lock"`
	NonUnique         enumtipe.CostomBool           `json:"non_unique"`
	Profile           *string                       `json:"profile"`
	Group             string                        `json:"group"`
	Home              string                        `json:"home"`
	Skeleton          string                        `json:"skeleton"`
	Comment           string                        `json:"comment"`
	Shell             string                        `json:"shell"`
	UID               *uint32                       `json:"uid"`
	Umask             string                        `json:"umask"`
	Groups            string                        `json:"groups"`
	Append            enumtipe.CostomBool           `json:"append"`
	Remove            enumtipe.CostomBool           `json:"remove"`
	Force             enumtipe.CostomBool           `json:"force"`
	GenerateSshKey    enumtipe.CostomBool           `json:"generate_ssh_key"`
	SshKeyBits        *int64                        `json:"ssh_key_bits"`
	SshKeyComment     string                        `json:"ssh_key_comment"`
	SshKeyFile        string                        `json:"ssh_key_file"`
	SshKeyPassphrase  string                        `json:"ssh_key_passphrase"`
	SshKeyType        string                        `json:"ssh_key_type"`
	State             enumtipe.CommonState          `json:"state"`
	Hidden            enumtipe.CostomBool           `json:"hidden"`
	Local             enumtipe.CostomBool           `json:"local"`
	System            enumtipe.CostomBool           `json:"system"`
	LoginClass        string                        `json:"login_class"`
	Authorization     string                        `json:"authorization"`
	Role              string                        `json:"role"`
	SeUser            string                        `json:"seuser"`
	Expires           *float64                      `json:"expires"`
}

func (a *AnsibleBuiltinUser) String() string {
	return helper.MarshalToString(a)
}

func (a *AnsibleBuiltinUser) MakeAnsibleTask() (string, error) {

	funcMap := template.FuncMap{
		"Indent": strings.Repeat,
	}

	tmpl := template.Must(template.New("ansible_builtin_user").Funcs(funcMap).Parse(Tmplt_ansible_builtin_user))
	var buff bytes.Buffer

	err := tmpl.Execute(&buff, *a)
	if err != nil {
		return "", err
	}

	return buff.String(), nil
}
