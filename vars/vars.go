package vars

var Debug bool
var ContainerRuntime = ""

func init() {
	Debug = false
	ContainerRuntime = "rcli"
}
