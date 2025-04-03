package main

import (
	"fmt"
	configreader "technical-overview/practical/config_reader"
)

func main() {
	fmt.Println("Example Config Reader")

	yReader := &configreader.YAMLReader{Filename: "../data/config.yaml"}
	jReader := &configreader.JSONReader{Filename: "../data/config.json"}

	fmt.Println(yReader)
	configreader.PrintConfig(yReader)

	fmt.Println()

	fmt.Println(jReader)
	configreader.PrintConfig(jReader)
}
