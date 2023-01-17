# Myncer ↔️ [pronounced min-suh]

Myncer (short for Music Syncer) aims to be an all-in-one solution for keeping your music up-to-date across all streaming platforms.

* Transfer playlists between music platforms
* Sync playlists (hourly, weekly, bi-weekly, monthly)
* Better sync options
    * one-way sync (ex. Spotify is the master and Tidal just mimics Spotify)
    * two-way merge (ex. Spotify & Tidal are merged into a master list and the master list is propagated to both). This one's the money.

## Motivation

I'm a huge audiophile and use Tidal for streaming to my Hi-Fi setup and Spotify to discover music. I'm a create-a-playlist for everything kinda guy and always struggle to keep them in sync.

I've gotten tired of using proprietary technologies like Soundiiz.com, tunemymusic.com, and such mostly cause I'm a cheapskate and blew all my dough on Hi-Fi speakers 😥

This aims to be an open-source alternative to these tools. If this project gains traction, happy to pay heed to the community's considerations, but for now, it's a personal project (Also, I'll have to consider how I'm gonna manage the server / serverless function costs - might end up creating an OF)


## Current Platforms supported

* Spotify
* Tidal
* Youtube (with the music videos)


## Technologies

An overview of how myncer is designed

### Authentication

Supabase - I wanted to get running off the ground fast. 

### Database

Supabase here again - I don't see any real motivation for using a SQL database here so ease of use takes priority.

### Web app

#### Framework
Next.js for this one - mostly cause I wanted to experiment with using a full stack E2E framework and this is a great

### CSS
Tailwind - similar motivation - it's a new technology and apparently improves efficiency.


## ToDo
* Framework for the web app
* Figure out how to correlate songs from one music provider to another (some global id or such?)
* Figure out how to run cron-like jobs on serverless functions (some kind of database schema I imagine that the serverless function watches periodically)