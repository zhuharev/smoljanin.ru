package models

import (
	"github.com/Unknwon/com"
	"github.com/fatih/color"
	"github.com/sisteamnik/articler"
	"github.com/zhuharev/smoljanin.ru/modules/linkpreview"
	"net/http"
	"time"
)

type SiteFeed struct {
	Id    int64
	Title string
	Body  string `xorm:"LONGTEXT"`

	ImagePreview string
	Image        string

	Published time.Time
	Source    string `xorm:"unique"`
	SiteId    int64  `xorm:"index"`

	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}

func SiteFetchNewFeed(siteId int64) error {
	site, e := GetSite(siteId)
	if e != nil {
		return e
	}

	// exit if fetched recently
	if !site.FeedFetched.IsZero() && time.Since(site.FeedFetched) < 5*time.Minute {
		return nil
	}

	links, e := linkpreview.Articler.Index(site.Domain)
	if e != nil {
		return e
	}

	for _, link := range links {
		if fetched, e := SiteLinkFetched(link); !fetched && e == nil {
			color.Green("fetch %s", link)
			cl := http.DefaultClient
			cl.Timeout = 3 * time.Second
			bts, e := com.HttpGetBytes(cl, link, http.Header{"User-Agent": {articler.HTTPUserAgent}})
			if e != nil {
				color.Red("%s", e)
				continue
			}
			art, e := linkpreview.Articler.ParseArticle(link, bts)
			if e != nil {
				color.Red("%s", e)
				continue
			}

			art.Text = string(bts)

			sf := NewSiteFeedFromArticle(art)
			sf.SiteId = siteId
			e = SaveSiteFeed(sf)
			if e != nil {
				color.Red("%s", e)
				continue
			}
		}
	}
	site.FeedFetched = time.Now()
	return SaveSite(site)
	//x.Where("created < ? and site_id = ?", time.Now().Truncate(5*time.Minute), siteId)
}

func SaveSiteFeed(sf *SiteFeed) error {
	if sf.Id == 0 {
		_, e := x.Insert(sf)
		return e
	} else {
		_, e := x.Id(sf.Id).Update(sf)
		return e
	}
}

func SiteLinkFetched(link string) (bool, error) {
	feed := new(SiteFeed)
	return x.Where("source = ?", link).Get(feed)
}

func NewSiteFeedFromArticle(art *articler.Article) *SiteFeed {
	sf := new(SiteFeed)
	sf.Title = art.Title
	sf.Body = art.Text
	sf.Source = art.Source
	sf.Published = art.Published
	return sf
}

func GetFeed(p int) ([]*SiteFeed, error) {
	var (
		res      []*SiteFeed
		pageSize = 10
	)
	e := x.OrderBy("published desc").Limit(pageSize, pageSize*(p-1)).Find(&res)
	return res, e
}

func FeedCount() (int64, error) {
	s := new(SiteFeed)
	return x.Count(s)
}

func GetSiteFeed(siteId int64, p int) ([]*SiteFeed, error) {
	var (
		res      []*SiteFeed
		pageSize = 10
	)
	e := x.Where("site_id = ?", siteId).OrderBy("published desc").Limit(pageSize, pageSize*(p-1)).Find(&res)
	return res, e
}

func GetSiteFeedItem(id int64) (*SiteFeed, error) {
	sf := new(SiteFeed)
	_, e := x.Id(id).Get(sf)
	return sf, e
}

func SiteFeedCount(siteId int64) (int64, error) {
	s := new(SiteFeed)
	return x.Where("site_id = ?", siteId).Count(s)
}
