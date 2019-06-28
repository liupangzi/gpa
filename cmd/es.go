package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var esCmd = &cobra.Command{
	Use:   "es",
	Short: "ElasticSearch API",
	Long:  "ElasticSearch API",
}

var mappingCmd = &cobra.Command{
	Use:   "mapping",
	Short: "Generate ElasticSearch mappings",
	Long:  "Generate ElasticSearch mappings",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mapping called", cmd, args)
	},
}

func init() {
	esCmd.AddCommand(mappingCmd)
}
