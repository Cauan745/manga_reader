package models

type MangaSearch struct {
	Name  string `json:"name"`
	Title string `json:"title"`
}

type Manga struct {
	Name     string          `json:"mangaName"`
	Title    string          `json:"mangaTitle"`
	Chapters []MangaChapters `json:"chapters"`

	// changes here needs changes in mangaService
}

type MangaChapters struct {
	Name  string `json:"name"`
	Title string `json:"title"`
}

type Chapter struct {
	Name string `json:"name"`
	// Links to the images
	PagesLinks []string `json:"links"`
}
