package cli

import "flag"

var (
	Cors    *bool
	Port    *string
	Version *bool
)

func ParseArgs() {
	Cors = flag.Bool("cors", false, "enable CORS for dev")
	Port = flag.String("port", ":9000", "change :port or ip:port")
	Version = flag.Bool("version", false, "print poc-go-svelte-websockets version")
	flag.Parse()
}
