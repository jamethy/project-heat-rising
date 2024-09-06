// Package util: I never know where to put this stuff
package util

import (
	"io"
	"log/slog"
	"net/url"
	"reflect"
	"runtime/debug"

	"github.com/google/go-querystring/query"
)

// AddQueryParameters adds the parameters in opt as URL query parameters to s. opt
// must be a struct whose fields may contain "url" tags.
func AddQueryParameters(s string, opt any) (string, error) {
	v := reflect.ValueOf(opt)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	qs, err := query.Values(opt)
	if err != nil {
		return s, err
	}

	u.RawQuery = qs.Encode()
	return u.String(), nil
}

func SafeClose(closer io.Closer) {
	if closer != nil {
		if err := closer.Close(); err != nil {
			slog.Warn("failed to close closer", "err", err)
			debug.PrintStack()
		}
	}
}

func Ptr[T any](t T) *T {
	return &t
}

func Unptr[T any](t *T) T {
	if t != nil {
		return *t
	}
	var tt T
	return tt
}
