package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"search-package/client"
	"search-package/models"
	"sync"

	"github.com/groovili/gogtrends"
)

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
	wg := sync.WaitGroup{}
	wg.Add(2)
	path := req.URL.Path
	strMatches := sc.query.FindStringSubmatch(path)
	// fmt.Printf("%q\n", strMatches)
	if len(strMatches) < 2 {
		w.Write([]byte("No query provided"))
		return
	}
	queryMatch := strMatches[1]
	fmt.Println("Searching github for " + queryMatch)
	// get github
	repoResp := &models.GithubRepoSearchResponse{}
	go client.SearchRepos(queryMatch, &wg, repoResp)
	// get trends
	var trends []*gogtrends.Timeline
	go client.SearchTrends(queryMatch, &wg, &trends)

	wg.Wait()

	j, err := json.Marshal(repoResp)
	if err != nil {
		panic("Could not serialize object to json")
	}

	mergeResults(repoResp, &trends)
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

func mergeResults(github *models.GithubRepoSearchResponse, trends *[]*gogtrends.Timeline) (interface{}, error) {
	return nil, nil
}
