package models

type ApiResult struct {
	Status        string `json:"status"`
	StatusMessage string `json:"status_message"`
	Data          Data   `json:"data"`
}

type Data struct {
	MovieCount int     `json:"movie_count"`
	Limit      int     `json:"limit"`
	PageNumber int     `json:"page_number"`
	Movies     []Movie `json:"movies"`
}

type Movie struct {
	Id                      int       `json:"id"`
	Url                     string    `json:"url"`
	Title                   string    `json:"title"`
	TitleEnglish            string    `json:"title_english"`
	TitleLong               string    `json:"title_long"`
	Slug                    string    `json:"slug"`
	Year                    int       `json:"year"`
	Rating                  float64   `json:"rating"`
	Runtime                 int       `json:"runtime"`
	Genres                  []string  `json:"genres"`
	Summary                 string    `json:"summary"`
	DescriptionFull         string    `json:"description_full"`
	Synopsis                string    `json:"synopsis"`
	YtTrailerCode           string    `json:"yt_trailer_code"`
	Language                string    `json:"language"`
	MpaRating               string    `json:"mpa_rating"`
	BackgroundImage         string    `json:"background_image"`
	BackgroundImageOriginal string    `json:"background_image_original"`
	SmallCoverImage         string    `json:"small_cover_image"`
	MediumCoverImage        string    `json:"medium_cover_image"`
	LargeCoverImage         string    `json:"large_cover_image"`
	State                   string    `json:"state"`
	Torrents                []Torrent `json:"torrents"`
	DateUploaded            string    `json:"date_uploaded"`
}

type Torrent struct {
	Url          string `json:"url"`
	Hash         string `json:"hash"`
	Quality      string `json:"quality"`
	Type         string `json:"type"`
	Seeds        int    `json:"seeds"`
	Peers        int    `json:"peers"`
	Size         string `json:"size"`
	DateUploaded string `json:"date_uploaded"`
}
