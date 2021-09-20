package main

import "github.com/tbmatuka/gowebhookexec/cmd/internal/gowebhookexec"

const (
  ConnectionHost = "127.0.0.1"
  ConnectionPort = "1234"
)

func main() {
  gowebhookexec.Listen(ConnectionHost, ConnectionPort)
}
