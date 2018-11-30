# amefuriso

A command for showing rainfall forecast using [YOLP weather API](https://developer.yahoo.co.jp/webapi/map/openlocalplatform/v1/weather.html).

For example,

```
% ./amefuriso -client-id $YOLP_CLIENT_ID -lat 37.15918532117067 -lon 138.28260262402358
2018/11/30 13:55:35 @{37.159185 138.2826}	observation	13:45	0.00 mm
2018/11/30 13:55:35 @{37.159185 138.2826}	forecast	13:55	0.55 mm
2018/11/30 13:55:35 @{37.159185 138.2826}	forecast	14:05	0.00 mm
2018/11/30 13:55:35 @{37.159185 138.2826}	forecast	14:15	0.00 mm
2018/11/30 13:55:35 @{37.159185 138.2826}	forecast	14:25	1.15 mm
2018/11/30 13:55:35 @{37.159185 138.2826}	forecast	14:35	0.00 mm
2018/11/30 13:55:35 @{37.159185 138.2826}	forecast	14:45	0.00 mm
```
