package gowebhookexec

import (
	"log"
	"net/http"
)

func Listen(config ViperConfig) {
	for handlerName, handlerViperConfig := range config.Handler {
		handlerConfig := newRequestHandlerConfig(handlerName, handlerViperConfig.Key, handlerViperConfig.CmdName)
		requestHandler := getRequestHandlerManager().newHandler(handlerConfig)
		http.HandleFunc(handlerConfig.Path, requestHandler.handleRequest)
	}

	log.Println("Listening on:", config.Host+":"+config.Port)

	if config.SslKey != "" && config.SslCert != "" {
		log.Fatal(http.ListenAndServeTLS(config.Host+":"+config.Port, config.SslCert, config.SslKey, nil))
	} else {
		log.Fatal(http.ListenAndServe(config.Host+":"+config.Port, nil))
	}
}
