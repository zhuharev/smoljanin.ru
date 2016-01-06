package models

import (
	"time"
)

type Site struct {
	Id     int64
	Domain string
	Www    bool
	Https  bool

	AliasTo int64
	AddedBy int64
	Created time.Time
	Deleted time.Time
	Updated time.Time
}
