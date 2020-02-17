package models

import "time"

// GithubRepo represents one Repo from github
type GithubRepo struct {
	Name            string          `json:"name"`
	FullName        string          `json:"full_name"`
	WatchersCount   int             `json:"watchers_count"`
	StargazersCount int             `json:"stargazers_count"`
	URL             string          `json:"url"`
	HTMLURL         string          `json:"html_url"`
	SearchScore     float64         `json:"score"`
	Description     string          `json:"description"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
	PushedAt        time.Time       `json:"pushed_at"`
	OpenIssues      int             `json:"open_issues_count"`
	Forks           int             `json:"forks_count"`
	Owner           GithubRepoOwner `json:"owner"`
}

// GithubRepoSearchResponse represents a Repo search response from github
type GithubRepoSearchResponse struct {
	TotalCount int          `json:"total_count"`
	Items      []GithubRepo `json:"items"`
}

// GithubRepoOwner represents a repo owner on github
type GithubRepoOwner struct {
	Login             string `json:"login"`
	ID                int64  `json:"id"`
	NodeID            string `json:"node_id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
}
