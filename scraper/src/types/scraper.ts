import { getImageLinks, searchManga, getChapters } from "../scrapers/mangaLivre"
import { Manga, MangaSearch } from "./manga"

export type searchMangaFunction = (name: string) => Promise<Array<MangaSearch>>
export type getChaptersFunction = (mangaName: string) => Promise<Manga>
export type getImageLinksFunction = (name: string, chapter: string) => Promise<Array<string|undefined>>

export type Scraper = {
    searchManga: searchMangaFunction,
    getChapters: getChaptersFunction,
    getImageLinks: getImageLinksFunction
}

export const mangaLivreScraper: Scraper = {
    searchManga,
    getChapters,
    getImageLinks: getImageLinks
}