package enumtipe

import "database/sql/driver"

type YumDnfUseBackend string

const (
	YumDnfUseBackendAuto YumDnfUseBackend = "auto"
	YumDnfUseBackendYum  YumDnfUseBackend = "yum"
	YumDnfUseBackendYum4 YumDnfUseBackend = "yum4"
	YumDnfUseBackendDnf  YumDnfUseBackend = "dnf"
	YumDnfUseBackendDnf4 YumDnfUseBackend = "dnf4"
	YumDnfUseBackendDnf5 YumDnfUseBackend = "dnf5"
)

var yd_ub_values = []string{
	YumDnfUseBackendAuto.String(),
	YumDnfUseBackendYum.String(),
	YumDnfUseBackendYum4.String(),
	YumDnfUseBackendDnf.String(),
	YumDnfUseBackendDnf4.String(),
	YumDnfUseBackendDnf5.String(),
}

func (YumDnfUseBackend) Values() (svs []string) {
	svs = append(svs, yd_ub_values...)
	return
}

func (cb YumDnfUseBackend) Value() (driver.Value, error) {
	return cb.String(), nil
}

func (cb YumDnfUseBackend) String() string {
	return string(cb)
}
