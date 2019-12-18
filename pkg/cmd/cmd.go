package cmd

import (
	"github.com/spf13/cobra"

	"github.com/micnncim/octo/pkg/cmd/repo"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "octo",
		Short:   "Yet another CLI that makes easier to use with GitHub",
		Example: `octo repo create octo --license-template apache-2.0 --allow-squash`,
	}
	cmd.AddCommand(repo.NewRepoCmd())
	return cmd
}
