package opfirewalld

type PortForward struct {
	Port   string `json:"port"`
	Proto  string `json:"proto"`
	Toaddr string `json:"toaddr"`
	Toport string `json:"toport"`
}
