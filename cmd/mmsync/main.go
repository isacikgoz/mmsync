package main

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var RootCmd = &cobra.Command{
	Use:   "mmsync",
	Short: "Utility tool for Mattermost developers",
}

func init() {
	RootCmd.PersistentFlags().String("config", filepath.Join("$XDG_CONFIG_HOME", "mmsync", "config"), "path to the configuration file")
	_ = viper.BindPFlag("config", RootCmd.PersistentFlags().Lookup("config"))

	checkoutCmd.Flags().BoolP("branch", "b", false, "causes a new branch to be created if it doesn’t exist")
	RootCmd.AddCommand(checkoutCmd)
	RootCmd.AddCommand(pullCmd)
}

func main() {
	RootCmd.SetArgs(os.Args[1:])

	if err := RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
