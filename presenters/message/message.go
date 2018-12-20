package message

import (
	"fmt"

	"github.com/int128/amefurisobot/domain"
)

func Format(message interface{}, weatherURL string) string {
	if m, ok := message.(domain.RainWillStartMessage); ok {
		return fmt.Sprintf("%s で %s から雨が降る見込みです。最新の予報は %s",
			m.Location.Name,
			m.Event.Time.Format("15:04"),
			weatherURL)
	}
	if m, ok := message.(domain.RainWillStopMessage); ok {
		return fmt.Sprintf("%s で %s に雨が止む見込みです。最新の予報は %s",
			m.Location.Name,
			m.Event.Time.Format("15:04"),
			weatherURL)
	}
	return "予報はありません"
}
