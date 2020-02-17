package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"search-package/client"
	"search-package/models"
	"sort"
	"sync"

	"github.com/groovili/gogtrends"
)

func normalize(val float64, max float64, min float64) float64 {
	return (val - min) / (max - min)
}

type searchController struct {
	query *regexp.Regexp
}

func setResponseHeaders(w *http.ResponseWriter) {
	(*w).Header().Set("Content-Type", "application/json")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
}

func (sc searchController) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	setResponseHeaders(&w)
	if req.Method == "OPTIONS" {
		return
	}

	path := req.URL.Path
	strMatches := sc.query.FindStringSubmatch(path)
	queryMatch := strMatches[1]
	// fmt.Printf("%q\n", strMatches)
	if len(strMatches) < 2 {
		w.Write([]byte("No query provided"))
		return
	}

	fmt.Println("Searching github for " + queryMatch)
	var repoResp models.GithubRepoSearchResponse
	wg := sync.WaitGroup{}
	wg.Add(3)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		r, err := client.SearchGithubRepos(queryMatch)
		if err != nil {
			panic(err)
		}
		repoResp = r
	}(&wg)

	var npmResp models.NpmRepoSearchResponse
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		r, err := client.SearchNpmRepos(queryMatch)
		if err != nil {
			panic(err)
		}
		npmResp = r
	}(&wg)

	var trends []*gogtrends.Timeline
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		r, err := client.SearchTrends(queryMatch)
		if err != nil {
			panic(err)
		}
		trends = r
	}(&wg)
	wg.Wait()

	results, _ := mergeResults(&repoResp, &trends, &npmResp)
	j, _ := json.Marshal(results)
	w.Write(j)
}

// RegisterSearchController registers the search controller as an http handler
func RegisterSearchController() {
	sc := searchController{
		query: regexp.MustCompile(`^/query/(.+)/?`),
	}

	http.Handle("/query", sc)
	http.Handle("/query/", sc)
}

func mergeResults(github *models.GithubRepoSearchResponse, trends *[]*gogtrends.Timeline, npm *models.NpmRepoSearchResponse) ([]models.RepoMergeResult, error) {
	gh := (*github)
	repos := make(map[string]models.RepoMergeResult)

	for _, repo := range gh.Items {
		repoURL := repo.HTMLURL
		if repoURL == "" {
			continue
		}
		for _, npmRepo := range npm.Objects {
			if npmRepo.Package.Links.Repository == repoURL {
				score := float64(repo.StargazersCount)/1000*.3 + npmRepo.Score.Detail.Popularity*.4 + npmRepo.Score.Detail.Quality*.3
				m := models.RepoMergeResult{repo, npmRepo.Package.Links.Npm, score}
				_, ok := repos[repoURL]
				if !ok {
					repos[repoURL] = m
				}
			}
		}
	}
	vals := make([]models.RepoMergeResult, 0)
	for _, v := range repos {
		vals = append(vals, v)
	}
	sort.Slice(vals, func(i, j int) bool {
		return vals[i].OurScore > vals[j].OurScore
	})
	return vals, nil
}
