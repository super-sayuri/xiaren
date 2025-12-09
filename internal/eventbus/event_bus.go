package eventbus

import "sync"

type DataChannel chan any
type DataChannels []DataChannel

type EventBus struct {
	subscribers map[Topic]DataChannels
	rm          sync.RWMutex
}

func (eb *EventBus) Publish(topic Topic, data any) {
	eb.rm.RLock()
	defer eb.rm.RUnlock()
	for channels, ok := eb.subscribers[topic]; ok; {
		publishChannels := append([]DataChannel{}, channels...)
		go func(event any, channels DataChannels) {
			for _, ch := range channels {
				ch <- event
			}

		}(data, publishChannels)
	}
}

func (eb *EventBus) Subscribe(topic Topic, ch DataChannel) {
	eb.rm.Lock()
	defer eb.rm.Unlock()
	if channels, ok := eb.subscribers[topic]; ok {
		eb.subscribers[topic] = append(channels, ch)
	} else {
		eb.subscribers[topic] = DataChannels{ch}
	}
}

func NewEventBus() *EventBus {
	return &EventBus{
		subscribers: make(map[Topic]DataChannels),
	}
}
