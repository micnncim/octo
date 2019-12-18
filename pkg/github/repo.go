package github

import (
	"context"

	"github.com/google/go-github/v28/github"
)

type Repo struct {
	Name  string
	Owner string
}

type CreateRepositoryRequest struct {
	Name              string
	Org               string
	Description       string
	Homepage          string
	Private           bool
	HasIssues         bool
	HasProjects       bool
	HasWiki           bool
	IsTemplate        bool
	TeamID            int64
	AutoInit          bool
	GitignoreTemplate string
	LicenseTemplate   string
	AllowSquashMerge  bool
	AllowMergeCommit  bool
	AllowRebaseMerge  bool
}

func (c *Client) CreateRepository(ctx context.Context, req *CreateRepositoryRequest) (*Repo, error) {
	repo, _, err := c.githubClient.Repositories.Create(ctx, req.Org, &github.Repository{
		Name:              &req.Name,
		Description:       &req.Description,
		Homepage:          &req.Homepage,
		Private:           &req.Private,
		HasIssues:         &req.HasIssues,
		HasProjects:       &req.HasProjects,
		HasWiki:           &req.HasWiki,
		IsTemplate:        &req.IsTemplate,
		TeamID:            &req.TeamID,
		AutoInit:          &req.AutoInit,
		GitignoreTemplate: &req.GitignoreTemplate,
		LicenseTemplate:   &req.LicenseTemplate,
		AllowSquashMerge:  &req.AllowSquashMerge,
		AllowMergeCommit:  &req.AllowMergeCommit,
		AllowRebaseMerge:  &req.AllowRebaseMerge,
	})
	if err != nil {
		return nil, err
	}
	return &Repo{
		Name:  repo.GetName(),
		Owner: repo.Owner.GetLogin(),
	}, nil
}
