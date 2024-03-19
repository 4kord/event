package event

type Subscription struct {
	subject    string
	id         string
	handler    EventHandler
	dispatcher *Dispatcher
}

func (s *Subscription) Unsubscribe() {
	s.dispatcher.remove(s.subject, s.id)
}

func (s *Subscription) IsValid() bool {
	return s.dispatcher.isValid(s.subject, s.id)
}

func (s *Subscription) Subject() string {
	return s.subject
}

func (s *Subscription) ID() string {
	return s.id
}
