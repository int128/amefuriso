package domain

type Forecast struct {
	Location      Location
	RainWillStart *Event
	RainWillStop  *Event
}

func (forecast *Forecast) HasAnyTopic() bool {
	return forecast.RainWillStart != nil || forecast.RainWillStop != nil
}

func NewForecast(weather Weather) Forecast {
	forecast := Forecast{Location: weather.Location}
	if start := weather.FindRainStarts(); !weather.IsRainingNow() && start != nil {
		forecast.RainWillStart = start
	}
	if stop := weather.FindRainStops(); weather.IsRainingNow() && stop != nil {
		forecast.RainWillStop = stop
	}
	return forecast
}
