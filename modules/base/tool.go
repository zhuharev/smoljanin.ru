package base

import (
	"github.com/microcosm-cc/bluemonday"
	"regexp"
)

var Sanitizer = bluemonday.UGCPolicy().AllowAttrs("class").Matching(regexp.MustCompile(`[\p{L}\p{N}\s\-_',:\[\]!\./\\\(\)&]*`)).OnElements("code")

func ShortSha(sha1 string) string {
	if len(sha1) == 40 {
		return sha1[:10]
	}
	return sha1
}
