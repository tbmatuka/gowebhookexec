package main

import "github.com/tbmatuka/gowebhookexec/cmd/internal/gowebhookexec"

func main() {
	config := gowebhookexec.LoadConfig()

	gowebhookexec.Listen(config)
}
