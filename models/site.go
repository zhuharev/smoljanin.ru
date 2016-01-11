package models

import (
	"fmt"
	"github.com/Unknwon/com"
	"github.com/fatih/color"
	"net/url"
	"strings"
	"time"
)

type Site struct {
	Id     int64
	Domain string `xorm:"unique"`
	Title  string
	Www    bool
	Https  bool

	PreviewUrl string

	Processed bool
	AliasTo   int64
	AddedBy   int64
	Created   time.Time
	Deleted   time.Time
	Updated   time.Time
}

func (s Site) Preview() string {
	if s.PreviewUrl != "" {
		return s.PreviewUrl
	}
	return "/cat/screen/" + com.ToStr(s.Id)
}

func NewSite(link, title string) (*Site, error) {
	s := new(Site)
	if !strings.HasPrefix(link, "http://") && !strings.HasPrefix(link, "https://") {
		link = "http://" + link
	} else if strings.HasPrefix(link, "https://") {
		s.Https = true
	}
	u, e := url.Parse(link)
	if e != nil {
		return nil, e
	}
	if strings.HasPrefix(u.Host, "www.") {
		u.Host = strings.TrimPrefix(u.Host, "www.")
		s.Www = true
	}
	s.Domain = u.Host
	s.Title = title
	if has, _ := HasSite(s.Domain); has {
		return nil, fmt.Errorf("Site exists")
	}
	_, e = x.Insert(s)
	return s, e
}

func GetSite(id int64) (*Site, error) {
	s := new(Site)
	_, e := x.Id(id).Get(s)
	return s, e
}

func HasSite(host string) (has bool, e error) {
	s := new(Site)
	s.Domain = host
	return x.Cols("id").Get(s)
}

func SaveSite(s *Site) error {
	_, e := x.Id(s.Id).Update(s)
	return e
}

func SiteList(page int) (res []*Site, e error) {
	res = make([]*Site, 0)
	if page < 1 {
		page = 1
	}
	e = x.Limit(10, (page-1)*10).Find(&res)
	return
}

func SiteCount() (int64, error) {
	c, e := x.Count(new(Site))
	color.Green("Total sites %d", c)
	return c, e
}
