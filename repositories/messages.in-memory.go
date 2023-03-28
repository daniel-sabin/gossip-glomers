package repositories

import "sync"

type MessagesRepositoryInMemory struct {
	store map[int]struct{}
	mutex sync.RWMutex
}

func NewMessagesRepositoryInMemory() *MessagesRepositoryInMemory {
	return &MessagesRepositoryInMemory{
		store: make(map[int]struct{}),
	}
}

func (m *MessagesRepositoryInMemory) Save(id int) {
	m.mutex.Lock()
	m.store[id] = struct{}{}
	m.mutex.Unlock()
}

func (m *MessagesRepositoryInMemory) MessageExists(id int) bool {
	m.mutex.Lock()
	_, exists := m.store[id]
	m.mutex.Unlock()

	return exists
}

func (m *MessagesRepositoryInMemory) Messages() map[int]struct{} {
	return m.store
}

func (m *MessagesRepositoryInMemory) MessagesCount() int {
	m.mutex.RLock()
	count := len(m.store)
	m.mutex.RUnlock()

	return count
}
