package gitignore

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/cassiofariasmachado/gitignore-cli/utils/log"
)

type GitIgnore struct {
	Name    string
	Path    string
	Content io.ReadCloser
}

var (
	baseUrl *string
)

func Configure(url *string) {
	if url == nil {
		log.Fatal("Base URL cannot be nil")
	}

	baseUrl = url
}

func GetGitIgnore(name string, path string) (*GitIgnore, error) {
	url := fmt.Sprintf(*baseUrl, name)

	res, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("Error getting .gitignore from github.com (StatusCode: %d)", res.StatusCode))
	}

	return &GitIgnore{
		Name:    name,
		Path:    path,
		Content: res.Body,
	}, nil
}

func (g GitIgnore) SaveFile() {
	file, err := os.Create(g.Path)
	if err != nil {
		log.Fatal("Error creating file: %v", err)
	}
	defer file.Close()

	_, err = io.Copy(file, g.Content)
	if err != nil {
		log.Fatal("Error saving .gitignore file: %v", err)
	}

	log.Print(".gitignore file saved to path: %s\n", g.Path)
}

func (g GitIgnore) Close() {
	g.Content.Close()
}
