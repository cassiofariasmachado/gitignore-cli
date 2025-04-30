package main

import (
	"flag"

	"github.com/cassiofariasmachado/gitignore-cli/internal/gitignore"
	"github.com/cassiofariasmachado/gitignore-cli/internal/utils/log"
)

var (
	debug   *bool
	path    *string
	name    *string
	baseUrl *string
)

func init() {
	name = flag.String("name", "", "Name of the gitignore file to download")
	path = flag.String("path", ".gitignore", "Path to save the gitignore file")
	debug = flag.Bool("debug", false, "Enable debug mode")
	baseUrl = flag.String("baseurl", "https://raw.githubusercontent.com/github/gitignore/main/%s.gitignore", "Base URL for downloading gitignore files")

	flag.Parse()

	if *name == "" {
		log.Fatal("Name of the .gitignore file is required")
	}
}

func main() {
	gitignore.Configure(baseUrl)
	log.Configure(debug)

	log.Debug("Debug mode: %t", *debug)
	log.Debug("GitIgnore Name: %s", *name)
	log.Debug("Path: %s", *path)

	g, err := gitignore.GetGitIgnore(*name, *path)

	if err != nil {
		log.Fatal("Error getting .gitignore: %v", err)
	}

	g.SaveFile()
	g.Close()
}
