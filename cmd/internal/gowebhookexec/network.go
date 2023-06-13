package gowebhookexec

import (
	"log"
	"net/http"
	"time"
)

func Listen(config ViperConfig) {
	if len(config.Handler) == 0 {
		log.Println("No handlers configured.")
	}

	for handlerName, handlerViperConfig := range config.Handler {
		handlerConfig := newRequestHandlerConfig(handlerName, handlerViperConfig.Key, handlerViperConfig.CmdName)

		// handle case where the default handler was configured in config file only expecting the default "date" cmd
		if handlerName == "default" && handlerConfig.CmdName == "" {
			handlerConfig.CmdName = "date"
		}

		if handlerConfig.CmdName == "" {
			log.Printf("Handler '%s' has no cmd configured.\n", handlerConfig.Name)

			continue
		}

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
	} else { //nolint:revive
		log.Fatal(server.ListenAndServe())
	}
}
