package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
	"sync"
	"time"

	"github.com/google/go-github/v74/github"
	log "github.com/sirupsen/logrus"
	"go.yaml.in/yaml/v3"
)

// Excluded add-ons that we don't want to show
var blacklist = []string{}

func main() {
	repos, ddevRepos, err := listAvailableAddons()
	checkErr(err)

	err = os.Chdir("..")
	checkErr(err)

	// Remove the _addons folder to get fresh data
	err = os.RemoveAll(filepath.Join("_addons"))
	checkErr(err)

	for _, repo := range repos {
		if isRepoBlacklisted(repo, ddevRepos) {
			log.Warnf("Skipping blacklisted repo: %s", repo.GetFullName())
			continue
		}
		err := generateAddonMarkdown(repo)
		if err != nil {
			log.Errorf("Failed to create markdown file for %s: %v", repo.GetFullName(), err)
		}
		err = generateOrgIndexFile(repo)
		if err != nil {
			log.Errorf("Failed to create index file for %s: %v", repo.GetFullName(), err)
		}
	}
}

// =============================================================================
// UTILITY FUNCTIONS
// =============================================================================

func checkErr(err error) {
	if err != nil {
		log.Panic("CheckErr(): ERROR:", err)
	}
}

func isRepoBlacklisted(repo *github.Repository, ddevRepos []*github.Repository) bool {
	// Check explicit blacklist
	if slices.Contains(blacklist, repo.GetFullName()) {
		return true
	}

	// Skip non-ddev repos that have the same name as a ddev org repo
	if repo.Owner.GetLogin() != "ddev" {
		for _, ddevRepo := range ddevRepos {
			if repo.GetName() == ddevRepo.GetName() {
				return true
			}
		}
	}

	return false
}

func hasFileChanged(filePath string, newContent string) bool {
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

// =============================================================================
// GITHUB API FUNCTIONS
// =============================================================================

// listAvailableAddons retrieves all DDEV add-ons from GitHub using the 'ddev-get' topic
// Returns all repos and ddev org repos separately
func listAvailableAddons() ([]*github.Repository, []*github.Repository, error) {
	ctx, client := GetGitHubClient()
	q := "topic:ddev-get fork:true archived:false"

	opts := &github.SearchOptions{Sort: "updated", Order: "desc", ListOptions: github.ListOptions{PerPage: 200}}
	var allRepos []*github.Repository
	for {
		repos, resp, err := client.Search.Repositories(ctx, q, opts)
		if err != nil {
			msg := fmt.Sprintf("Unable to get list of available services: %v", err)
			if resp != nil {
				msg = msg + fmt.Sprintf(" rateinfo=%v", resp.Rate)
			}
			return nil, nil, fmt.Errorf("%s", msg)
		}
		allRepos = append(allRepos, repos.Repositories...)
		if resp.NextPage == 0 {
			break
		}

		// Set the next page number for the next request
		opts.Page = resp.NextPage
	}
	out := ""
	for _, r := range allRepos {
		out = out + fmt.Sprintf("%s: %s\n", r.GetFullName(), r.GetDescription())
	}
	if len(allRepos) == 0 {
		return nil, nil, fmt.Errorf("no add-ons found")
	}

	// Filter ddev org repos from allRepos
	var ddevRepos []*github.Repository
	for _, repo := range allRepos {
		if repo.Owner.GetLogin() == "ddev" {
			ddevRepos = append(ddevRepos, repo)
		}
	}

	return allRepos, ddevRepos, nil
}

// getLastCommitDate retrieves the date of the latest commit on the repository's default branch.
// This assumes the default branch is "main" or "master", which covers the vast majority of cases.
// It avoids an extra API call per repo, so it's not 100% reliable, but it's fast and good enough for most use cases.
func getLastCommitDate(repo *github.Repository) (string, error) {
	ctx, client := GetGitHubClient()
	// Get the repo to know the default branch
	commits, _, err := client.Repositories.ListCommits(ctx, repo.Owner.GetLogin(), repo.GetName(), &github.CommitsListOptions{
		ListOptions: github.ListOptions{PerPage: 1},
	})
	if err != nil {
		return "", err
	}
	if len(commits) == 0 {
		return "", fmt.Errorf("no commits found")
	}
	return commits[0].GetCommit().GetAuthor().GetDate().Format(time.DateOnly), nil
}

// fetchRepositoryReadme retrieves the README.md content from the given repository
func fetchRepositoryReadme(repo *github.Repository) (string, error) {
	ctx, client := GetGitHubClient()
	readme, _, err := client.Repositories.GetReadme(ctx, repo.Owner.GetLogin(), repo.GetName(), nil)
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

// fetchInstallConfig retrieves and parses the install.yaml content from the given repository
func fetchInstallConfig(repo *github.Repository) (*InstallDesc, error) {
	ctx, client := GetGitHubClient()

	// Fetch install.yaml from the repo (adjust the path if necessary)
	fileContent, _, _, err := client.Repositories.GetContents(ctx, repo.Owner.GetLogin(), repo.GetName(), "install.yaml", nil)
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

// getScheduledWorkflowStatus returns the status of the last scheduled workflow
func getScheduledWorkflowStatus(repo *github.Repository) string {
	ctx, client := GetGitHubClient()
	since := time.Now().Add(-24 * time.Hour)

	runs, _, err := client.Actions.ListRepositoryWorkflowRuns(ctx, repo.Owner.GetLogin(), repo.GetName(), &github.ListWorkflowRunsOptions{
		Event:  "schedule",
		Branch: repo.GetDefaultBranch(),
		ListOptions: github.ListOptions{
			PerPage: 10,
		},
	})
	if err != nil || runs == nil || len(runs.WorkflowRuns) == 0 {
		return "unknown"
	}

	for _, run := range runs.WorkflowRuns {
		if run.GetCreatedAt().After(since) {
			if run.GetConclusion() == "" {
				return "unknown"
			}
			return run.GetConclusion()
		}
	}
	return "disabled"
}

// =============================================================================
// FILE GENERATION FUNCTIONS
// =============================================================================

// generateAddonMarkdown creates a markdown file for each repository in the structure _addons/<org>/<repo>.md
func generateAddonMarkdown(repo *github.Repository) error {
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
	readmeContent, err := fetchRepositoryReadme(repo)
	if err != nil {
		log.Warnf("Could not retrieve README for repo %s: %v", repo.GetFullName(), err)
		readmeContent = "README not available."
	}

	installYaml, err := fetchInstallConfig(repo)
	if err != nil {
		return fmt.Errorf("could not retrieve install.yaml for repo %s: %v", repo.GetFullName(), err)
	}

	dependencies := "[]"
	if len(installYaml.Dependencies) > 0 {
		dependencies = fmt.Sprintf(`["%s"]`, strings.Join(installYaml.Dependencies, `", "`))
	}

	lastCommitDate, err := getLastCommitDate(repo)
	if err != nil {
		lastCommitDate = repo.GetUpdatedAt().Format(time.DateOnly)
	}

	workflowStatus := getScheduledWorkflowStatus(repo)

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
workflow_status: %s
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
		lastCommitDate,
		workflowStatus,
		repo.GetStargazersCount(),
		strings.TrimSpace(readmeContent),
	)

	if !hasFileChanged(filePath, newContent) {
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

// generateOrgIndexFile creates an index.html file for each organization in the structure _addons/<org>/index.html
func generateOrgIndexFile(repo *github.Repository) error {
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

	if !hasFileChanged(filePath, newContent) {
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

// =============================================================================
// CONTENT PROCESSING FUNCTIONS
// =============================================================================

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

// formatGitHubSpoilers ensures that GitHub spoilers are rendered correctly in Jekyll by adding markdown attributes
func formatGitHubSpoilers(content string) string {
	return strings.NewReplacer(
		`<details>`, `<details markdown="1">`,
		`<summary>`, `<summary markdown="span">`,
	).Replace(content)
}

// =============================================================================
// DATA STRUCTURES
// =============================================================================

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

// =============================================================================
// GITHUB CLIENT MANAGEMENT
// =============================================================================

var (
	// githubContext is the Go context used for GitHub API requests
	githubContext context.Context
	// githubClient is the singleton instance of Client
	githubClient *github.Client
	// githubClientOnce ensures githubClient is initialized only once
	githubClientOnce sync.Once
)

// GetGitHubClient returns a singleton GitHub client and context, initializing it if necessary.
func GetGitHubClient() (context.Context, *github.Client) {
	githubClientOnce.Do(func() {
		githubContext = context.Background()
		// Respect proxies set in the environment
		githubClient = github.NewClientWithEnvProxy()
		if githubToken := GetGitHubToken(); githubToken != "" {
			githubClient = githubClient.WithAuthToken(githubToken)
		}
	})
	return githubContext, githubClient
}

// GetGitHubToken returns the GitHub access token from the environment variable.
func GetGitHubToken() string {
	for _, token := range []string{"DDEV_ADDON_REGISTRY_TOKEN", "DDEV_GITHUB_TOKEN", "GH_TOKEN", "GITHUB_TOKEN"} {
		if githubToken := os.Getenv(token); githubToken != "" {
			return githubToken
		}
	}
	return ""
}
