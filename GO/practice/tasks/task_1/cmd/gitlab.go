package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func gitlab_procces() {
	fmt.Println(repositories)
}

var gitlabCmd = &cobra.Command{
	Use:   "gitlab",
	Short: "Usage gitlab command",
	Run: func(cmd *cobra.Command, args []string) {
		if !validActions[action] {
			cmd.Help()
			fmt.Println("Incorrect value for ACTION")
			os.Exit(1)
		}
		gitlab_procces()
	},
}
