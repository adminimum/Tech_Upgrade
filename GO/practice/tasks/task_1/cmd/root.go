package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var conf_file_path = "conf/config.json"

var (
	action       string
	repositories string
	description  string
	path         string
	namespace_id int
	project_id   int
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
	gitlabCmd.Flags().StringVarP(&description, "descript", "d", "", "Description for created repository")
	gitlabCmd.Flags().StringVarP(&path, "path", "p", "", "URL path for created repository")
	gitlabCmd.Flags().IntVarP(&namespace_id, "namespace", "n", 0, "Numerical id of gitlab namespace for creared repository")
	gitlabCmd.Flags().IntVarP(&project_id, "projid", "i", 0, "Numerical id of gitlab project for deleting repository")
	gitlabCmd.PersistentFlags().StringVarP(&repositories, "repos", "r", "", "List of repositories names hyphenated")

	githubCmd.Flags().StringVarP(&action, "action", "a", "", "Action with repo:['create', 'delete', 'copy', 'rename', 'list']")
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
