package models

// GithubRepo represents one Repo from github
type GithubRepo struct {
	Name            string  `json:"name"`
	FullName        string  `json:"full_name"`
	WatchersCount   int     `json:"watchers_count"`
	StargazersCount int     `json:"stargazers_count"`
	URL             string  `json:"url"`
	HTMLURL         string  `json:"html_url"`
	SearchScore     float64 `json:"score"`
}

// GithubRepoSearchResponse represents a Repo search response from github
type GithubRepoSearchResponse struct {
	TotalCount int          `json:"total_count"`
	Items      []GithubRepo `json:"items"`
}
