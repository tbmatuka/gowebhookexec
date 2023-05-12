GoWebhookExec
===============

GoWebhookExec is a service for handling webhooks written in Go. It lets you trigger a shell command with a webhook.

# Features

* Configure multiple webhooks
* Protect each webhook with a key
* Pass attributes safely to your scripts (more details in the [os.exec documentation](https://pkg.go.dev/os/exec))
* Pass files or other large payloads through the request body
* Use SSL to encrypt your HTTP connection
* Real time output to the HTTP response while the command is running
* Execution lock per handler makes sure you don't trigger more than one run at the time (but you should still have locking in your deploy script as well)

# Installation

At the moment your only options are building the app yourself or downloading the [release](https://github.com/tbmatuka/gowebhookexec/releases/latest) package which contains a built binary:
`https://github.com/tbmatuka/gowebhookexec/releases/download/v1.0/webhook-exec-v1.0-linux-amd64.tar.gz`

# Usage

## Command options

```
Usage of webhook-exec:
  --default-cmd string   Command for the default config. (default "date")
  --default-key string   Secret key for the default config. (default "")
  --host string          IP address or interface to listen on. Listens everywhere if empty. (default "")
  --port string          TCP port to listen on. (default "1234")
  --sslcert string       Path to the SSL certificate file. (default "")
  --sslkey string        Path to the SSL key file. (default "")
```

## Configuration files

Configuration files are loaded from `/etc/webhook-exec.yaml` and `$HOME/.config/webhook-exec/webhook-exec.yaml`.

```
host: eth0
port: 4567
sslcert: /etc/ssl/domain.crt
sslkey: /etc/ssl/domain.key

handler:
  default:
    key: secret
    cmd: "/usr/local/deploy.sh all"
  app1:
    key: secret1
    cmd: "/usr/local/deploy.sh app1"
```

## HTTP requests

Request URL format:
`http(s)://host:port/handler_name/key/?--option=true&argument1&argument2`

Basic request with empty key:
`$ curl "localhost:1234/default/?arg"`

Basic request with key set to `secret`:
`$ curl "localhost:1234/default/secret/?arg"`

Passing a file to the script's `stdin`:
`$ curl --data-binary "@file.txt.gz" "localhost:1234/default/secret/?arg"`

# Use case examples

## Continuous deployment (CD)

Trigger a deploy script on your server from your CI. Pass a version/tag to the script to deploy exactly the version that you want. You can optionally pass a configuration file or even your whole app as the request body.

## Restarting services

If you have an unreliable service that you can't fix and your only option is to restart it when it fails, you can use a webhook to restart it. You could set up an external monitoring service to restart it automatically or give the link to someone who doesn't have permissions to access the server but can be trusted to restart the service.

## Running tasks on a docker hosts from a container

Sometimes you need to trigger commands on your docker/VM host from inside the container. You could even have a container provide its own config or config template (for example if you have a HTTP proxy running on the host in front of your containers).

## Getting health status of services

You can have a shell script gets renders health status or versions of services you are interested in. Instead of running a cron job and pushing the results of the script, you can either see them in your browser or grab them with a tool that aggregates them.
