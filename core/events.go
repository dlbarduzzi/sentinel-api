package core

import "github.com/dlbarduzzi/sentinel/tools/event"

type EventRequest struct {
	App App
	event.Event
}
