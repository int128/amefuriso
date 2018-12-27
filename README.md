# amefuriso [![CircleCI](https://circleci.com/gh/int128/amefuriso.svg?style=shield)](https://circleci.com/gh/int128/amefuriso)

This is a Slack bot for notifying rainfall forecast using [Yahoo Japan Weather API](https://developer.yahoo.co.jp/webapi/map/openlocalplatform/v1/weather.html).

<img src="https://user-images.githubusercontent.com/321266/50439946-48c17500-0937-11e9-9207-784a9aa27058.png" width="320">

## Getting Started

TODO: This app is not yet available for public.

## Contributions

This is an open source software.
Feel free to open issues and pull requests.

### Architecture

This application is written in Go and designed for App Engine.
It consists of the following packages:

- `main` - Bootstraps the application and wires dependencies.
- `handlers` - Handles requests.
- `usecases` - Provides application use cases.
- `domain` - Provides domain of weather forecast.
- `gateways` - Provides conversion between domain models and external models.
- `infrastructure` - Invokes external APIs.

You can regenerate interface mocks as follows:

```sh
go generate -v ./...
```

### How to develop and deploy

```sh
# Install SDK
brew cask install google-cloud-sdk
gcloud components install app-engine-go

# Run
dev_appserver.py .

# Deploy
gcloud app deploy --project=amefuriso
```
