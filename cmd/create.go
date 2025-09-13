package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a crud",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {

		filePath := args[0]
		dat, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Errorf("Erro ao abrir arquivo", err)
		}

		fmt.Print(string(dat))

	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
