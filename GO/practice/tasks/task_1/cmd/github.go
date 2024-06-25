package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func github_procces() {
	fmt.Println("LOL, Github api and gitlab api look alike. It's better to learn something new")
}

var githubCmd = &cobra.Command{
	Use:   "github",
	Short: "Usage github command",
	Run: func(cmd *cobra.Command, args []string) {
		if !validActions[action] {
			cmd.Help()
			fmt.Println("Incorrect value for ACTION")
			os.Exit(1)
		}
		github_procces()
	},
}
