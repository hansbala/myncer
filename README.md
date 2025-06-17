# Myncer 

<p align="center">
  <img src="https://raw.githubusercontent.com/hansbala/myncer/master/myncer-web/public/myncer.png" alt="Myncer Logo" width="200"/>
</p>

Myncer aims to be an universal music sync engine for keeping your music synced and up-to-date across all streaming platforms.

## Features

- Transfer playlists between music platforms.
- Sync playlists on a hourly, weekly, bi-weekly, or monthly schedule.
- Sync options:
  - One way syncs: (ex. Spotify is the master and Tidal just mimics Spotify)
  - Merge syncs: (ex. Spotify & Tidal are merged into a list and that list is propagated to both). This one's the money.

## Datasources Supported

- Spotify
- Tidal
- Youtube (with the music videos)

## Development

### Prerequisites:
- [Nix](https://nixos.org/download.html) for development environment

### Initial setup

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
- `TailwindCSS`
- `shadcn/ui` component library

### Server API

A minimal-ish web server written in Golang.

### Database

In production this uses a Cloud SQL postgres database.

## Motivation

I'm a huge audiophile and use Tidal for streaming to my Hi-Fi setup and Spotify to discover music. I'm a create-a-playlist for everything kinda guy and always struggle to keep them in sync. I've gotten tired of using proprietary technologies like Soundiiz.com, tunemymusic.com, and such mostly cause I'm a cheapskate and blew all my dough on Hi-Fi speakers 😥. This aims to be an open-source alternative to these tools.

