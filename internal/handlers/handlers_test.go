package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type postData struct {
	Key   string
	Value string
}

var tests = []struct {
	name               string
	url                string
	method             string
	params             []postData
	expectedStatusCode int
}{
	{name: "home", url: "/", method: "GET", params: []postData{}, expectedStatusCode: http.StatusOK},
	{name: "about", url: "/about", method: "GET", params: []postData{}, expectedStatusCode: http.StatusOK},
	{name: "generals-quarters", url: "/generals-quarters", method: "GET", params: []postData{}, expectedStatusCode: http.StatusOK},
	{name: "majors-suite", url: "/majors-suite", method: "GET", params: []postData{}, expectedStatusCode: http.StatusOK},
	{name: "search-availability", url: "/search-availability", method: "GET", params: []postData{}, expectedStatusCode: http.StatusOK},
	{name: "make-reservation", url: "/make-reservation", method: "GET", params: []postData{}, expectedStatusCode: http.StatusOK},
	{name: "post-search-availability", url: "/search-availability", method: "POST", params: []postData{
		{Key: "start", Value: "2020-01-01"},
		{Key: "end", Value: "2020-01-02"},
	}, expectedStatusCode: http.StatusOK},
	{name: "search-available-json", url: "/search-availability-json", method: "POST", params: []postData{
		{Key: "start", Value: "2020-01-01"},
		{Key: "end", Value: "2020-01-02"},
	}, expectedStatusCode: http.StatusOK},
	{name: "make-reservation-post", url: "/make-reservation", method: "POST", params: []postData{
		{Key: "first_name", Value: "John"},
		{Key: "last_name", Value: "Doe"},
		{Key: "email", Value: "mail@mail.com"},
		{Key: "phone", Value: "555-555-5555"},
	}, expectedStatusCode: http.StatusOK},
}

func TestHandlers(t *testing.T) {
	routes := getRoutes()
	ts := httptest.NewTLSServer(routes)

	defer ts.Close()

	for _, e := range tests {
		var res *http.Response
		var err error

		if e.method == "GET" {
			res, err = ts.Client().Get(ts.URL + e.url)

		} else if e.method == "POST" {
			values := url.Values{}

			for _, x := range e.params {
				values.Add(x.Key, x.Value)
			}

			res, err = ts.Client().PostForm(ts.URL+e.url, values)
		}

		if err != nil {
			t.Error(err)
		}

		if res.StatusCode != e.expectedStatusCode {
			t.Errorf("for %s expected %d but got %d", e.name, e.expectedStatusCode, res.StatusCode)
		}
	}
}
