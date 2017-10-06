package api

import (
	"context"
	"net/http"
	"strings"
	"time"

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

func ListFollowingTargets(user string) []string {
	res := []string{}

	client := NewAuthenticatedClient()
	ctx := context.Background()

	page := 1
	got := 0

	for {
		opt := &github.ListOptions{Page: page, PerPage: 100}
		targets, _, err := client.Users.ListFollowing(ctx, user, opt)
		if err != nil {
			panic(err)
		}

		for _, target := range targets {
			res = append(res, *target.Login)
			// fmt.Println(*target.Login)
		}

		got = len(targets)
		if got != 100 {
			break
		} else {
			page += 1
		}
	}

	return res
}

func IsGitHubUserOldEnough(user string) bool {
	client := NewAuthenticatedClient()
	ctx := context.Background()

	userObj, _, err := client.Users.Get(ctx, user)
	if err != nil {
			panic(err)
	}

	t := userObj.CreatedAt.Time
	now := time.Now()
	days := now.Sub(t).Hours() / 24

	if days >= 30 {
		return true
	}
	return false
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

func IsGitHubUserStarringRepo(user string, repo string) bool {
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

		for _, repo2 := range starredRepos {
			if *repo2.Repository.FullName == repo {
				return true
			}
			// fmt.Println(*repo.FullName)
		}

		got = len(starredRepos)
		if got != 100 {
			break
		} else {
			page += 1
		}
	}

	return false
}

func IsGitHubUserFlagged(user string) bool {
	client := &http.Client{}

	request, err := http.NewRequest("GET", "https://github.com/" + user, nil)
	if err != nil {
		panic(err)
	}

	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	status := response.StatusCode
	return status == 404
}

func IsGitHubUserActive(user string) bool {
	client := NewAuthenticatedClient()
	ctx := context.Background()

	repos, _, err := client.Repositories.List(ctx, user, nil)
	if err != nil {
		panic(err)
	}

	if len(repos) < 3 {
		return false
	}

	for _, repo := range repos {
		if !*(repo.Fork) {
			return true
		}
	}
	return false
}
