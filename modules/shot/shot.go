package shot

import (
	"github.com/ungerik/go-dry"
)

func Shot(host string) ([]byte, error) {
	shotter := "http://calm-island-9934.herokuapp.com/?url=" + host + "&width=1200&height=600&clipRect=0%2C0%2C1200%2C600"
	return dry.FileGetBytes(shotter)
}
