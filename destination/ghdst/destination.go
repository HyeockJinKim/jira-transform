package ghdst

import (
	"context"
	"fmt"

	"github.com/HyeockJinKim/jira-transform/data"
	"github.com/google/go-github/v67/github"
)

type Args struct {
	Token string
	Owner string
}

type Destination struct {
	cli   *github.Client
	owner string
}

func NewDestination(args Args) *Destination {
	cli := github.NewClient(nil).WithAuthToken(args.Token)
	return &Destination{
		cli:   cli,
		owner: args.Owner,
	}
}

func (d *Destination) Mirror(ctx context.Context, issue data.Issue) error {
	_, _, err := d.cli.Issues.Create(ctx, d.owner, issue.Repo, &github.IssueRequest{
		Title:  &issue.Title,
		Body:   &issue.Description,
		Labels: &issue.Labels,
	})
	if err != nil {
		return fmt.Errorf("failed to create issue: %w", err)
	}
	return nil
}
