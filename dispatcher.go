package event

import (
	"sync"

	"github.com/google/uuid"
)

type EventHandler func(e interface{})

type Dispatcher struct {
	mu        sync.RWMutex
	callbacks map[string][]*Subscription
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		callbacks: make(map[string][]*Subscription),
	}
}

func (d *Dispatcher) Subscribe(subject string, h EventHandler) *Subscription {
	d.mu.Lock()
	defer d.mu.Unlock()

	sub := &Subscription{
		subject:    subject,
		id:         uuid.New().String(),
		handler:    h,
		dispatcher: d,
	}
	d.callbacks[subject] = append(d.callbacks[subject], sub)

	return sub
}

func (d *Dispatcher) Emit(subject string, v interface{}) {
	d.mu.RLock()
	defer d.mu.RUnlock()

	if subs, ok := d.callbacks[subject]; ok {
		for i := range subs {
			subs[i].handler(v)
		}
	}
}

func (d *Dispatcher) remove(subject string, id string) {
	d.mu.Lock()
	defer d.mu.Unlock()

	for i, subscription := range d.callbacks[subject] {
		if subscription.id == id {
			d.callbacks[subject] = append(d.callbacks[subject][:i], d.callbacks[subject][i+1:]...)
			break
		}
	}
}

func (d *Dispatcher) isValid(subject string, id string) bool {
	d.mu.Lock()
	defer d.mu.Unlock()

	for _, subscription := range d.callbacks[subject] {
		if subscription.id == id {
			return true
		}
	}

	return false
}
