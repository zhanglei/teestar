package api

import (
	"context"
	"net/http"
	"strings"
	"time"

	"golang.org/x/oauth2"
	"github.com/weilaihui/go-gitee/gitee"
)

func NewAuthenticatedClient() *gitee.Client {
	ctx := context.Background()
	conf := &oauth2.Config{
	    ClientID:     "47e10a732882363062588a4cb26e0eea80eb4e3d32f60ce8193f7f96e467abac",
	    ClientSecret: "228f59e9306d14b95de611db8ffcb39524a9c6e499dc1c09eed1652a63781d57",
	    Scopes:       []string{"user_info", "projects", "pull_requests", "issues", "notes", "keys", "hook", "groups", "gists"},
	    Endpoint: oauth2.Endpoint{
	        AuthURL:  "https://gitee.com/oauth/auth",
	        TokenURL: "https://gitee.com/oauth/token",
	    },
	}
	token,err := conf.PasswordCredentialsToken(ctx,"noreply@daiheimao.top","1qaz2wsx")

	tp := gitee.OAuthTransport{
		Token: token,
	}

	return gitee.NewClient(tp.Client())
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
		opt := &gitee.ActivityListStarredOptions{"created", "asc", gitee.ListOptions{Page: page, PerPage: 100}}
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
		opt := &gitee.ListOptions{Page: page, PerPage: 100}
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

func IsGiteeUserOldEnough(user string) bool {
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

func HasGiteeUser(user string) bool {
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

func HasGiteeRepo(repo string) bool {
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

func IsGiteeUserStarringRepo(user string, repo string) bool {
	client := NewAuthenticatedClient()
	ctx := context.Background()

	page := 1
	got := 0

	for {
		opt := &gitee.ActivityListStarredOptions{"created", "asc", gitee.ListOptions{Page: page, PerPage: 100}}
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

func IsGiteeUserFlagged(user string) bool {
	client := &http.Client{}

	request, err := http.NewRequest("GET", "https://gitee.com/" + user, nil)
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

func IsGiteeUserActive(user string) bool {
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
