package vars

var Debug bool
var ContainerRuntime = ""
var SSHDefaultPort = 0

func init() {
	Debug = false
	ContainerRuntime = "rcli"
	SSHDefaultPort = 22
}
