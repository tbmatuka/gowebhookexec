package gowebhookexec

type requestHandlerConfig struct {
  Name string
  Path string
  Key string
  CmdName string
}

func newRequestHandlerConfig(name string, key string, cmdName string) requestHandlerConfig {
  config := new(requestHandlerConfig)

  config.Name = name
  config.Path = "/" + name + "/"
  config.Key = key
  config.CmdName = cmdName

  return *config
}
