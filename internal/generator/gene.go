package generator

import (
	"bytes"
	"carlware/gene/internal/models"
	"carlware/gene/internal/utils"
	"html/template"
	"io/ioutil"
	"os"

	"fmt"

	"gopkg.in/yaml.v2"
)

func Generate(document, temp, output string) error {
	documentFile, err := ioutil.ReadFile(document)
	if err != nil {
		fmt.Println("error", err)
	}

	doc := &models.Document{}
	err = yaml.Unmarshal(documentFile, doc)
	if err != nil {
		fmt.Println("error", err)
	}

	tmplString, err := utils.FileToString(temp)
	if err != nil {
		fmt.Println("error", err)
	}

	tmpl, err := template.New("test").Parse(tmplString)
	if err != nil {
		fmt.Println("error", err)
	}
	out := &bytes.Buffer{}

	err = tmpl.Execute(out, doc)
	if err != nil {
		fmt.Println("err", err)
	}

	err = utils.CreateFolders(output)
	if err != nil {
		fmt.Println("cannot create folder", err)
	}
	err = ioutil.WriteFile(output, out.Bytes(), os.ModePerm)
	if err != nil {
		fmt.Println("err", err)
	}

	fmt.Println(doc)
	return nil
}
