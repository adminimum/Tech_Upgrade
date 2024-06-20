package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	action       string
	repositories string
)

var validActions = map[string]bool{
	"create": true,
	"delete": true,
	"copy":   true,
	"rename": true,
	"list":   true,
}

var rootCmd = &cobra.Command{
	Use:   "API_GIT",
	Short: "API_GIT is an app for managing repositories",
	Long: `App for manupilating repositories on different git providers:
	GITHUB/GITLAB`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	gitlabCmd.Flags().StringVarP(&action, "action", "a", "", "Action with repo:['create', 'delete', 'copy', 'rename', 'list']")
	githubCmd.Flags().StringVarP(&action, "action", "a", "", "Action with repo:['create', 'delete', 'copy', 'rename', 'list']")
	gitlabCmd.PersistentFlags().StringVarP(&repositories, "repos", "r", "", "List of repositories names hyphenated")
	githubCmd.PersistentFlags().StringVarP(&repositories, "repos", "r", "", "List of repositories names hyphenated")
	rootCmd.AddCommand(gitlabCmd, githubCmd)

	gitlabCmd.MarkFlagRequired("action")
	githubCmd.MarkFlagRequired("action")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
