package main

import (
	"carlware/gene/internal/generator"
	"fmt"
)

func main() {
	err := generator.Generate("./internal/design/gene.yml", "./templates/model.tpml", "./src/Models/certificate/certificat.ts")
	if err != nil {
		fmt.Println("error", err)
	}
}
