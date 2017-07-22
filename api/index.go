package api

import (
	"context"

	"github.com/google/go-github/github"
)

func ListRepos(user string) []string {
	var res []string

	client := github.NewClient(nil)

	// list all organizations for user "willnorris"
	ctx := context.Background()
	repos, _, _ := client.Repositories.List(ctx, user, nil)

	for _, repo := range repos {
		res = append(res, *repo.FullName)
		// fmt.Println(*repo.FullName)
	}

	return res
}
