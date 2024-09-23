package main

import (
	"flag"

	"google.golang.org/protobuf/compiler/protogen"
)

var (
	targetDir = flag.String("o", ".", "output directory")
)

func main() {
	var flags flag.FlagSet
	opts := &protogen.Options{
		ParamFunc: flags.Set,
	}
	opts.Run(run)
}

func run(p *protogen.Plugin) error {
	p.SupportedFeatures = p.SupportedFeatures + 1 // FEATURE_PROTO3_OPTIONAL
	return nil
}
