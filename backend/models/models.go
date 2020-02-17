package models

// RepoItem represents one Repo from github
type RepoItem struct {
	Name            string `json:"name"`
	FullName        string `json:"full_name"`
	WatchersCount   int    `json:"watchers_count"`
	StargazersCount int    `json:"stargazers_count"`
	URL             string `json:"url"`
	HTMLURL         string `json:"html_url"`
}

// RepoResponse represents a Repo search response from github
type RepoResponse struct {
	TotalCount int        `json:"total_count"`
	Items      []RepoItem `json:"items"`
}
