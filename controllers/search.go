package controllers

import (
	"encoding/json"
	"net/http"
	"regexp"
	"search-package/client"
)

type searchController struct {
	query *regexp.Regexp
}

func (sc searchController) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path

	repoResp, err := client.SearchRepos(path[7:])
	if err != nil {
		panic(err)
	}

	j, err := json.Marshal(repoResp)
	if err != nil {
		panic("Could not serialize object to json")
	}
	w.Header().Set("Content-Type", "application/json")
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
