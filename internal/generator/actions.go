package generator

import (
	"carlware/gene/internal/models"
	"carlware/gene/internal/utils"
	"fmt"
	"path"
)

func ExecActions(oper *models.Operation, doc *models.Document) {

	for _, act := range doc.Actions {
		process := &models.Process{
			Model:   doc.Model,
			Actions: doc.Actions,
			Action:  &act,
			Name:    act.Name,
		}

		dst, err := utils.EvalString(oper.Process.Path, process)
		if err != nil {
			fmt.Println("error while evaluating the path", oper.Process.Path)
			panic(err)
		}
		name, err := utils.EvalString(oper.Process.Filename, process)
		if err != nil {
			fmt.Println("error while evaluating the filename", oper.Process.Filename)
			panic(err)
		}
		dstPath := path.Join(doc.Destination, string(dst), string(name))
		tempPath := path.Join(doc.TemplatesPath, oper.Process.Template)

		err = Generate(process, act.Name, tempPath, dstPath)
		if err != nil {
			fmt.Println("error while generating files", dstPath, tempPath)
			panic(err)
		}

		fmt.Println("action name", act.Name)
	}
}
