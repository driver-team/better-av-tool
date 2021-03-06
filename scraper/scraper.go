package scraper

import "net/http"

type Scraper interface {
	// Remote
	FetchDoc(num string) error
	SetHTTPClient(client *http.Client)
	SetDocUrl(url string)

	// Local
	GetPlot() string
	GetTitle() string
	GetDirector() string
	GetRuntime() string
	GetTags() []string
	GetMaker() string
	GetActors() []string
	GetLabel() string
	GetNumber() string
	GetCover() string
	GetWebsite() string
	GetPremiered() string
	GetSeries() string

	NeedCut() bool
}
