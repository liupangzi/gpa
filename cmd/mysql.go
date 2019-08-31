package cmd

import (
	"github.com/spf13/cobra"
)

var mysqlCmd = &cobra.Command{
	Use:   "mysql",
	Short: "gpa mysql suite",
	Long:  "gpa mysql suite",
}

func init() {
	mysqlCmd.AddCommand(modelCmd, implCmd)
}
