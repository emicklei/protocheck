package main

import (
	"flag"
	"log/slog"

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
	for _, each := range p.Files {
		if !importsCheck(each) {
			continue
		}
		for _, other := range each.Messages {
			slog.Info("message", "name", other.Desc.Name(), "pkg", each.GoPackageName)
		}
	}
	return nil
}

// importsCheck returns whether the file imports check.proto.
func importsCheck(f *protogen.File) bool {
	for i := 0; i < f.Desc.Imports().Len(); i++ {
		entry := f.Desc.Imports().Get(i)
		if entry.Path() == "check.proto" {
			return true
		}
	}
	return false
}
