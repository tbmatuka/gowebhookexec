package gowebhookexec

import (
  "log"
  "net/http"
)

func Listen(host string, port string) {
  config := newRequestHandlerConfig("test", "secret", "webhook.sh")
  requestHandler := getRequestHandlerManager().newHandler(config)

  http.HandleFunc(config.Path, requestHandler.handleRequest)

  log.Fatal(http.ListenAndServe(host + ":" + port, nil))
}
