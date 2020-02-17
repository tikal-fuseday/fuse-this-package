package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"search-package/models"

	"github.com/groovili/gogtrends"
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

// SearchTrends searches google trends
func SearchTrends(search string) ([]*gogtrends.Timeline, error) {
	ctx := context.Background()
	explore, err := gogtrends.Explore(ctx,
		&gogtrends.ExploreRequest{
			ComparisonItems: []*gogtrends.ComparisonItem{
				{
					Keyword: search,
					Geo:     "US",
					Time:    "today+12-m",
				},
			},
			Category: 31, // Programming category
			Property: "",
		}, "EN")
	if err != nil {
		return nil, err
	}
	overTime, err := gogtrends.InterestOverTime(ctx, explore[0], "EN")
	if err != nil {
		return nil, err
	}
	return overTime, nil
}
