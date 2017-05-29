package search

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type Options struct {
	AccessToken string
	Labels      string
	Lang        string
	Order       string
	Sort        string
	Page        int
	PerPage     int
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
	res, _, err := s.client.Search.Issues(context.Background(), s.Query(), &github.SearchOptions{
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

// Query returns the search query.
func (s *Search) Query() string {
	f := []string{
		"state:open",
		"type:issue",
	}

	for _, s := range strings.Split(strings.ToLower(s.opts.Lang), ",") {
		f = append(f, fmt.Sprintf("language: \"%s\"", s))
	}

	for _, s := range strings.Split(strings.ToLower(s.opts.Labels), ",") {
		f = append(f, fmt.Sprintf("label: \"%s\"", s))
	}

	return strings.Join(f, " ")
}
