package enumtipe

import "database/sql/driver"

type YumDnfState string

const (
	YumDnfStateAbsent    YumDnfState = "absent"
	YumDnfStateInstalled YumDnfState = "installed"
	YumDnfStateLatest    YumDnfState = "latest"
	YumDnfStatePresent   YumDnfState = "present"
	YumDnfStateRemoved   YumDnfState = "removed"
)

var ydstate_values = []string{
	YumDnfStateAbsent.String(),
	YumDnfStateInstalled.String(),
	YumDnfStateLatest.String(),
	YumDnfStatePresent.String(),
	YumDnfStateRemoved.String(),
}

func (YumDnfState) Values() (svs []string) {
	svs = append(svs, ydstate_values...)
	return
}

func (cb YumDnfState) Value() (driver.Value, error) {
	return cb.String(), nil
}

func (cb YumDnfState) String() string {
	return string(cb)
}
