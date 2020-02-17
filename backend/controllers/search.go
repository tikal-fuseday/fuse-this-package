package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"search-package/client"

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

type trendsController struct {
	query *regexp.Regexp
}

func (sc searchController) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	setResponseHeaders(&w)
	if req.Method == "OPTIONS" {
		return
	}

	path := req.URL.Path
	strMatches := sc.query.FindStringSubmatch(path)
	// fmt.Printf("%q\n", strMatches)
	if len(strMatches) < 2 {
		w.Write([]byte("No query provided"))
		return
	}

	repoResp, err := client.SearchNpmRepos(strMatches[1])

	if err != nil {
		panic(err)
	}

	j, err := json.Marshal(repoResp)
	if err != nil {
		panic("Could not serialize object to json")
	}
	w.Write(j)
}

// trends
func (tc trendsController) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	path := req.URL.Path
	strMatches := tc.query.FindStringSubmatch(path)

	if len(strMatches) < 2 {
		w.Write([]byte("No trends query provided"))
		return
	}

	fmt.Println("Searching trends for " + strMatches[1])

	ctx := context.Background()
	explore, err := gogtrends.Explore(ctx,
		&gogtrends.ExploreRequest{
			ComparisonItems: []*gogtrends.ComparisonItem{
				{
					Keyword: strMatches[1],
					Geo:     "US",
					Time:    "today+12-m",
				},
			},
			Category: 31, // Programming category
			Property: "",
		}, "EN")

	overTime, err := gogtrends.InterestOverTime(ctx, explore[0], "EN")
	if err != nil {
		panic(err)
	}

	j, err := json.Marshal(overTime)
	if err != nil {
		panic("Could not serialize")
	}
	w.Write(j)
}

// RegisterSearchController registers the search controller as an http handler
func RegisterSearchController() {
	sc := searchController{
		query: regexp.MustCompile(`^/query/(.+)/?`),
	}

	http.Handle("/query", sc)
	http.Handle("/query/", sc)

	// Google Trends
	tc := trendsController{
		query: regexp.MustCompile(`^/trends/(.+)/?`),
	}
	http.Handle("/trends", tc)
	http.Handle("/trends/", tc)

}
