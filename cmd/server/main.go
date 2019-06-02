package main

import (
	"internal/server"
)

func main() {
	server.MarkdownServer(80, "posts/", "templates/*", "static/")
}
