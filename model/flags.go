package model

import (
	"flag"
)

type Flags struct {
	Port string
}

func (flags *Flags) InitFlags() {
	port := flag.String("port", "4040", "Specifies the host port of server. Default: 4040.")
	flag.Parse()
	flags.Port = *port
}

var Flag Flags