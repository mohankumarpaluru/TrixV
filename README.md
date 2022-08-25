# tube

`tube` is a Youtube-like (_without censorship and features you don't need!_)
Video Sharing App written in Go which also supports automatic transcoding to
MP4 H.265 AAC, multiple collections and RSS feed.

## Features

- Easy to add videos (just move a file into the folder)
- Easy to upload videos (just use the builtin uploader and automatic transcoder!)
- Builtin ffmpeg-based Transcoder that automatically converts your uploaded content to MP4 H.264 / AAC
- Builtin automatic thumbnail generator
- No database (video info pulled from file metadata)
- No JavaScript (the player UI is entirely HTML, except for the uploader which degrades!))
- Easy to customize CSS and HTML template
- Automatically generates RSS feed (at `/feed.xml`)
- Clean, simple, familiar UI

### Screenshots

![screenshot-1](screenshot-1.png?raw=true "Main Screen and Video Player")
![screenshot-2](screenshot-2.png?raw=true "Video Upload Screen")

## Getting Started

### Prebuilt Release Binaries

1. Go grab the latest binary from the
   [Releases](https://git.mills.io/prologic/tube/releases) page for your
   platform / operating system.
2. Extract the archive.
3. Run `./tube`

Open http://127.0.0.1:8000/ in your Browser!

### Published Docker Images

```#!sh
$ docker pull prologic/tube
$ docker run -p 8000:8000 -v /path/to/data:/data
```

Open http://DOCKER_MACHINE_IP:8000/ in your Browser!

Where `DOCKER_MACHINE_IP` is the IP Address of your Docker Node.

### Building From Source

```#!sh
$ git clone https://git.mills.io/prologic/tube
$ cd tube
$ make build
$ ./tube
```

Open http://127.0.0.1:8000/ in your Browser!

### A Production Deployment

A Production Deployment can be found at https://tube.mills.io/ -- This is run
and managed via a Docker Swarm cluster with a Docker-Compose / Stack very
similiar to the one you can find in the repo: [docker-compose.yml](docker-compose.yml)

Beyond this a "Production Deployment" is out-of-scope at this time for the
documentation being provided here. Please don't hesitate to file an
[Issue](https://git.mills.io/prologic/tube/issues/new) however for ask for help
or advice or contact the author directly!

## Configuration

`tube` can be configured to suit your particular needs and comes by default with
a sensbile set of defaults. There is also a default configuration at the
top-level [config.json](config.json) that you can use as a start point and
modify to suite your needs.

To Run `tube` with a provided configuration just pass the `-c /path/to/config`
option; for example:

```#!sh
$ tube -c config.json
```

Everything in the configuration is optional as the builtin defaults are used
if you do not supply anything, omit some sections or values or the configuration
is invalid. Refer to the [default config.json](config.json) for the builtin
defaults (_this files matches the builtin defaults_).

Here are some documentation on key configuration items:

### Library Options and Upload / Video Paths(s)

```#!json
{
    "library": [
        {
            "path": "videos",
            "prefix": ""
        }
    ],
}
```

Set `path` to the value of the path where you want to store videos and where
`tube` will look for new videos.

### Server Options / Upload Path and Max Upload Size

```#!json
{
    "server": {
        "host": "0.0.0.0",
        "port": 8000,
        "store_path": "tube.db",
        "upload_path": "uploads",
        "max_upload_size": 104857600
    }
}
```

- Set `host` to the interface you wish to bind to. If you want to only bind
  your local machine (_ie: localhost_) set this to `127.0.0.1`.
- Set `port` to any port you wish to bind the listening socket of the server
  to. It doesn't matter what it is as long as there it doesn't collide with
  a port already in use on your system.
- Set `store_path` to a directory where `tube` will store statistics on videos
  viewed.
- Set `upload_path` to a directory that you wish to use as a temporary working
  space for `tube` to store uploaded videos and process them. This can be a
  tmpfs file system for example for faster I/O.
- Set `max_upload_size` to the maximum number of bytes you wish to impose on
  uploaded and imported videos. Upload(s)/Import(s) that exceed this size will
  by denied by the server. This is a saftey measure so as to not DoS the
  Tube server instance. Set it to a sensible value you see fit.

### Thumbnailer / Transcoder Timeouts

```#!json
{
    "thumbnailer": {
        "timeout": 60
    },
    "transcoder": {
        "timeout": 300,
        "sizes": null
    }
}
```

- Set `timeout` to the no. of seconds to permit for thumbnail generation and
  video transcoding. This value has to be large enough for thumbnail generation
  and transcoding to take place depending on the `max_upload_size` permitted.
  These values also depend on the underlying performance of the machine Tube
  runs on. Use sensible values for your `max_upload_size` + system performance.
  This is a safety measure to ensure background processed do not run away
  and/or hog system resources. The thumbnailer and transcoder processes will
  be killed if their execution time exceeds these values.

- Set `sizes` to an map of `size` => `suffix` that you wish to support for
  transcoding videos to lower quality on Upload/Import. This is especially
  useful for serving up videos to users that have poor bandwidth or where
  data charges are high for them. The following is a valid map:

```#!json
{
    "transcoder": {
        "sizes": {
          "hd720": "720p",
          "hd480": "480p",
          "nhd":   "360p",
          "film":  "240p"
        }
    }
}
```

### Optionally Require Password for Uploading

You might be hosting a page where the public can view video, but you
don't want others to be able to upload and add content.

By specifying a password as an environment variable when running tube 
you can require this password to be provided when you access `/upload`.
The username will always be `uploader`.

```#!sh
$ auth_password=upload123 tube -c config.json
```

### Feed (RSS) Configuration

```#!json
{
    "feed": {
        "external_url": "",
        "title": "Feed Title",
        "link": "http://your-url.example/about",
        "description": "Feed Description",
        "author": {
            "name": "Author Name",
            "email": "author@somewhere.example"
        },
        "copyright": "Copyright Text"
    }
}
```

- Fill these values out as you see fit. If you are familiar with RSS
  these should be straight forward :)

### Content Proprietary Notices Configuration

{
    "copyright": {
        "content": "All Content herein Public Domain and User Contributed."
    }
}

## Contributors

Thank you to all those that have contributed to this project, battle-tested it,
used it in their own projects or products, fixed bugs, improved performance
and even fix tiny typos in documentation! Thank you and keep contributing!

You can find an [AUTHORS](AUTHORS) file where we keep a list of contributors
to the project. If you contriibute a PR please consider adding your name there.
There is also Github's own [Contributors](https://git.mills.io/prologic/tube/graphs/contributors) statistics.

[![](https://sourcerer.io/fame/prologic/prologic/tube/images/0)](https://sourcerer.io/fame/prologic/prologic/tube/links/0)
[![](https://sourcerer.io/fame/prologic/prologic/tube/images/1)](https://sourcerer.io/fame/prologic/prologic/tube/links/1)
[![](https://sourcerer.io/fame/prologic/prologic/tube/images/2)](https://sourcerer.io/fame/prologic/prologic/tube/links/2)
[![](https://sourcerer.io/fame/prologic/prologic/tube/images/3)](https://sourcerer.io/fame/prologic/prologic/tube/links/3)

## License

tube source code is available under the MIT [License](LICENSE).

Previously based off of [tube](https://github.com/wybiral/tube) by [davy wybiral
](https://github.com/wybiral). (See [LICENSE.old](LICENSE.old))
