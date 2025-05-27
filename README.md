# Myncer ↔️ [pronounced min-suh]

Myncer (short for Music Syncer) aims to be an all-in-one solution for keeping your music up-to-date across all streaming platforms.

- Transfer playlists between music platforms
- Sync playlists (hourly, weekly, bi-weekly, monthly)
- Better sync options
  - one-way sync (ex. Spotify is the master and Tidal just mimics Spotify)
  - two-way merge (ex. Spotify & Tidal are merged into a master list and the master list is propagated to both). This one's the money.

## Motivation

I'm a huge audiophile and use Tidal for streaming to my Hi-Fi setup and Spotify to discover music. I'm a create-a-playlist for everything kinda guy and always struggle to keep them in sync.

I've gotten tired of using proprietary technologies like Soundiiz.com, tunemymusic.com, and such mostly cause I'm a cheapskate and blew all my dough on Hi-Fi speakers 😥

This aims to be an open-source alternative to these tools. I do plan on hosting this at [myncer.hansbala.com](https://myncer.hansbala.com). It shouldn't be too costly.

## Current Platforms supported

- Spotify
- Tidal
- Youtube (with the music videos)

## Technologies

An overview of how myncer is designed

### Web app

Simple React app written in Typescript. Uses the following technologies:
- TailwindCSS
- shadcn/ui component libraries.

### Backend

Written fully in Golang.

### Database

In production this uses a Cloud SQL postgres database. In local dev environment, it spins up a dev postgres instance.

## Development

### Initial setup.

```console
git clone git@github.com:hansbala/myncer.git
cd myncer/
make nix
```

### Full stack development

After setting up the nix environment, you would also need to setup the env variables for each service. Then,

```console
# Spins up a test database.
make database
# Starts the backend.
make backend-run
# Starts the frontend web app.
make frontend
```
