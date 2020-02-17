package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"search-package/client"
	"search-package/models"

	"github.com/groovili/gogtrends"
)

func normalize(val float32, max float32, min float32) float32 {
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
	repoResp, err := client.SearchRepos(queryMatch)
	if err != nil {
		panic(err)
	}
	
	trends, err := client.SearchTrends(queryMatch)
	if err != nil {
		panic(err)
	}

	
	j, err := json.Marshal(repoResp)
	if err != nil {
		panic("Could not serialize object to json")
	}
	mergeResults(&repoResp, &trends)
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
