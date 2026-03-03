export type Manga = {
    mangaName: string | undefined,
    mangaTitle: string
    chapters: {
      name: string | undefined, title:string
    }[]
}

export type MangaSearch = {
	name: string | undefined,
	title: string
}