package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"search-package/client"
	"search-package/models"
	"sort"

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
	repoResp, err := client.SearchGithubRepos(queryMatch)
	if err != nil {
		panic(err)
	}

	npmResp, err := client.SearchNpmRepos(queryMatch)
	if err != nil {
		panic(err)
	}

	trends, err := client.SearchTrends(queryMatch)
	if err != nil {
		panic(err)
	}

	// j, err := json.Marshal(repoResp)
	// if err != nil {
	// 	panic("Could not serialize object to json")
	// }
	results, err := mergeResults(&repoResp, &trends, &npmResp)
	j, err := json.Marshal(results)
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
	repos := make([]models.RepoMergeResult, 0)

	for _, repo := range gh.Items {
		repoURL := repo.HTMLURL
		if repoURL == "" {
			continue
		}
		for _, npmRepo := range npm.Objects {
			if npmRepo.Package.Links.Repository == repoURL {
				score := normalize(float64(repo.StargazersCount), 1000, 1)*.3 + normalize(npmRepo.Score.Detail.Popularity, 1000, 1)*.4 + normalize(npmRepo.Score.Detail.Quality, 1000, 1)*.3
				m := models.RepoMergeResult{repo, npmRepo.Package.Links.Npm, score}
				repos = append(repos, m)
			}
		}
	}

	sort.Slice(repos, func(i, j int) bool {
		return repos[i].OurScore > repos[j].OurScore
	})
	return repos, nil
}
