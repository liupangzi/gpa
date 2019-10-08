package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:     "gpa",
	Short:   "Golang Persistence API",
	Long:    "Golang programming interface that describes the management of MySQL data in applications.",
	Version: "v1.0.2",
}

func init() {
	RootCmd.AddCommand(mysqlCmd)
}
