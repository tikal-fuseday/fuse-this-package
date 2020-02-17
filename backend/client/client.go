package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"search-package/models"
)

const (
	authToken    = "145d0586fffb2790d22f45bb76e8f629131094fc"
	githubApiURL = "https://api.github.com"
	npmApiURL    = "http://registry.npmjs.com"
)

var c *http.Client = &http.Client{}

// SearchGithubRepos retrieves a github repo search response
func SearchGithubRepos(search string) (models.GithubRepoSearchResponse, error) {
	fmt.Println("Fetching github results for :", search)
	reposURL := githubApiURL + "/search/repositories?sort=stars&order=desc&q=" + search + "+language:javascript"
	repoResp := models.GithubRepoSearchResponse{}

	req, _ := http.NewRequest("GET", reposURL, nil)
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("Authorization", "token "+authToken)
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

// SearchNpmRepos retrieves a npm package search response
func SearchNpmRepos(search string) (models.NpmRepoSearchResponse, error) {
	//	filters := models.NPMFilter{Quality: 0.0, Maintenance: 0.0, Popularity: 1.0}
	fmt.Println("Fetching npm results for :", search)
	reposURL := npmApiURL + "/-/v1/search?text=" + search + ""

	repoResp := models.NpmRepoSearchResponse{}

	req, _ := http.NewRequest("GET", reposURL, nil)

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
	fmt.Println(repoResp)
	return repoResp, nil
}
