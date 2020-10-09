package generator

import (
	"carlware/gene/internal/models"
	"fmt"
)

func FillActionsWithModel(acts []models.Action, mFields []models.Field) []models.Action {
	actions := []models.Action{}

	for _, act := range acts {
		if act.Inherit {
			fields := ExcludeKeys(mFields, act.Exclude)
			action := models.Action{
				Name:       act.Name,
				Properties: act.Properties,
				Request: models.Request{
					Properties: act.Request.Properties,
					Fields:     append(act.Request.Fields, fields...),
				},
				Response: models.Response{
					Properties: act.Response.Properties,
					Fields:     act.Response.Fields,
				},
			}
			fmt.Println("exclude", act.Exclude)
			fmt.Println("fields", fields)
			actions = append(actions, action)
		} else {
			actions = append(actions, act)
		}
	}
	return actions
}

func Gen(templatePath string) {
	doc, err := ParseTemplate(templatePath)
	if err != nil {
		fmt.Println("error while parsing template")
		panic(err)
	}

	doc.Actions = FillActionsWithModel(doc.Actions, doc.Model.Fields)

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

	fmt.Println("files generated successfully")
}
