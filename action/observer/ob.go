package main

import "fmt"

type EventListener interface {
	Update(v interface{})
}

type AListener struct{}

func (AListener) Update(v interface{}) {
	fmt.Println("aa")
}

type EventManager struct {
	ListenerMap map[string]EventListener
}

func NewEventManager() *EventManager {
	return &EventManager{
		ListenerMap: make(map[string]EventListener, 16),
	}
}

func (e *EventManager) Subscribe(name string, listener EventListener) {
	e.ListenerMap[name] = listener
}

func (e *EventManager) UnSubscribe(name string) {
	if _, ok := e.ListenerMap[name]; !ok {
		return
	}
	delete(e.ListenerMap, name)
}

func (e *EventManager) Notify() {
	for _, listener := range e.ListenerMap {
		listener.Update(333)
	}
}

func main() {
	em := NewEventManager()
	em.Subscribe("aa", &AListener{})
	em.Notify()
}
