# Roadmap

This is mostly a collection of ideas and wishes.
The best way to make those things happen, is to actually do them.

If you want to start working on any of those features, please take
a look at the [contributing guidelines](CONTRIBUTING.md) and open
an issue to let us know and to discuss the details.

## In Progress

- refactor [uploadHandler](https://git.mills.io/prologic/tube/src/commit/c0ca374a16a75acbf380b133dde6529d7f66bb2b/app/app.go#L226)

## Prioritized

- extract hardcoded ffmpeg parameters into config.json
- Add ability to override `config.json` configuration via env variables. This would make a containerized version of tube much easier to deploy.

## Unsorted

- edit video
    - edit title / description
    - edit thumbnail (select position + generate, or upload?)
- delete video
- background transcoding / scaling
- importer framework
    - allow integration with tools like [yt-dlp](https://github.com/yt-dlp/yt-dlp)
    - allow integration with any tool that fetches a video file from a url
- library backend framework
    - Support for S3 Bucket for file storage
    - Support for recursive scanning of a library path
    - Support for read-only library sections
- player/server: on-the-fly scale-down
- uploader: check uploaded files for allowed format/codec combinations to avoid unnecessary transcoding
- on the fly transcoding (accessing large collection, w/o need for batch transcoding)
