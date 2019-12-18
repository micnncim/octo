package github

import (
	"context"
	"errors"

	"github.com/google/go-github/v28/github"
	"golang.org/x/oauth2"
)

type Client struct {
	githubClient *github.Client
	token        string
}

func NewClient(token string) (*Client, error) {
	if token == "" {
		return nil, errors.New("missing GITHUB_TOKEN")
	}
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	return &Client{
		githubClient: github.NewClient(tc),
	}, nil
}
