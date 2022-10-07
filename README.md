[![GitHub Workflow Status (branch)](https://img.shields.io/github/workflow/status/isacikgoz/mmsync/CI)](https://github.com/mattisacikgozermost/mmsync/actions/workflows/ci.yml?query=branch%3Amaster)

# mmsync

Utility tool for Mattermost developers.

## Install

`go install github.com/isacikgoz/mmsync`

## Usage

```man
Usage:
  mmsync [command]

Available Commands:
  checkout    git-checkout - Switch branches or restore working tree files for Mattermost repositories
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  pull        git-pull -  Incorporates changes from a remote repository into the current branch

Flags:
      --config string   path to the configuration file (default "$XDG_CONFIG_HOME/mmsync/config")
  -h, --help            help for mmsync

Use "mmsync [command] --help" for more information about a command.

```
