package cmd

import (
	"github.com/spf13/cobra"
)

var mysqlCmd = &cobra.Command{
	Use:   "mysql",
	Short: "gpa mysql suite",
	Long:  "mysql",
}

func init() {
	mysqlCmd.AddCommand(modelCmd, implCmd)
}
