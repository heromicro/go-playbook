package opfirewalld

import "database/sql/driver"

type FirewalldTarget string

const (
	FirewalldTargetAbsent  FirewalldTarget = "absent"
	FirewalldTargetDefault FirewalldTarget = "default"
	FirewalldTargetAccept  FirewalldTarget = "ACCEPT"
	FirewalldTargetDrop    FirewalldTarget = "DROP"
	FirewalldTargetReject  FirewalldTarget = "\"%%REJECT%%\""
)

var fw_target_values = []string{
	FirewalldTargetAbsent.String(),
	FirewalldTargetDefault.String(),
	FirewalldTargetAccept.String(),
	FirewalldTargetDrop.String(),
	FirewalldTargetReject.String(),
}

func (FirewalldTarget) Values() (svs []string) {
	svs = append(svs, fw_target_values...)
	return
}

func (cb FirewalldTarget) Value() (driver.Value, error) {
	return cb.String(), nil
}

func (cb FirewalldTarget) String() string {
	return string(cb)
}
