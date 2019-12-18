package cmd

import (
	"github.com/spf13/cobra"

	"github.com/micnncim/octo/pkg/cmd/repo"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "octo",
		Short:   "Yet another CLI that makes easier to use with GitHub",
		Example: `octo create micnncim octo --license-template apache-2.0 --allow-squash`,
		RunE: func(c *cobra.Command, args []string) error {
			return nil
		},
	}
	cmd.AddCommand(repo.NewRepoCmd())
	return cmd
}
