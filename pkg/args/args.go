package args

type Args struct {
	Creates    string `json:"creates"`
	ChDir      string `json:"ch_dir"`
	Removes    string `json:"removes"`
	Executable string `json:"executable"`
}
