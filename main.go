package main

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type Package struct {
	FullName      string
	Description   string
	StarsCount    int
	ForksCount    int
	LastUpdatedBy string
}

func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
	  &oauth2.Token{AccessToken: "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"},
	)
	tc := oauth2.NewClient(ctx, ts)
	// get go-github client
	client := github.NewClient(tc)

	repo, _, err := client.Repositories.Get(ctx, "Golang-Coach", "Lessons")

	if err != nil {
		fmt.Printf("Problem in getting repository information %v\n", err)
		os.Exit(1)
	}

	pack := &Package{
		FullName: *repo.FullName,
		Description: *repo.Description,
		ForksCount: *repo.ForksCount,
		StarsCount: *repo.StargazersCount,
	}

	fmt.Printf("%+v\n", pack)
}
	
