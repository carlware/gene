package generator

import (
	"carlware/gene/internal/models"
	"carlware/gene/internal/utils"
	"fmt"
	"path"
)

func ExecGeneric(oper *models.Operation, doc *models.Document) {
	op := &models.Generic{
		Name:       oper.Name,
		Template:   path.Join(doc.TemplatesPath, oper.Generic.Template),
		Properties: oper.Generic.Properties,
		Model:      &doc.Model,
		Actions:    doc.Actions,
	}
	op.Model.Fields = ExcludeKeys(doc.Model.Fields, oper.Generic.ExcludeFields)

	dst, err := utils.EvalString(oper.Generic.Path, op)
	if err != nil {
		fmt.Println("error while evaluating the path", oper.Generic.Path)
		panic(err)
	}
	op.Path = path.Join(doc.Destination, string(dst))

	err = Generate(op, op.Name, op.Template, op.Path)
	if err != nil {
		fmt.Println("error while generating files", err)
		panic(err)
	}

	fmt.Println("files generated successfully")
}
