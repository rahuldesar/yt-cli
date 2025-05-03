# github.com/rahuldesar/yt-cli

| Search | Home(\*) | TRENDING | SUBSCRIPTION | CHannels | LIKED | WATCH-LATER | HISTORY | PLAYLISTS | SETTINGS |
| ------ | -------- | -------- | ------------ | -------- | ----- | ----------- | ------- | --------- | -------- |

## Basic Features

### Core

- Play video / Play audio
- `?` - options
- `h`/`l` = `<-` `->`
- `j`/`k` = `UP` `DOWN`
- `P` = toggle preview?
- Download Feature

  - Quiet mode?

  - Preview consists
    - Image
    - Duration
    - Channel Name
    - Title
    - Descriptions
    - Comments?
      - can comment? probably not

### Like to have

- Proper debugging system
- Select preferred quality (fallback to best available if preferred is not available)

  - this should override yt-dlp somehow

- Fetch all data concurrently
  - Multiple `yt-dlp` queries Possible ???
- Option to : reorder Menubar
- `Numbers` to select Menu
- Option to not refresh everytime(within certain time). so there wont be any loading time while opening app.
- `R`: refresh to refetch data of currently active menu/or all?
- `filter` keywords and channel builtin config
- Image Preview ( kinda core, but I have no idea how to do this ) - kitty supports this, but for other browsers????

## Search

(need autocomplete for `:`)
:channel/:c <channel_name> - search channel
:playlist/:p <playlist_name> - search playlist
:history/:h - search history
:recent/:r - recent

`yt-dlp "https://www.youtube.com/feed/channels" --flat-playlist -J`

## HOME

`url` : https://www.youtube.com/feed/recommended

### Some notes

```sh Data Formats
## Your Feed
## /feed/recommended works too
yt-dlp https://www.youtube.com -J --flat-playlist --extractor-args "youtubetab:approximate_date" --playlist-start 1 --playlist-end 30 --cookies-from-browser "chrome:Profile_1"

feed/history/subscription/search/liked_video/trending/explore
https://www.youtube.com/results?search_query={query}

.entries | map(keys) | unique
"__x_forwarded_for_ip"
"_type"
"availability",
"channel",
"channel_id",
"channel_is_verified",
"channel_url",
"description",
"duration",
"id",
"ie_key",
"live_status",
"release_timestamp",
"thumbnails",
"timestamp",
"title",
"uploader",
"uploader_id",
"uploader_url",
"url",
"view_count"


use diff for these :
playlist
channel
https://www.youtube.com/feed/channels - channel list



yt-dlp https://www.youtube.com/feed/recommended -J --flat-playlist --extractor-args "youtubetab:approximate_date" --playlist-start 1 --playlist-end 30 --cookies-from-browser "chrome:Profile_1"



## Helpful stuff

`--mark-watched`


```
