package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"search-package/models"
)

const (
	authToken = "145d0586fffb2790d22f45bb76e8f629131094fc"
	apiURL    = "https://api.github.com"
)

var c *http.Client = &http.Client{}

// SearchRepos retrieves a github repo search response
func SearchRepos(search string) (models.GithubRepoSearchResponse, error) {
	fmt.Println("SearchRepos fetching results for :", search)
	reposURL := apiURL + "/search/repositories?sort=stars&order=desc&q=" + search + "+language:javascript"
	repoResp := models.GithubRepoSearchResponse{}

	req, _ := http.NewRequest("GET", reposURL, nil)
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("Authorization", "token " + authToken)
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
