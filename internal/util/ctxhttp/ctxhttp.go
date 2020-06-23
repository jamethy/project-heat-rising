package ctxhttp

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/context/ctxhttp"
	"github.com/jamethy/project-rising-heat/internal/util"
)

type GetParams struct {
	Ctx         context.Context
	HttpClient  *http.Client
	URL         string
	Query       interface{}
	Destination interface{}
}

func Get(params GetParams) (*http.Response, error) {

	if params.Query != nil {
		var err error
		params.URL, err = util.AddQueryParameters(params.URL, params.Query)
		if err != nil {
			return nil, fmt.Errorf("bad query params: %w", err)
		}
	}

	res, err := ctxhttp.Get(params.Ctx, params.HttpClient, params.URL)
	if err != nil {
		return nil, fmt.Errorf("failed to get: %w", err)
	}
	defer util.SafeClose(res.Body)

	if res.StatusCode != 200 {
		b, _ := ioutil.ReadAll(res.Body)
		return res, fmt.Errorf("non-200 response: %d - %s", res.StatusCode, string(b))
	}
	err = json.NewDecoder(res.Body).Decode(params.Destination)
	if err != nil {
		return res, fmt.Errorf("failed to parse response: %w", err)
	}

	return res, nil
}

func Do(ctx context.Context, client *http.Client, req *http.Request) (*http.Response, error) {
	return ctxhttp.Do(ctx, client, req)
}
