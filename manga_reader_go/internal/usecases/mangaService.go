package usecases

import (
	"encoding/json"
	"fmt"
	"net/http"

	"manga-reader/internal/models"
)

type ScraperApiConfig struct {
	ApiUrl         string
	SearchPath     string
	GetMangaPath   string
	GetChapterPath string
}

type MangaService struct {
	ApiConfig ScraperApiConfig
}

func fetch[result any](url string) (result, error) {
	var res result

	resp, err := http.Get(url)
	if err != nil {
		return res, err
	}

	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (s *MangaService) SearchManga(mangaName string) ([]models.MangaSearch, error) {
	requestUrl := fmt.Sprintf("%s%s/%s", s.ApiConfig.ApiUrl, s.ApiConfig.SearchPath, mangaName)

	results, err := fetch[[]models.MangaSearch](requestUrl)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (s *MangaService) GetManga(name string) (models.Manga, error) {
	requestUrl := fmt.Sprintf("%s%s/%s", s.ApiConfig.ApiUrl, s.ApiConfig.GetMangaPath, name)

	manga, err := fetch[models.Manga](requestUrl)
	if err != nil {
		return manga, err
	}

	return manga, nil
}

func (s *MangaService) GetChapter(mangaName string, chapterName string) (models.Chapter, error) {
	requestUrl := fmt.Sprintf("%s%s/%s/%s", s.ApiConfig.ApiUrl, s.ApiConfig.GetChapterPath, mangaName, chapterName)

	chapter, err := fetch[models.Chapter](requestUrl)
	if err != nil {

		fmt.Println(err)
		return chapter, err
	}

	return chapter, nil
}
