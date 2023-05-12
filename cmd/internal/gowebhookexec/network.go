package gowebhookexec

import (
	"log"
	"net/http"
	"time"
)

func Listen(config ViperConfig) {
	for handlerName, handlerViperConfig := range config.Handler {
		handlerConfig := newRequestHandlerConfig(handlerName, handlerViperConfig.Key, handlerViperConfig.CmdName)
		requestHandler := getRequestHandlerManager().newHandler(handlerConfig)
		http.HandleFunc(handlerConfig.Path, requestHandler.handleRequest)
	}

	server := &http.Server{
		Addr:              config.Host + ":" + config.Port,
		ReadHeaderTimeout: 3 * time.Second,
	}

	log.Println("Listening on:", server.Addr)

	if config.SslKey != "" && config.SslCert != "" {
		log.Fatal(server.ListenAndServeTLS(config.SslCert, config.SslKey))
	} else {
		log.Fatal(server.ListenAndServe())
	}
}
