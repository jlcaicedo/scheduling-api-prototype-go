package schedules

import (
	"sync"
	"time"

	"github.com/google/uuid"
)

type Schedule struct {
	ID    string    `json:"id"`
	Title string    `json:"title"`
	Time  time.Time `json:"time"`
}

type Store struct {
	mu   sync.RWMutex
	data map[string]Schedule
}

func NewStore() *Store {
	return &Store{data: make(map[string]Schedule)}
}

func (s *Store) List() []Schedule {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := make([]Schedule, 0, len(s.data))
	for _, v := range s.data {
		out = append(out, v)
	}
	return out
}

func (s *Store) Create(title string, t time.Time) Schedule {
	s.mu.Lock()
	defer s.mu.Unlock()
	id := uuid.NewString()
	sch := Schedule{ID: id, Title: title, Time: t}
	s.data[id] = sch
	return sch
}
