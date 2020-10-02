package cmd

import (
	"carlware/gene/cli/cmd/gene"
	"carlware/gene/cli/cmd/version"

	"github.com/spf13/cobra"
)

// RootCmd describes root command of the tool
var mainCmd = &cobra.Command{
	Use:   "gene",
	Short: "Very Simple Template Generator",
}

func init() {
	mainCmd.AddCommand(version.VersionCmd)
	mainCmd.AddCommand(gene.GeneCmd)
}

// Execute main command
func Execute() error {
	return mainCmd.Execute()
}
