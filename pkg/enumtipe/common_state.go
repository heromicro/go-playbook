package enumtipe

import "database/sql/driver"

type CommonState string

const (
	CommonStateAbsent  CommonState = "absent"
	CommonStatePresent CommonState = "present"
	CommonStateLatest  CommonState = "latest"
)

var cstate_values = []string{
	CommonStateAbsent.String(),
	CommonStatePresent.String(),
	CommonStateLatest.String(),
}

func (CommonState) Values() (svs []string) {
	svs = append(svs, cstate_values...)
	return
}

func (cb CommonState) Value() (driver.Value, error) {
	return cb.String(), nil
}

func (cb CommonState) String() string {
	return string(cb)
}
