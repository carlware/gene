package gene

import (
	"carlware/gene/internal/generator"

	"github.com/spf13/cobra"
)

var templatePath string

var GeneCmd = &cobra.Command{
	Use:     "gen",
	Aliases: []string{"s"},
	Short:   "Generates files from templates",
	Run: func(cmd *cobra.Command, args []string) {
		generator.Gen(templatePath)
	},
}

func init() {
	GeneCmd.Flags().StringVarP(&templatePath, "template", "t", "", "Template YAML File")
}
