package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/hashicorp/go-multierror"
	"github.com/isacikgoz/mmsync/internal/git"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/exp/slices"
)

var checkoutCmd = &cobra.Command{
	Use:          "checkout",
	Short:        "git-checkout - Switch branches or restore working tree files for Mattermost repositories",
	RunE:         checkoutCmdF,
	Args:         cobra.MinimumNArgs(1),
	SilenceUsage: true,
}

var pullCmd = &cobra.Command{
	Use:          "pull",
	Short:        "git-pull -  Incorporates changes from a remote repository into the current branch",
	RunE:         pullCmdF,
	SilenceUsage: true,
}

func checkoutCmdF(c *cobra.Command, args []string) error {
	cfg, err := resolveConfig(c.OutOrStdout())
	if err != nil {
		return err
	}

	fmt.Println("scanning repositories...")
	var result *multierror.Error

	for _, repo := range cfg.Repositories {
		// skip if the repository is not there
		if _, err := os.Stat(repo.Path); os.IsNotExist(err) {
			continue
		}

		cmdArgs := []string{"checkout", args[0]}
		if b, _ := c.Flags().GetBool("branch"); b {
			cmdArgs = slices.Insert(cmdArgs, 1, "-b")
		}

		err := git.RunGitCommand(cfg, repo, cmdArgs...)
		if err != nil {
			result = multierror.Append(result, err)
			continue
		}

		fmt.Fprintf(c.OutOrStdout(), "%s: successfully checked out to %s\n", repo.Name, args[0])
	}

	return result.ErrorOrNil()
}

func pullCmdF(c *cobra.Command, args []string) error {
	cfg, err := resolveConfig(c.OutOrStdout())
	if err != nil {
		return err
	}

	fmt.Println("scanning repositories...")
	var result *multierror.Error

	for _, repo := range cfg.Repositories {
		// skip if the repository is not there
		if _, err := os.Stat(repo.Path); os.IsNotExist(err) {
			continue
		}

		remote := repo.Remote
		if len(args) > 0 {
			remote = args[0]
		}

		cmdArgs := append([]string{"pull"}, args...)

		err := git.RunGitCommand(cfg, repo, cmdArgs...)
		if err != nil {
			result = multierror.Append(result, err)
		}

		fmt.Fprintf(c.OutOrStdout(), "%s: successfully pulled from %s\n", repo.Name, remote)
	}

	return result.ErrorOrNil()
}

func resolveConfig(w io.Writer) (*git.Config, error) {
	if !viper.IsSet("config") {
		fmt.Fprintln(w, "no config detected, continuing with defaults...")
		return git.DefaultConfig(), nil
	}

	fileContents, err := os.ReadFile(viper.GetString("config"))
	if err != nil {
		return nil, fmt.Errorf("there was a problem reading the config file: %w", err)
	}

	var cfg git.Config
	if err := json.Unmarshal(fileContents, &cfg); err != nil {
		return nil, fmt.Errorf("there was a problem parsing the config file: %w", err)
	}

	return &cfg, nil
}
