package client

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"search-package/models"
)

const (
	authToken = "blahblah"
	apiURL    = "https://api.github.com"
)

var c *http.Client = &http.Client{}

// SearchRepos retrieves a github repo search response
func SearchRepos(search string) (models.RepoResponse, error) {
	reposURL := apiURL + "/search/repositories?sort=stars&order=desc&q=" + search + "+language:javascript"
	repoResp := models.RepoResponse{}

	req, _ := http.NewRequest("GET", reposURL, nil)
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	resp, err := c.Do(req)
	if err != nil {
		return repoResp, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	jsonError := json.Unmarshal(body, &repoResp)
	if jsonError != nil {
		return repoResp, jsonError
	}
	return repoResp, nil
}
