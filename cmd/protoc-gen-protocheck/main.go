package main

import (
	"flag"
	"log/slog"
	"strings"

	"github.com/emicklei/protocheck/cmd/protoc-gen-protocheck/lang/golang"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/descriptorpb"
)

var lang *string

func main() {
	var flags flag.FlagSet
	lang = flags.String("lang", "go", "one of the supported programming languages")
	opts := &protogen.Options{
		ParamFunc: flags.Set,
	}
	opts.Run(run)
}

// https://rotemtam.com/2021/03/22/creating-a-protoc-plugin-to-gen-go-code/
func run(p *protogen.Plugin) error {
	p.SupportedFeatures = p.SupportedFeatures + 1 + 2 // FEATURE_PROTO3_OPTIONAL, FEATURE_SUPPORTS_EDITIONS
	p.SupportedEditionsMinimum = descriptorpb.Edition_EDITION_PROTO3
	p.SupportedEditionsMaximum = descriptorpb.Edition_EDITION_2023

	for _, each := range p.Files {
		if !each.Generate {
			continue
		}
		if !importsCheck(each) {
			continue
		}
		slog.Debug("destination",
			"pkg", each.GoPackageName,
			"path", each.Desc.Path(),
			"prefix", each.GeneratedFilenamePrefix,
		)

		switch *lang {
		case "go":
			if err := golang.Process(p, each); err != nil {
				return err
			}
		default:
			slog.Warn("unsupported language", "lang", *lang)
		}
	}
	return nil
}

// importsCheck returns whether the file imports check.proto.
func importsCheck(f *protogen.File) bool {
	for i := 0; i < f.Desc.Imports().Len(); i++ {
		entry := f.Desc.Imports().Get(i)
		if strings.HasSuffix(entry.Path(), "check.proto") {
			return true
		}
	}
	return false
}
