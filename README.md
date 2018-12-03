# amefuriso [![CircleCI](https://circleci.com/gh/int128/amefuriso.svg?style=shield)](https://circleci.com/gh/int128/amefuriso)

A command for showing rainfall forecast using [YOLP weather API](https://developer.yahoo.co.jp/webapi/map/openlocalplatform/v1/weather.html).

For example,

```
% ./amefuriso -client-id $YOLP_CLIENT_ID -lat 35.663613 -lon 139.732293
| 11:20 |         |  0.00 mm/h |
| 11:25 |         |  0.75 mm/h | 🌧
| 11:30 |         |  0.00 mm/h |
| 11:35 |         |  0.00 mm/h |
| 11:40 |         |  0.00 mm/h |
| 11:45 |         |  0.00 mm/h |
| 11:50 |         |  0.00 mm/h |
| 11:55 |         |  0.00 mm/h |
| 12:00 |         |  0.00 mm/h |
| 12:05 |         |  1.25 mm/h | 🌧 🌧
| 12:10 |         |  0.00 mm/h |
| 12:15 |         |  0.00 mm/h |
| 12:20 |         |  0.00 mm/h |
| 12:25 |  -6 min |  0.00 mm/h |
| 12:30 |  -1 min |  0.00 mm/h |
| 12:35 |  +4 min |  0.00 mm/h |
| 12:40 |  +9 min |  0.00 mm/h |
| 12:45 | +14 min |  0.00 mm/h |
| 12:50 | +19 min |  0.00 mm/h |
| 12:55 | +24 min |  0.00 mm/h |
| 13:00 | +29 min |  0.00 mm/h |
| 13:05 | +34 min |  0.00 mm/h |
| 13:10 | +39 min |  0.00 mm/h |
| 13:15 | +44 min |  0.00 mm/h |
| 13:20 | +49 min |  0.00 mm/h |
```

## TODOs

- [ ] Rainfall graph image
- [ ] Forecast notification
- [ ] Slack integration
