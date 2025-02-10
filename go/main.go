package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/google/go-github/v69/github"
	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
)

func main() {
	repos, err := listAvailableAddons(false)
	checkErr(err)

	err = os.Chdir("..")
	checkErr(err)

	for _, repo := range repos {
		err := createRepoMarkdown(repo)
		if err != nil {
			log.Errorf("Failed to create markdown file for %s: %v", repo.GetFullName(), err)
		}
		err = createIndexFile(repo)
		if err != nil {
			log.Errorf("Failed to create index file for %s: %v", repo.GetFullName(), err)
		}
	}
}

// listAvailableAddons lists the add-ons that are listed on GitHub
func listAvailableAddons(officialOnly bool) ([]*github.Repository, error) {
	client := GetGithubClient(context.Background())
	q := "topic:ddev-get fork:true"
	if officialOnly {
		q = q + " org:" + "ddev"
	}

	opts := &github.SearchOptions{Sort: "updated", Order: "desc", ListOptions: github.ListOptions{PerPage: 200}}
	var allRepos []*github.Repository
	for {

		repos, resp, err := client.Search.Repositories(context.Background(), q, opts)
		if err != nil {
			msg := fmt.Sprintf("Unable to get list of available services: %v", err)
			if resp != nil {
				msg = msg + fmt.Sprintf(" rateinfo=%v", resp.Rate)
			}
			return nil, fmt.Errorf("%s", msg)
		}
		allRepos = append(allRepos, repos.Repositories...)
		if resp.NextPage == 0 {
			break
		}

		// Set the next page number for the next request
		opts.ListOptions.Page = resp.NextPage
	}
	out := ""
	for _, r := range allRepos {
		out = out + fmt.Sprintf("%s: %s\n", r.GetFullName(), r.GetDescription())
	}
	if len(allRepos) == 0 {
		return nil, fmt.Errorf("no add-ons found")
	}
	return allRepos, nil
}

// GetGithubClient creates the required GitHub client
func GetGithubClient(ctx context.Context) github.Client {
	// Use authenticated client for higher rate limit, normally only needed for tests
	githubToken := os.Getenv("DDEV_ADDON_REGISTRY_TOKEN")
	if githubToken != "" {
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: githubToken},
		)
		tc := oauth2.NewClient(ctx, ts)
		return *github.NewClient(tc)
	}

	return *github.NewClient(nil)
}

func checkErr(err error) {
	if err != nil {
		log.Panic("CheckErr(): ERROR:", err)
	}
}

// createRepoMarkdown creates a markdown file for each repository in the structure _addons/<org>/<repo>.md
func createRepoMarkdown(repo *github.Repository) error {
	// Define the directory and filename
	org := repo.Owner.GetLogin()         // Get organization/user name
	repoName := repo.GetName()           // Get repository name
	dir := filepath.Join("_addons", org) // Create path _addons/<org>
	err := os.MkdirAll(dir, os.ModePerm) // Ensure the directory exists
	if err != nil {
		return fmt.Errorf("failed to create directory: %v", err)
	}

	// Define the markdown file path
	filePath := filepath.Join(dir, fmt.Sprintf("%s.md", repoName))

	// Get README content from the repository
	readmeContent, err := getRepoReadme(repo)
	if err != nil {
		log.Warnf("Could not retrieve README for repo %s: %v", repo.GetFullName(), err)
		readmeContent = "README not available."
	}

	// Create the front matter (YAML-like header)
	categories := "community"
	if org == "ddev" {
		categories = "official"
	}
	content := fmt.Sprintf(`---
title: %s
github_url: %s
description: "%s"
user: %s
repo: %s
categories:
  - %s
created_at: %s
updated_at: %s
stars: %d
---

%s
`, repo.GetFullName(), repo.GetHTMLURL(), strings.ReplaceAll(repo.GetDescription(), `"`, `\"`), org, repoName, categories, repo.GetCreatedAt().Format(time.DateOnly), repo.GetUpdatedAt().Format(time.DateOnly), repo.GetStargazersCount(), readmeContent)

	// Write content to the markdown file
	err = os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %v", err)
	}

	log.Infof("Markdown file created for repository: %s", repo.GetFullName())
	return nil
}

// createIndexFile creates a markdown file for each repository in the structure _addons/<org>/<repo>.md
func createIndexFile(repo *github.Repository) error {
	// Define the directory and filename
	org := repo.Owner.GetLogin()         // Get organization/user name
	dir := filepath.Join("_addons", org) // Create path _addons/<org>
	err := os.MkdirAll(dir, os.ModePerm) // Ensure the directory exists
	if err != nil {
		return fmt.Errorf("failed to create directory: %v", err)
	}

	// Define the markdown file path
	filePath := filepath.Join(dir, "index.html")

	// Generate new content for the file
	newContent := fmt.Sprintf(`---
layout: page
title: DDEV Add-on Registry / %s
group: %s
---

{%% include addon_table.html filter_by_user="%s" %%}
`, org, org, org)

	// Check if the file already exists
	if _, err := os.Stat(filePath); err == nil {
		// Read the existing file content
		existingContent, err := os.ReadFile(filePath)
		if err != nil {
			return fmt.Errorf("failed to read existing file: %v", err)
		}

		// Compare the existing content with the new content
		if string(existingContent) == newContent {
			return nil // No need to update the file
		}
	}

	// Write the new content to the file
	err = os.WriteFile(filePath, []byte(newContent), 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %v", err)
	}

	log.Infof("index file created or updated for org: %s", org)
	return nil
}

// getRepoReadme retrieves the README.md content from the given repository
func getRepoReadme(repo *github.Repository) (string, error) {
	client := GetGithubClient(context.Background())
	readme, _, err := client.Repositories.GetReadme(context.Background(), repo.Owner.GetLogin(), repo.GetName(), nil)
	if err != nil {
		return "", fmt.Errorf("could not retrieve README: %v", err)
	}

	// Decode README content (GitHub API returns it as Base64-encoded)
	content, err := readme.GetContent()
	if err != nil {
		return "", fmt.Errorf("could not decode README content: %v", err)
	}

	// Replace relative links with full links
	updatedContent := replaceRelativeLinks(content, repo)

	return updatedContent, nil
}

// replaceRelativeLinks replaces relative links with full links in the README content,
// handling both regular links and images. It ignores anchor links (e.g., "#introduction").
func replaceRelativeLinks(content string, repo *github.Repository) string {
	blobURL := fmt.Sprintf("https://github.com/%s/%s/blob/%s", repo.Owner.GetLogin(), repo.GetName(), repo.GetDefaultBranch())
	rawURL := fmt.Sprintf("https://raw.githubusercontent.com/%s/%s/%s", repo.Owner.GetLogin(), repo.GetName(), repo.GetDefaultBranch())

	// Regex to find markdown-style links, excluding anchors
	linkRegex := regexp.MustCompile(`\[(.*?)\]\(([^http#].*?)\)`)
	// Regex to find markdown-style image links: ![alt](relative-link)
	imageRegex := regexp.MustCompile(`!\[(.*?)\]\(([^http].*?)\)`)

	// Replace relative image links with raw.githubusercontent URL
	updatedContent := imageRegex.ReplaceAllStringFunc(content, func(link string) string {
		matches := imageRegex.FindStringSubmatch(link)
		if len(matches) > 2 {
			altText := matches[1]      // Alt text
			relativeLink := matches[2] // Relative link for image
			fullLink := fmt.Sprintf("%s/%s", rawURL, strings.TrimPrefix(relativeLink, "/"))
			return fmt.Sprintf("![%s](%s)", altText, fullLink)
		}
		return link
	})

	// Replace relative links (non-image) with blob URL, excluding anchors
	updatedContent = linkRegex.ReplaceAllStringFunc(updatedContent, func(link string) string {
		matches := linkRegex.FindStringSubmatch(link)
		if len(matches) > 2 {
			text := matches[1]         // Link text
			relativeLink := matches[2] // Relative link
			fullLink := fmt.Sprintf("%s/%s", blobURL, strings.TrimPrefix(relativeLink, "/"))
			return fmt.Sprintf("[%s](%s)", text, fullLink)
		}
		return link
	})

	return updatedContent
}
