package shot

import (
	"io"
	"net/http"
)

func Shot(host string) (io.Reader, error) {
	shotter := "http://calm-island-9934.herokuapp.com/?url=" + host + "&width=1200&height=600&clipRect=0%2C0%2C1200%2C600"
	resp, e := http.Get(shotter)
	if e != nil {
		return nil, e
	}
	return resp.Body, nil
}
