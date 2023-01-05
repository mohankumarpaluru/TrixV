package main

import (
	"fmt"
	"os"

	"git.mills.io/prologic/tube/app"
	log "github.com/sirupsen/logrus"
	flag "github.com/spf13/pflag"
)

var (
	debug   bool
	version bool
	config  string
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] [file]\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.BoolVarP(&version, "version", "v", false, "display version information")
	flag.BoolVarP(&debug, "debug", "d", false, "enable debug logging")
	flag.StringVarP(&config, "config", "c", "config.json", "path to configuration file")
}

func main() {
	flag.Parse()

	if debug {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	if version {
		fmt.Printf("tube version %s", FullVersion())
		os.Exit(0)
	}

	cfg := app.DefaultConfig()
	log.Infof("Reading configuration from %s", config)
	err := cfg.ReadFile(config)
	if err != nil {
		if flag.Lookup("config").Changed {
			log.Fatal(err)
		}
		log.Infof("Reading %s failed. Starting with builtin defaults.", config)
	}
	// I'd like to add something like this here: log.Debug("Active config: %s", cfg.toJson())
	a, err := app.NewApp(cfg)
	if err != nil {
		log.Fatal(err)
	}
	addr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	log.Printf("Local server: http://%s", addr)
	err = a.Run()
	if err != nil {
		log.Fatal(err)
	}
}
