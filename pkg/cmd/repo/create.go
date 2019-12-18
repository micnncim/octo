package repo

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/micnncim/octo/pkg/github"
)

type createRepoOpt struct {
	name              string
	org               string
	description       string
	homepage          string
	private           bool
	hasIssues         bool
	hasProjects       bool
	hasWiki           bool
	isTemplate        bool
	teamID            int64
	autoInit          bool
	gitignoreTemplate string
	licenseTemplate   string
	allowSquashMerge  bool
	allowMergeCommit  bool
	allowRebaseMerge  bool
}

func newRepoCreateCmd() *cobra.Command {
	opt := &createRepoOpt{}
	cmd := &cobra.Command{
		Use:     "create",
		Aliases: []string{"c"},
		Example: `octo create <org>/<repo> --license-template apache-2.0 --allow-squash
octo create <repo> --auto-init --gitignore-template go --allow-squash
`,
		RunE: func(_ *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("invalid arguments")
			}
			org, repo := parseRepo(args[0])
			req := &github.CreateRepositoryRequest{
				Name:              repo,
				Org:               org,
				Description:       opt.description,
				Homepage:          opt.homepage,
				Private:           opt.private,
				HasIssues:         opt.hasIssues,
				HasProjects:       opt.hasProjects,
				HasWiki:           opt.hasWiki,
				IsTemplate:        opt.isTemplate,
				TeamID:            opt.teamID,
				AutoInit:          opt.autoInit,
				GitignoreTemplate: opt.gitignoreTemplate,
				LicenseTemplate:   opt.licenseTemplate,
				AllowSquashMerge:  opt.allowSquashMerge,
				AllowMergeCommit:  opt.allowMergeCommit,
				AllowRebaseMerge:  opt.allowRebaseMerge,
			}
			c, err := github.NewClient(os.Getenv("GITHUB_TOKEN"))
			if err != nil {
				return err
			}
			ctx := context.Background()
			r, err := c.CreateRepository(ctx, req)
			if err != nil {
				return err
			}
			fmt.Printf("https://github.com/%s/%s\n", r.Owner, r.Name)
			return nil
		},
	}

	cmd.Flags().StringVarP(&opt.description, "description", "d", "", "The description of repository")
	cmd.Flags().BoolVarP(&opt.private, "private", "p", false, "Whether the repository is private or not")
	cmd.Flags().BoolVar(&opt.autoInit, "auto-init", false, "Whether automatically initializing repository or not")
	cmd.Flags().StringVar(&opt.gitignoreTemplate, "gitignore-template", "", "The gitignore template")
	cmd.Flags().StringVar(&opt.licenseTemplate, "license-template", "", "The license template of repository")
	cmd.Flags().BoolVar(&opt.allowRebaseMerge, "allow-rebase-merge", false, "Whether allowing rebase merge or not")
	cmd.Flags().BoolVar(&opt.allowSquashMerge, "allow-squash-merge", false, "Whether allowing squash merge or not")
	cmd.Flags().BoolVar(&opt.allowMergeCommit, "allow-merge-commit", false, "Whether allowing merge commit or not")

	return cmd
}

func parseRepo(s string) (org, repo string) {
	ss := strings.Split(s, "/")
	switch len(ss) {
	case 1:
		repo = ss[0]
	case 2:
		org, repo = ss[0], ss[1]
	}
	return
}
