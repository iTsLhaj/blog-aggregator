package main

import (
	"context"
	"encoding/xml"
	"html"
	"io/ioutil"
	"net/http"
	"strings"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	var body []byte
	var feed RSSFeed
	var req *http.Request
	var res *http.Response
	var err error

	req, err = http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "gator")

	res, err = http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	body = []byte(html.UnescapeString(string(body)))
	for _, e := range []string{"&rsquo;", "&ndash;", "&ldquo;", "&rdquo;", "&hellip;"} {
		body = []byte(strings.Replace(string(body), e, "", -1))
	}
	err = xml.Unmarshal(body, &feed)
	if err != nil {
		return nil, err
	}

	return &feed, nil
}
