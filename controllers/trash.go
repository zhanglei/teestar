package controllers

import (
	"errors"
	"strings"

	"github.com/hsluoyz/gitstar/api"
)

type UserTargetStatus struct {
	StarringRepos []string
	StarredRepos  []string
}

func getAllUserAndOrganRepos(user string) []string {
	var repos []string

	tokens := strings.Split(user, ",")
	for _, token := range tokens {
		repos = append(repos, api.ListRepos(token)...)
	}

	return repos
}

func getRealUser(user string) string {
	tokens := strings.Split(user, ",")
	if len(tokens) == 0 {
		panic(errors.New("invalid user"))
	}

	return tokens[0]
}
