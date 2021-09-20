GoWebhookExec
===============

GoWebhookExec is a service for handling webhooks written in Go. It lets you trigger a shell command with a webhook.

# Features

* Configure multiple webhooks
* Protect each webhook with a key
* Pass safe attributes to your scripts

# Usage examples

## Continuous deployment (CD)

Trigger a deploy script on your server from your CI. Pass a version/tag to the script to deploy exactly the version that
you want.

## Restarting services

If you have an unreliable service that you can't fix and your only option is to restart it when it fails, you can use a
webhook to restart it. You could set up an external monitoring service to restart it automatically or give the link to
someone who doesn't have permissions to access the server but can be trusted to restart the service.
