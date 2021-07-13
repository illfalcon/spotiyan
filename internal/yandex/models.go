package yandex

type Response struct {
	Result []Result `json:"result"`
}

type Result struct {
	Title   string   `json:"title"`
	Artists []Artist `json:"artists"`
	Albums  []Album  `json:"albums"`
}

type Artist struct {
	Name string `json:"name"`
}

type Album struct {
	Title string `json:"title"`
}

type Track struct {
	Title   string
	Artists []string
	Albums  []string
}

func (r Response) ToTrack() (Track, error) {
	if len(r.Result) == 0 {
		return Track{}, NewNoResultsFromYandex(0)
	}

	return Track{
		Title:   "",
		Artists: nil,
		Albums:  nil,
	}, nil
}
