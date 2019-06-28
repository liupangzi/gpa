package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var mysqlCmd = &cobra.Command{
	Use:   "mysql",
	Short: "mysql",
	Long:  "mysql",
}

var modelCmd = &cobra.Command{
	Use:   "model",
	Short: "Generate MySQL model from SQL ddl",
	Long:  "Generate MySQL model from SQL ddl",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("model called")
	},
}

func init() {
	mysqlCmd.AddCommand(modelCmd)
}
