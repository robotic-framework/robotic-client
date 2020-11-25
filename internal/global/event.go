package global

import "github.com/eden-framework/plugin-event/event"

var EventConfig = struct {
	Event *event.MessageBus
}{
	Event: &event.MessageBus{
		Driver: event.EVENT_DRIVER__BUILDIN,
	},
}
