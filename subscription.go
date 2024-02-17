package event

type Subscription struct {
	subject    string
	id         string
	handler    EventHandler
	dispatcher *Dispatcher
}

func (d *Subscription) Unsubscribe() {
	d.dispatcher.remove(d.subject, d.id)
}

func (d *Subscription) IsValid() bool {
	return d.dispatcher.isValid(d.subject, d.id)
}
