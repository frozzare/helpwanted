package main

import (
	"flag"
	"log"
	"os"

	"github.com/frozzare/helpwanted/search"
	"github.com/google/go-github/github"
	"github.com/olekukonko/tablewriter"
)

var (
	accessToken = flag.String("access-token", "", "Your personal GitHub access token (optional)")
	sort        = flag.String("sort", "", "The sort field. Can be comments, created, or updated. Default: results are sorted by best match.")
	order       = flag.String("order", "desc", "The sort order if sort parameter is provided. One of asc or desc. Default: desc")
	page        = flag.Int("page", 1, "Specify further pages. Default: 1.")
	perPage     = flag.Int("per-page", 30, "Number of items per page. Default: 30.")
	query       = flag.String("query", "label:\"help wanted\"", "Modify GitHub search query. Default: label: \"help wanted\"")
)

func main() {
	flag.Parse()

	issues, err := search.NewSearch(&search.Options{
		AccessToken: *accessToken,
		Order:       *order,
		Page:        *page,
		PerPage:     *perPage,
		Sort:        *sort,
		Query:       *query,
	}).Find()

	if err != nil {
		log.Printf("GitHub search failed: %s\n", err.Error())
	} else {
		render(issues)
	}
}

func render(issues []github.Issue) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Title", "URL"})

	for _, i := range issues {
		table.Append([]string{i.GetTitle(), i.GetHTMLURL()})
	}

	table.Render()
}
