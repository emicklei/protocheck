package main

import "flag"

var (
	targetDir = flag.String("o", ".", "output directory")
)

func main() {
	flag.Parse()
}
