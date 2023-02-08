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
	//go:embed version.txt
	version string
	//go:embed url.txt
	url string
)

func main() {
	log.Println("initializing", version)

	url = strings.TrimSuffix(url, "/")

	c, err := client.New(version, url)
	if err != nil {
		fmt.Println("Failed client new:", err)
		os.Exit(1)
	}
	c.Patch()
}
