package main

import (
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/atotto/clipboard"
	"github.com/pkg/browser"
)

var (
	os, ns  string
	openURL = browser.OpenURL
	readAll = clipboard.ReadAll
	parse   = url.Parse
)

func main() {
	t := time.NewTicker(3 * time.Second)
	for {
		select {
		case <-t.C:
			do()
		}
	}
	t.Stop()
}

func do() {
	ns, _ = readAll()
	if os != ns {
		for _, n := range strings.Split(ns, "\n") {
			openIfURL(n)
		}
	}
	os = ns
}

func openIfURL(n string) {
	u, err := parse(n)
	if err == nil && isTarget(u) {
		err = openURL(u.String())
		if err != nil {
			log.Println("%w", err)
		}
	}
}

func isTarget(u *url.URL) bool {
	return u.Scheme == "https"
}
