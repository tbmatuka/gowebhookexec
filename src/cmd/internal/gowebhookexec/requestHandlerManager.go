package gowebhookexec

import "sync"

type requestHandlerManager struct {
  requestHandlers map[string]*requestHandler
}

var containerRequestHandlerManagerLock = &sync.Mutex{}
var containerRequestHandlerManager *requestHandlerManager

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

  requestHandlerManager.requestHandlers[config.Name] = requestHandler

  return requestHandler
}
