# amefurisobot [![CircleCI](https://circleci.com/gh/int128/amefurisobot.svg?style=shield)](https://circleci.com/gh/int128/amefurisobot)

A bot for notifying rainfall forecast using [Yahoo Japan Weather API](https://developer.yahoo.co.jp/webapi/map/openlocalplatform/v1/weather.html).

```sh
# Install SDK
brew cask install google-cloud-sdk
gcloud components install app-engine-go

# Run
dev_appserver.py .

# Deploy
gcloud app deploy --project=amefuriso .
```
