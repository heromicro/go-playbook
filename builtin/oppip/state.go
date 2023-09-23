package oppip

import "database/sql/driver"

type PipState string

const (
	PipStateAbsent         PipState = "absent"
	PipStateForcereinstall PipState = "forcereinstall"
	PipStateLatest         PipState = "latest"
	PipStatePresent        PipState = "present"
)

var pipstate_values = []string{
	PipStateAbsent.String(),
	PipStateForcereinstall.String(),
	PipStateLatest.String(),
	PipStatePresent.String(),
}

func (PipState) Values() (svs []string) {
	svs = append(svs, pipstate_values...)
	return
}

func (cb PipState) Value() (driver.Value, error) {
	return cb.String(), nil
}

func (cb PipState) String() string {
	return string(cb)
}
