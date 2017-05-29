package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/frozzare/helpwanted/search"
	"github.com/google/go-github/github"
	"github.com/olekukonko/tablewriter"
)

var (
	accessToken = flag.String("access-token", "", "Your personal GitHub access token (optional)")
	labels      = flag.String("labels", "help wanted", "Comma separated list of labels.")
	lang        = flag.String("lang", "", "Comma separated list of language repositories should be written in.")
	order       = flag.String("order", "desc", "The sort order if sort parameter is provided. One of asc or desc.")
	page        = flag.Int("page", 1, "Specify further pages.")
	perPage     = flag.Int("per-page", 30, "Number of items per page.")
	sort        = flag.String("sort", "", "The sort field. Can be comments, created, or updated. Default: results are sorted by best match.")
)

func main() {
	flag.Parse()

	search := search.NewSearch(&search.Options{
		AccessToken: *accessToken,
		Labels:      *labels,
		Lang:        *lang,
		Order:       *order,
		Page:        *page,
		PerPage:     *perPage,
		Sort:        *sort,
	})

	fmt.Printf("Searching: %s\n", search.Query())

	issues, err := search.Find()

	if err != nil {
		fmt.Printf("GitHub search failed: %s\n", err.Error())
	} else if len(issues) == 0 {
		fmt.Println("No issues found")
	} else {
		render(issues)
	}
}

func render(issues []github.Issue) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"TITLE", "CREATED", "URL"})

	for _, i := range issues {
		table.Append([]string{i.GetTitle(), formatTime(i.GetCreatedAt()), i.GetHTMLURL()})
	}

	table.Render()
}

func formatTime(t time.Time) string {
	layout := "2006-01-02 15:04:05"
	return t.Format(layout)
}
