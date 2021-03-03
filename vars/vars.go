package vars

var Debug bool
var Silent bool
var SSHDefaultPort = 0

func init() {
	Debug = false
	Silent = false
	SSHDefaultPort = 22
}
