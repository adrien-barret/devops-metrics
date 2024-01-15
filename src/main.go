package main

import (
	"context"
	"log"
	"net/http"

	"github.com/adrien-barret/ghinstallation/v2"
	// "github.com/adrien-barret/ghinstallation/v2"
	"github.com/google/go-github/github" // with go modules disabled
)

const GitHubEnterpriseURL = "https://github.example.com/api/v3"

func main() {
	// Shared transport to reuse TCP connections.
	tr := http.DefaultTransport

	// Wrap the shared transport for use with the app ID 1 authenticating with installation ID 99.
	itr, err := ghinstallation.NewKeyFromFile(tr, 1, 99, "devops-metrics-dora.2024-01-15.private-key.pem")
	if err != nil {
		log.Fatal(err)
	}
	itr.BaseURL = GitHubEnterpriseURL
	ctx := context.Background()
	// Use installation transport with github.com/google/go-github
	client, clientErr := github.NewEnterpriseClient(GitHubEnterpriseURL, GitHubEnterpriseURL, &http.Client{Transport: itr})
	if clientErr != nil {
		log.Fatal(clientErr)
	}
	repos, _, err := client.Repositories.List(ctx, "", nil)
	if _, ok := err.(*github.RateLimitError); ok {
		log.Println("hit rate limit")
	}

	log.Println(repos)
}

// deb27b0b1cd7c2b1fcf64cf5b5d69f621036b623
