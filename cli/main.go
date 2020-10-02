package main

import "carlware/gene/cli/cmd"

func main() {
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
