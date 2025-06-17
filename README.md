# Myncer

Myncer aims to be an all-in-one solution for keeping your music up-to-date across all streaming platforms.

- Transfer playlists between music platforms
- Sync playlists (hourly, weekly, bi-weekly, monthly)
- Better sync options
  - one-way sync (ex. Spotify is the master and Tidal just mimics Spotify)
  - two-way merge (ex. Spotify & Tidal are merged into a master list and the master list is propagated to both). This one's the money.

## Motivation

I'm a huge audiophile and use Tidal for streaming to my Hi-Fi setup and Spotify to discover music. I'm a create-a-playlist for everything kinda guy and always struggle to keep them in sync. I've gotten tired of using proprietary technologies like Soundiiz.com, tunemymusic.com, and such mostly cause I'm a cheapskate and blew all my dough on Hi-Fi speakers 😥. This aims to be an open-source alternative to these tools. I do plan on hosting this at [myncer.hansbala.com](https://myncer.hansbala.com). It shouldn't be too costly.

## Datasources Supported

- Spotify
- Tidal
- Youtube (with the music videos)

## Development

### Prerequisites:
- [Nix](https://nixos.org/download.html) for development environment

### Initial setup.

```console
make nix
make config
```

After running `make config`, make sure to edit `server/config.dev.textpb` and `web/.env`

### Fullstack Run

```console
make up
```

To kill everything
```console
make down
```

### Focus Development

Database
```console
make db-up
```

Server
```console
make server-dev
```

Web App
```console
make web-dev
```

## Technologies

An overview of how myncer is designed

### Web app

Simple React app written in Typescript. Uses the following technologies:
- TailwindCSS
- shadcn/ui component libraries.

### Server

Written fully in Golang.

### Database

In production this uses a Cloud SQL postgres database. In local dev environment, it spins up a dev postgres instance.

