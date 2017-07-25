package api

import (
	"context"
	"strings"

	"github.com/google/go-github/github"
)

func NewAuthenticatedClient() *github.Client {
	tp := github.BasicAuthTransport{
		Username: "mytesttest@test.com",
		Password: "1qaz2wsx",
	}

	return github.NewClient(tp.Client())
}

func ListRepos(user string) []string {
	var res []string

	client := NewAuthenticatedClient()
	ctx := context.Background()

	repos, _, err := client.Repositories.List(ctx, user, nil)
	if err != nil {
		panic(err)
	}

	for _, repo := range repos {
		res = append(res, *repo.FullName)
		// fmt.Println(*repo.FullName)
	}

	return res
}

func ListStarringRepos(user string) []string {
	res := []string{}

	client := NewAuthenticatedClient()
	ctx := context.Background()

	page := 1
	got := 0

	for {
		opt := &github.ActivityListStarredOptions{"created", "asc", github.ListOptions{Page: page, PerPage: 100}}
		starredRepos, _, err := client.Activity.ListStarred(ctx, user, opt)
		if err != nil {
			panic(err)
		}

		for _, repo := range starredRepos {
			res = append(res, *repo.Repository.FullName)
			// fmt.Println(*repo.FullName)
		}

		got = len(starredRepos)
		if got != 100 {
			break
		} else {
			page += 1
		}
	}

	return res
}

func HasGitHubUser(user string) bool {
	client := NewAuthenticatedClient()
	ctx := context.Background()

	_, resp, err := client.Users.Get(ctx, user)
	if err != nil {
		if resp.Response.StatusCode == 404 {
			return false
		} else {
			panic(err)
		}
	}

	return true
}

func HasGitHubRepo(repo string) bool {
	client := NewAuthenticatedClient()
	ctx := context.Background()

	tokens := strings.Split(repo, "/")
	if len(tokens) != 2 {
		return false
	}

	_, resp, err := client.Repositories.Get(ctx, tokens[0], tokens[1])
	if err != nil {
		if resp.Response.StatusCode == 404 {
			return false
		} else {
			panic(err)
		}
	}

	return true
}
