App for manupilating repositories on different git providers:
        GITHUB/GITLAB

Usage:
  API_GIT [flags]
  API_GIT [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  github      Usage github command
  gitlab      Usage gitlab command
  help        Help about any command

Flags:
  -h, --help   help for API_GIT

Use "API_GIT [command] --help" for more information about a command.

GITLAB:

Usage:
  API_GIT gitlab [flags]

Flags:
  -a, --action string     Action with repo:['create', 'delete', 'copy', 'rename', 'list']
  -d, --descript string   Description for created repository
  -h, --help              help for gitlab
  -n, --namespace int     Numerical id of gitlab namespace for creared repository
  -p, --path string       URL path for created repository
  -i, --projid int        Numerical id of gitlab project for deleting repository
  -r, --repos string      List of repositories names hyphenated

required flag(s) "action" not set
