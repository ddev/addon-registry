package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
	"time"

	"github.com/google/go-github/v69/github"
	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"gopkg.in/yaml.v3"
)

// Excluded add-ons that we don't want to show
var blacklist = []string{
	// We have an official add-on for this
	"sebastian-ehrling/ddev-opensearch",
	// XHGui is bundled with DDEV
	"oblakstudio/ddev-xhgui-pro",
}

func main() {
	repos, err := listAvailableAddons(false)
	checkErr(err)

	err = os.Chdir("..")
	checkErr(err)

	// Remove the _addons folder to get fresh data
	err = os.RemoveAll(filepath.Join("_addons"))
	checkErr(err)

	for _, repo := range repos {
		if inBlacklist(repo) {
			continue
		}
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

func inBlacklist(repo *github.Repository) bool {
	return slices.Contains(blacklist, repo.GetFullName())
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

	installYaml, err := getRepoInstallYaml(repo)
	if err != nil {
		return fmt.Errorf("could not retrieve install.yaml for repo %s: %v", repo.GetFullName(), err)
	}

	dependencies := "[]"
	if len(installYaml.Dependencies) > 0 {
		dependencies = fmt.Sprintf(`["%s"]`, strings.Join(installYaml.Dependencies, `", "`))
	}

	// Create the front matter (YAML-like header)
	addonType := "contrib"
	if org == "ddev" {
		addonType = "official"
	}
	newContent := fmt.Sprintf(`---
title: %s
github_url: %s
description: "%s"
user: %s
repo: %s
repo_id: %d
ddev_version_constraint: "%s"
dependencies: %s
type: %s
created_at: %s
updated_at: %s
stars: %d
---

%s
`,
		repo.GetFullName(),
		repo.GetHTMLURL(),
		strings.ReplaceAll(repo.GetDescription(), `"`, `\"`),
		org,
		repoName,
		repo.GetID(),
		installYaml.DdevVersionConstraint,
		dependencies,
		addonType,
		repo.GetCreatedAt().Format(time.DateOnly),
		repo.GetUpdatedAt().Format(time.DateOnly),
		repo.GetStargazersCount(),
		strings.TrimSpace(readmeContent),
	)

	if !isFileChanged(filePath, newContent) {
		log.Infof("No changes for repo: %s", repo.GetFullName())
		return nil
	}

	// Write content to the markdown file
	err = os.WriteFile(filePath, []byte(newContent), 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %v", err)
	}

	log.Infof("Updated repo: %s", repo.GetFullName())
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

	if !isFileChanged(filePath, newContent) {
		return nil
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

	// GitHub spoilers rendering
	updatedContent = formatGitHubSpoilers(updatedContent)

	return updatedContent, nil
}

// replaceRelativeLinks replaces relative links with full links in the README content,
// handling both regular links and images. It ignores anchor links (e.g., "#introduction").
func replaceRelativeLinks(content string, repo *github.Repository) string {
	blobURL := fmt.Sprintf("https://github.com/%s/%s/blob/%s", repo.Owner.GetLogin(), repo.GetName(), repo.GetDefaultBranch())
	rawURL := fmt.Sprintf("https://raw.githubusercontent.com/%s/%s/%s", repo.Owner.GetLogin(), repo.GetName(), repo.GetDefaultBranch())

	// Match all Markdown links
	linkRegex := regexp.MustCompile(`\[(.*?)\]\((.*?)\)`)
	imageRegex := regexp.MustCompile(`!\[(.*?)\]\((.*?)\)`)

	// Replace relative image links with raw.githubusercontent URL
	updatedContent := imageRegex.ReplaceAllStringFunc(content, func(link string) string {
		matches := imageRegex.FindStringSubmatch(link)
		if len(matches) > 2 {
			altText := matches[1]
			relativeLink := matches[2]
			// Ignore if the link starts with http or #
			if strings.HasPrefix(relativeLink, "http") || strings.HasPrefix(relativeLink, "#") {
				return link
			}
			fullLink := fmt.Sprintf("%s/%s", rawURL, strings.TrimPrefix(relativeLink, "/"))
			return fmt.Sprintf("![%s](%s)", altText, fullLink)
		}
		return link
	})

	// Replace relative links (non-image) with blob URL, excluding anchors and external links
	updatedContent = linkRegex.ReplaceAllStringFunc(updatedContent, func(link string) string {
		matches := linkRegex.FindStringSubmatch(link)
		if len(matches) > 2 {
			text := matches[1]
			relativeLink := matches[2]
			// Ignore if the link starts with http or #
			if strings.HasPrefix(relativeLink, "http") || strings.HasPrefix(relativeLink, "#") {
				return link
			}
			fullLink := fmt.Sprintf("%s/%s", blobURL, strings.TrimPrefix(relativeLink, "/"))
			return fmt.Sprintf("[%s](%s)", text, fullLink)
		}
		return link
	})

	return updatedContent
}

func formatGitHubSpoilers(content string) string {
	return strings.NewReplacer(
		`<details>`, `<details markdown="1">`,
		`<summary>`, `<summary markdown="span">`,
	).Replace(content)
}

// isFileChanged checks if a file has changed
func isFileChanged(filePath string, newContent string) bool {
	if _, err := os.Stat(filePath); err == nil {
		existingContent, err := os.ReadFile(filePath)
		if err != nil {
			return true
		}
		if string(existingContent) == newContent {
			return false
		}
	}
	return true
}

type InstallDesc struct {
	// Name must be unique in a project; it will overwrite any existing add-on with the same name.
	Name                  string            `yaml:"name"`
	ProjectFiles          []string          `yaml:"project_files"`
	GlobalFiles           []string          `yaml:"global_files,omitempty"`
	DdevVersionConstraint string            `yaml:"ddev_version_constraint,omitempty"`
	Dependencies          []string          `yaml:"dependencies,omitempty"`
	PreInstallActions     []string          `yaml:"pre_install_actions,omitempty"`
	PostInstallActions    []string          `yaml:"post_install_actions,omitempty"`
	RemovalActions        []string          `yaml:"removal_actions,omitempty"`
	YamlReadFiles         map[string]string `yaml:"yaml_read_files"`
}

// getRepoInstallYaml retrieves and parses the install.yaml content from the given repository
func getRepoInstallYaml(repo *github.Repository) (*InstallDesc, error) {
	client := GetGithubClient(context.Background())

	// Fetch install.yaml from the repo (adjust the path if necessary)
	fileContent, _, _, err := client.Repositories.GetContents(context.Background(), repo.Owner.GetLogin(), repo.GetName(), "install.yaml", nil)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve install.yaml: %v", err)
	}

	// Decode the file content (GitHub API returns Base64-encoded content)
	content, err := fileContent.GetContent()
	if err != nil {
		return nil, fmt.Errorf("could not decode install.yaml content: %v", err)
	}

	// Parse the YAML content
	var parsedYaml InstallDesc
	err = yaml.Unmarshal([]byte(content), &parsedYaml)
	if err != nil {
		return nil, fmt.Errorf("could not parse install.yaml: %v", err)
	}

	return &parsedYaml, nil
}
