# vydon

A Helm chart for Vydon that contains the api, app, and worker

**Homepage:** <https://vydon.io>

## Source Code

* <https://github.com/vydon-io/vydon>

## Requirements

| Repository | Name | Version |
|------------|------|---------|
| file://../../backend/charts/api | api | v0 |
| file://../../frontend/apps/web/charts/app | app | v0 |
| file://../../worker/charts/worker | worker | v0 |

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| api.enabled | bool | `true` | Enable or Disable Neoysnc Api |
| app.enabled | bool | `true` | Enable or Disable Neoysnc App |
| worker.enabled | bool | `true` | Enable or Disable Vydon Worker |
