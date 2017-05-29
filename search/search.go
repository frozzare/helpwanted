package search

import (
	"context"
	"net/http"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type Options struct {
	AccessToken string
	Order       string
	Sort        string
	Page        int
	PerPage     int
	Query       string
}

type Search struct {
	client *github.Client
	opts   *Options
}

// NewSearch creates new search struct.
func NewSearch(opts *Options) *Search {
	var tc *http.Client

	if len(opts.AccessToken) > 0 {
		ctx := context.Background()
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: opts.AccessToken},
		)
		tc = oauth2.NewClient(ctx, ts)
	}

	return &Search{
		client: github.NewClient(tc),
		opts:   opts,
	}
}

// Find search GitHub for issues.
func (s *Search) Find() ([]github.Issue, error) {
	res, _, err := s.client.Search.Issues(context.Background(), s.opts.Query, &github.SearchOptions{
		Order: s.opts.Order,
		Sort:  s.opts.Sort,
		ListOptions: github.ListOptions{
			Page:    s.opts.Page,
			PerPage: s.opts.PerPage,
		},
	})
	if err != nil {
		return []github.Issue{}, err
	}

	return res.Issues, nil
}
