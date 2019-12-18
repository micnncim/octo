package repo

import (
	"github.com/spf13/cobra"
)

func NewRepoCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "repo",
	}
	cmd.AddCommand(newRepoCreateCmd())
	return cmd
}
