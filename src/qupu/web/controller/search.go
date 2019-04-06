package controller

import (
	"context"
	"encoding/json"
	"gopkg.in/olivere/elastic.v5"
	model2 "model"
	"net/http"
	"qupu/web/model"
	"qupu/web/view"
	"reflect"
	"strconv"
	"strings"
)

type SearchResultHandler struct {
	view   view.SearchResultView
	client *elastic.Client
}

func CreateSearchResultHandle(filename string) SearchResultHandler {
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	return SearchResultHandler{
		client: client,
		view:   view.CreateSearchResultView(filename),
	}
}

func (s SearchResultHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	q := strings.TrimSpace(req.FormValue("q"))
	from, err := strconv.Atoi(req.FormValue("from"))
	if err != nil {
		from = 0
	}

	data, err := s.getSearchResult(q, from)

	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
	}
	dataStr, err := json.Marshal(data)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
	}

	_, err = res.Write([]byte(dataStr))

	//err = s.view.Render(res, data)

	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
	}
}

func (s SearchResultHandler) getSearchResult(q string, from int) (model.SearchResult, error) {
	var data model.SearchResult
	resp, err := s.client.
		Search("qupu").
		Query(elastic.NewQueryStringQuery(q)).
		From(from).
		Do(context.Background())

	if err != nil {
		return data, err
	}

	data.Start = from
	data.Hits = int(resp.TotalHits())
	data.Songs = resp.Each(reflect.TypeOf(model2.Spectrum{}))
	return data, nil

}
