package generator

import (
	"fmt"
)

func Gen(templatePath string) {
	doc, err := ParseTemplate(templatePath)
	if err != nil {
		fmt.Println("error while parsing template")
		panic(err)
	}

	for _, oper := range doc.Operations {
		switch oper.Type {
		case "generic":
			ExecGeneric(&oper, doc)
		case "process":
			ExecActions(&oper, doc)
		default:
			fmt.Printf("operation '%s': is not valid.", oper.Type)
		}
	}

}
