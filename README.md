# amefuriso [![CircleCI](https://circleci.com/gh/int128/amefuriso.svg?style=shield)](https://circleci.com/gh/int128/amefuriso)

A command for showing rainfall forecast using [YOLP weather API](https://developer.yahoo.co.jp/webapi/map/openlocalplatform/v1/weather.html).

For example,

```
% ./amefuriso -client-id $YOLP_CLIENT_ID -lat 43.1907584 -lon 140.9944283
Weather at (140.994430, 43.190758)
| 17:45 |  0.00 mm/h |
| 17:50 |  0.00 mm/h |
| 17:55 |  0.00 mm/h |
| 18:00 |  0.00 mm/h |
| 18:05 |  0.00 mm/h |
| 18:10 |  0.75 mm/h | 🌧
| 18:15 |  2.13 mm/h | 🌧 🌧 🌧
| 18:20 |  0.95 mm/h | 🌧
| 18:25 |  0.75 mm/h | 🌧
| 18:30 |  0.00 mm/h |
| 18:35 |  0.00 mm/h |
| 18:40 |  0.00 mm/h |
| 18:45 |  0.00 mm/h |
|-------|------------|---
| 18:50 |  1.45 mm/h | 🌧 🌧
| 18:55 |  0.00 mm/h |
| 19:00 |  0.00 mm/h |
| 19:05 |  0.00 mm/h |
| 19:10 |  0.00 mm/h |
| 19:15 |  0.00 mm/h |
| 19:20 |  0.00 mm/h |
| 19:25 |  0.00 mm/h |
| 19:30 |  0.35 mm/h | 🌧
| 19:35 |  0.45 mm/h | 🌧
| 19:40 |  0.00 mm/h |
| 19:45 |  0.00 mm/h |
```

## TODOs

- [ ] Rainfall graph image
- [ ] Forecast notification
- [ ] Slack integration
