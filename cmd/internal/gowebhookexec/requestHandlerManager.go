package gowebhookexec

import (
	"log"
	"sync"
)

type requestHandlerManager struct {
	requestHandlers map[string]*requestHandler
}

var containerRequestHandlerManagerLock = &sync.Mutex{} //nolint:gochecknoglobals
var containerRequestHandlerManager *requestHandlerManager //nolint:gochecknoglobals

func getRequestHandlerManager() *requestHandlerManager {
	if containerRequestHandlerManager == nil {
		// only lock on initialization
		containerRequestHandlerManagerLock.Lock()

		// check again after lock
		if containerRequestHandlerManager == nil {
			containerRequestHandlerManager = new(requestHandlerManager)
			containerRequestHandlerManager.requestHandlers = make(map[string]*requestHandler)
		}

		containerRequestHandlerManagerLock.Unlock()
	}

	return containerRequestHandlerManager
}

func (requestHandlerManager *requestHandlerManager) newHandler(config requestHandlerConfig) *requestHandler {
	requestHandler := newRequestHandler(&config)

	log.Println("Starting handler:", config.Path)

	requestHandlerManager.requestHandlers[config.Name] = requestHandler

	return requestHandler
}
