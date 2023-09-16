package enumtipe

import "database/sql/driver"

type UpdatePasswordChoice string

const (
	UpdatePasswordChoiceAlways   UpdatePasswordChoice = "always"
	UpdatePasswordChoiceOnCreate UpdatePasswordChoice = "on_create"
)

var up_choices = []string{
	UpdatePasswordChoiceAlways.String(),
	UpdatePasswordChoiceOnCreate.String(),
}

func (UpdatePasswordChoice) Values() (svs []string) {
	svs = append(svs, up_choices...)
	return
}

func (cb UpdatePasswordChoice) Value() (driver.Value, error) {
	return cb.String(), nil
}

func (cb UpdatePasswordChoice) String() string {
	return string(cb)
}
