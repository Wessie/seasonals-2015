package main

import (
	"flag"
	"log"
	"path/filepath"

	"github.com/Wessie/seasonals"
	"github.com/Wessie/seasonals/conf"
)

func main() {
	var confName string
	flag.StringVar(&confName, "config", "", "location of configuration file to use.")
	if err := flag.Parse(); err != nil {
		log.Printf("failed to parse commandline flags: %s\n", err)
	}

	if confName == "" {
		dir := appdirs.UserConfigDir("seasonals", "seasonals", "", false)
		if dir == "" {
			log.Println("failed to determine location of configuration file.")
			return
		}

		confName = filepath.Join(dir, "seasonals.toml")
	}

	confName, err := filepath.Abs(confName)
	if err != nil {
		log.Printf("failed to locate configuration file: %s", err)
		return
	}

	if err = conf.LoadConfiguration(confName); err != nil {
		log.Printf("failed to load configuration file: %s", err)
		return
	}

	if err = seasonals.Start(); err != nil {
		log.Printf("unhandled internal error: %s", err)
		return
	}
}
