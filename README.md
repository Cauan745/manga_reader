# Manga Reader

Projeto em desenvolvimento para pesquisar e ler mangás.

O frontend ainda não está pronto. A interface está sendo construída com foco em TypeScript e JavaScript, assim como o scraper.

O projeto está dividido em três partes:

- front: interface web em TypeScript com TanStack Start, ainda em desenvolvimento
- manga_reader_go: API em Go com PostgreSQL
- scraper: serviço de busca e coleta de capítulos em TypeScript

## Como rodar

### API

cd manga_reader_go
make start-containers
make dev

A API fica disponível na porta 4000.

### Frontend

cd front
npm install
npm run dev

O frontend fica disponível na porta 3000.

### Scraper

cd scraper
npm install
npm run dev

O scraper é executado localmente com o Wrangler.
