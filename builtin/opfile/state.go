package opfile

import "database/sql/driver"

type CostomState string

const (
	CostomStateAbsent    CostomState = "absent"
	CostomStateDirectory CostomState = "directory"
	CostomStateFile      CostomState = "file"
	CostomStateHard      CostomState = "hard"
	CostomStateLink      CostomState = "link"
	CostomStateTouch     CostomState = "touch"
)

var state_values = []string{
	CostomStateAbsent.String(),
	CostomStateDirectory.String(),
	CostomStateFile.String(),
	CostomStateHard.String(),
	CostomStateLink.String(),
	CostomStateTouch.String(),
}

func (CostomState) Values() (sts []string) {
	sts = append(sts, state_values...)
	return
}

func (cb CostomState) Value() (driver.Value, error) {
	return cb.String(), nil
}

func (cb CostomState) String() string {
	return string(cb)
}
