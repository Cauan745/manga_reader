import { Hono } from 'hono'
import { cors } from 'hono/cors'
import { mangaLivreScraper, Scraper } from './types/scraper'

const app = new Hono()

// Defines the used manga service
const mangaService: Scraper = mangaLivreScraper;

app.use("/*", cors({
  origin: "*"
}))

app.get("/search/:name", async (c) => {
  const name = c.req.param('name');

  const results = await mangaService.searchManga(name);

  return c.json(results)
})

app.get('/:name', async (c) => {

  const name = c.req.param('name')

  console.log(name)

  const chapters = await mangaService.getChapters(name);

  return c.json(chapters)

})

app.get('/:name/:chapter', async (c) => {

  const name = c.req.param("name")

  const chapter = c.req.param("chapter")

  const links = await mangaService.getImageLinks(name, chapter)

  return c.json({
    name,
    links
  })
})

export default app
