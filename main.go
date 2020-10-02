package main

import (
	"carlware/gene/internal/generator"
)

func main() {
	// generator.Generate("./internal/design/gene.yml", "./templates/model.tpml", "./src/Models/certificate/certificat.ts")
	generator.Gen("./internal/design/gene.yml")
}
