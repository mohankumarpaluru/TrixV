package app

import (
	"encoding/json"
	"os"
)

// Config settings for main App.
type Config struct {
	Library     []*PathConfig      `json:"library"`
	Server      *ServerConfig      `json:"server"`
	Thumbnailer *ThumbnailerConfig `json:"thumbnailer"`
	Transcoder  *TranscoderConfig  `json:"transcoder"`
	Feed        *FeedConfig        `json:"feed"`
	Copyright   *Copyright         `json:"copyright"`
}

// PathConfig settings for media library path.
type PathConfig struct {
	Path                   string `json:"path"`
	Prefix                 string `json:"prefix"`
	PreserveUploadFilename bool   `json:"preserve_upload_filename,omitempty"`
}

// ServerConfig settings for App Server.
type ServerConfig struct {
	Host                   string `json:"host"`
	Port                   int    `json:"port"`
	StorePath              string `json:"store_path"`
	UploadPath             string `json:"upload_path"`
	PreserveUploadFilename bool   `json:"preserve_upload_filename,omitempty"`
	MaxUploadSize          int64  `json:"max_upload_size"`
}

// ThumbnailerConfig settings for Transcoder
type ThumbnailerConfig struct {
	Timeout           int `json:"timeout"`
	PositionFromStart int `json:"position_from_start"`
}

// Sizes a map of ffmpeg -s option to suffix. e.g: hd720 -> #720p
type Sizes map[string]string

// TranscoderConfig settings for Transcoder
type TranscoderConfig struct {
	Timeout int   `json:"timeout"`
	Sizes   Sizes `json:"sizes"`
}

// FeedConfig settings for App Feed.
type FeedConfig struct {
	ExternalURL string `json:"external_url"`
	Title       string `json:"title"`
	Link        string `json:"link"`
	Description string `json:"description"`
	Author      struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	} `json:"author"`
	Copyright string `json:"copyright"`
}

// Copyright text for App.
type Copyright struct {
	Content string `json:"content"`
}

// DefaultConfig returns Config initialized with default values.
func DefaultConfig() *Config {
	return &Config{
		Library: []*PathConfig{
			&PathConfig{
				Path:                   "videos",
				Prefix:                 "",
				PreserveUploadFilename: false,
			},
		},
		Server: &ServerConfig{
			Host:                   "0.0.0.0",
			Port:                   8000,
			StorePath:              "tube.db",
			UploadPath:             "uploads",
			PreserveUploadFilename: false,
			MaxUploadSize:          104857600,
		},
		Thumbnailer: &ThumbnailerConfig{
			Timeout: 60,
		},
		Transcoder: &TranscoderConfig{
			Timeout: 300,
			Sizes:   Sizes(nil),
		},
		Feed: &FeedConfig{
			ExternalURL: "http://localhost:8000",
		},
		Copyright: &Copyright{
			Content: "All Content herein Public Domain and User Contributed.",
		},
	}
}

// ReadFile reads a JSON file into Config.
func (c *Config) ReadFile(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	d := json.NewDecoder(f)
	return d.Decode(c)
}
