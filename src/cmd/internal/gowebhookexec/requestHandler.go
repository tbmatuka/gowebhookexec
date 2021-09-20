package gowebhookexec

import (
  "bufio"
  "log"
  "net/http"
  "os/exec"
  "strings"
  "sync"
)

type requestHandler struct {
  Name string
  Path string
  key string
  cmdName string
  lock sync.Mutex
}

func newRequestHandler(config *requestHandlerConfig) *requestHandler {
  requestHandler := new(requestHandler)

  requestHandler.Name = config.Name
  requestHandler.Path = config.Path
  requestHandler.key = config.Key
  requestHandler.cmdName = config.CmdName

  return requestHandler
}

func (requestHandler *requestHandler) handleRequest(response http.ResponseWriter, request *http.Request) {
  // check key
  pathArgs := strings.Split(strings.Trim(request.URL.Path[len(requestHandler.Path):], "/"), "/")
  requestKey := pathArgs[0]
  if requestKey != requestHandler.key {
    response.WriteHeader(http.StatusForbidden)
    response.Write([]byte("Invalid key."))

    return
  }

  // prevent multiple hooks at the same time
  requestHandler.lock.Lock()
  defer func() {
    requestHandler.lock.Unlock()
  }()

  // execute command
  cmd := exec.Command(requestHandler.cmdName)

  // set remoteAddr as env variable
  cmd.Env = append(cmd.Env, "remoteAddr=" + strings.Split(request.RemoteAddr, ":")[0])

  // set query params as arguments
  for _, queryString := range strings.Split(request.URL.RawQuery, "&") {
    cmd.Args = append(cmd.Args, queryString)
  }

  // connect body to stdin
  cmd.Stdin = request.Body

  // get output
  stdoutPipe, _ := cmd.StdoutPipe()
  cmd.Stderr = cmd.Stdout

  stdout := bufio.NewScanner(stdoutPipe)

  log.Print("start ", cmd.String())

  err := cmd.Start()
  if err != nil {
    log.Panicln(err)
  }

  flusher := response.(http.Flusher)
  for stdout.Scan() {
    response.Write([]byte(stdout.Text()+"\n"))
    flusher.Flush()
  }

  err = cmd.Wait()

  log.Print("end ", cmd.String())
}
