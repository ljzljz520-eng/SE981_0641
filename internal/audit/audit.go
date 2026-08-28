package audit

import (
	"fmt"
	"sync"
	"time"
)

type Event struct {
	Action, Entity, Actor string
	At                    time.Time
}
type Log struct {
	mu     sync.Mutex
	events []Event
}

func New() *Log { return &Log{events: []Event{}} }
func (l *Log) Record(action, entity, actor string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.events = append(l.events, Event{action, entity, actor, time.Unix(0, 0)})
}
func (l *Log) Events() []Event {
	l.mu.Lock()
	defer l.mu.Unlock()
	return append([]Event{}, l.events...)
}
func (e Event) String() string { return fmt.Sprintf("%s:%s:%s", e.Action, e.Entity, e.Actor) }
