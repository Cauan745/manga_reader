import { parse, HTMLElement } from 'node-html-parser'
import {getChaptersFunction, getImageLinksFunction, Scraper, searchMangaFunction } from '../types/scraper';
import { Manga, MangaSearch } from '../types/manga';

const headers = {
    // Mimic a standard Windows/Chrome browser
    "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
    "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8",
    "Accept-Language": "en-US,en;q=0.5"
}

const extractImagesLinkFromMangaLivre = (html: string): Array<string|undefined> => {
  const root = parse(html);

  const images = root.querySelectorAll('[id^="image-"]');

  const results = images.map(image => image.getAttribute("src"));

  return results
}

const extractChaptersFromMangaLivre = (html: string, mangaName: string): Manga => {
  const root = parse(html);
  
  const mangaTitle = root.querySelectorAll(`.post-title > h1:nth-child(1)`)[0].innerText

  const htmlChapters = root.querySelectorAll(`div[class="chapter-box"] > a`);

  const chapters = htmlChapters.map(chapter => {
    return {
      name: chapter.getAttribute("href")?.replace(`https://mangalivre.to/manga/${mangaName}/`, "").replace("/",""),
      title: chapter.innerText.trim()
    }
  });

  const results:Manga = {mangaName, mangaTitle, chapters}

  console.log(results)

  return results
}

const extractSearchResultsFromMangaLivre = (html: string): Array<MangaSearch> => {
  const root = parse(html)

  const queryResult = root.querySelectorAll('h3[class="h4"] > :first-child')

  const results = queryResult.map(link => { 
    return {
      name: link.getAttribute("href")?.replace("https://mangalivre.to/manga/", "").replace("/", ""),
      title: link.innerText
    }
  })

  return results;
}

export const searchManga: searchMangaFunction = async (name) => {
  const response = await fetch(`https://mangalivre.to/?s=${name}`, {headers})

  let html = await response.text();

  let results = extractSearchResultsFromMangaLivre(html);

  return results;
}

export const getChapters: getChaptersFunction = async (mangaName) => {

  console.log("======== Manga Page ============")

  const response = await fetch(`https://mangalivre.to/manga/${mangaName}`, {headers})

  console.log(response)

  let html = await response.text();

  console.log(html)
  
  console.log("================================")

  let results = extractChaptersFromMangaLivre(html, mangaName);

  return results;
}

export const getImageLinks: getImageLinksFunction = async (name: string, chapter: string) => {

  const response = await fetch(`https://mangalivre.to/manga/${name}/${chapter}/`, {headers})
  
  const html = await response.text();
  
  const images = extractImagesLinkFromMangaLivre(html)
  
  return images
}
