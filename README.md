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

```json
<!-- entries format -->
{
  "_type": "url",
  "ie_key": "Youtube",
  "id": "iOdJOTtukaE",
  "url": "https://www.youtube.com/watch?v=iOdJOTtukaE",
  "title": "I Slept On The Floor Every Night For A Week",
  "description": "Are there amazing benefits to sleeping on the floor? That's what I wanted to find out so I decided to sleep on the floor for 7 nights. I discovered some things I was not expecting! Did I fall...",
  "duration": 1607.0,
  "channel_id": "UCTBjtACFlDmd-g6Gtva9biQ",
  "channel": "pigmie",
  "channel_url": "https://www.youtube.com/channel/UCTBjtACFlDmd-g6Gtva9biQ",
  "uploader": "pigmie",
  "uploader_id": "@FocusedLucas",
  "uploader_url": "https://www.youtube.com/@FocusedLucas",
  "thumbnails": [
    {
      "url": "https://i.ytimg.com/vi/iOdJOTtukaE/hq720.jpg?sqp=-oaymwEcCOgCEMoBSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLCR5DmB8ZeGlnASAe7gxM-76LdVfg",
      "height": 202,
      "width": 360
    },
    {
      "url": "https://i.ytimg.com/vi/iOdJOTtukaE/hq720.jpg?sqp=-oaymwEcCNAFEJQDSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLDSppnHKTlX0x-x7IisG_FUdmwWAw",
      "height": 404,
      "width": 720
    }
  ],
  "timestamp": null,
  "release_timestamp": null,
  "availability": null,
  "view_count": 1788846,
  "live_status": null,
  "channel_is_verified": true,
  "__x_forwarded_for_ip": null
}
```

## Helpful stuff

`--mark-watched`
