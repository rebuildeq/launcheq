package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	_ "embed"

	"github.com/xackery/launcheq/client"
)

var (
	Version     string
	PatcherUrl  string
	FileListUrl string
)

func main() {
	log.Println("initializing", Version)

	PatcherUrl = strings.TrimSuffix(PatcherUrl, "/")
	FileListUrl = strings.TrimSuffix(FileListUrl, "/")

	c, err := client.New(Version, PatcherUrl, FileListUrl)
	if err != nil {
		fmt.Println("Failed client new:", err)
		os.Exit(1)
	}
	c.Patch()
}
