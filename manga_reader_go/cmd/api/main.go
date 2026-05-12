package main

import (
	"fmt"
	"os"

	"manga-reader/internal/usecases"
)

type application struct{}

func main() {
	scraperApiConfig := usecases.ScraperApiConfig{
		ApiUrl:         "http://localhost:8787",
		SearchPath:     "/search",
		GetMangaPath:   "",
		GetChapterPath: "",
	}

	mangaService := usecases.MangaService{
		ApiConfig: scraperApiConfig,
	}

	results, err := mangaService.SearchManga("naruto")
	if err != nil {
		panic(err)
	}

	fmt.Println(results)

	manga, err := mangaService.GetManga(results[0].Name)
	if err != nil {
		panic(err)
	}

	fmt.Println(manga)

	chapter, err := mangaService.GetChapter(results[0].Name, manga.Chapters[0].Name)

	fmt.Println(chapter)

	htmlString := ""

	for _, link := range chapter.PagesLinks {
		htmlString = fmt.Sprintf(`%s\n<img src="%s">`, htmlString, link)
	}

	erro := os.WriteFile("capitulo.html", []byte(htmlString), 0o644)

	if erro != nil {
		panic(err)
	}
}
