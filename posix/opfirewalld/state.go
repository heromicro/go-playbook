package opfirewalld

import "database/sql/driver"

type FirewalldState string

const (
	FirewalldStateAbsent   FirewalldState = "absent"
	FirewalldStatePresent  FirewalldState = "present"
	FirewalldStateEnabled  FirewalldState = "enabled"
	FirewalldStateDisabled FirewalldState = "disabled"
)

var cstate_values = []string{
	FirewalldStateAbsent.String(),
	FirewalldStatePresent.String(),
	FirewalldStateEnabled.String(),
	FirewalldStateDisabled.String(),
}

func (FirewalldState) Values() (svs []string) {
	svs = append(svs, cstate_values...)
	return
}

func (cb FirewalldState) Value() (driver.Value, error) {
	return cb.String(), nil
}

func (cb FirewalldState) String() string {
	return string(cb)
}
