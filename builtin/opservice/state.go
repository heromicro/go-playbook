package opservice

import "database/sql/driver"

type ServiceState string

const (
	ServiceStateReloaded  ServiceState = "reloaded"
	ServiceStateRestarted ServiceState = "restarted"
	ServiceStateStarted   ServiceState = "started"
	ServiceStateStopped   ServiceState = "stopped"
)

var sstate_values = []string{
	ServiceStateReloaded.String(),
	ServiceStateRestarted.String(),
	ServiceStateStarted.String(),
	ServiceStateStopped.String(),
}

func (ServiceState) Values() (sts []string) {
	sts = append(sts, sstate_values...)
	return
}

func (cb ServiceState) Value() (driver.Value, error) {
	return cb.String(), nil
}

func (cb ServiceState) String() string {
	return string(cb)
}
