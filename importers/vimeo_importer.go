package importers

import (
	"fmt"
	"strings"

	"git.mills.io/prologic/vimeodl"
)

type VimeoImporter struct{}

func (i *VimeoImporter) GetVideoInfo(url string) (videoInfo VideoInfo, err error) {
	if strings.HasPrefix(strings.ToLower(url), "vimeo:") {
		url = strings.TrimSpace(strings.SplitN(url, ":", 2)[1])
	}

	if !strings.HasPrefix(url, "http") {
		url = "https://player.vimeo.com/video/" + url
	}

	if !strings.HasPrefix(url, "https://player.vimeo.com/video/") {
		playerURL, err := vimeodl.GetPlayerURL(url)
		if err != nil {
			err := fmt.Errorf("error finding player url: %w", err)
			return VideoInfo{}, err
		}
		url = playerURL
	}

	if !strings.HasSuffix(url, "/") {
		url += "/"
	}

	url += "config"

	config, err := vimeodl.GetVideoConfig(url)
	if err != nil {
		err := fmt.Errorf("error retrieving video config: %w", err)
		return VideoInfo{}, err
	}

	videoInfo.VideoURL = vimeodl.PickBestVideo(config)

	videoInfo.ThumbnailURL = vimeodl.PickBestThumbnail(config)

	videoInfo.ID = fmt.Sprintf("%d", config.Video.Id)
	videoInfo.Title = config.Video.Title

	return
}
