package domain

type RainWillStartMessage struct {
	Location Location
	Event    Event
}

type RainWillStopMessage struct {
	Location Location
	Event    Event
}

const NoForecastMessage = "NoForecastMessage"

func NewForecastMessage(weather Weather) interface{} {
	if start := weather.FindRainStarts(); !weather.IsRainingNow() && start != nil {
		return RainWillStartMessage{
			Location: weather.Location,
			Event:    *start,
		}
	}
	if stop := weather.FindRainStops(); weather.IsRainingNow() && stop != nil {
		return RainWillStopMessage{
			Location: weather.Location,
			Event:    *stop,
		}
	}
	return NoForecastMessage
}

type Message struct {
	Text     string
	ImageURL string
}
