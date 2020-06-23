package util

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/url"
	"reflect"
	"runtime/debug"

	"github.com/google/go-querystring/query"
)

// AddQueryParameters adds the parameters in opt as URL query parameters to s. opt
// must be a struct whose fields may contain "url" tags.
func AddQueryParameters(s string, opt interface{}) (string, error) {
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
			log.Println("failed to close: ", err)
			debug.PrintStack()
		}
	}
}

func PrettyPrint(j interface{}) {
	b, err := json.MarshalIndent(j, "", "  ")
	if err != nil {
		fmt.Println("failed to marshal json ", err)
		return
	}

	fmt.Println(string(b))
}
