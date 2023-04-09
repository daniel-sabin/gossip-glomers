package repositories

import "sync"

type MessagesRepositoryInMemory struct {
	store map[int]struct{}
	ids   []int
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

func (m *MessagesRepositoryInMemory) Messages() []int {
	m.ids = make([]int, 0, m.MessagesCount())
	m.mutex.Lock()
	for id := range m.store {
		m.ids = append(m.ids, id)
	}
	m.mutex.Unlock()

	return m.ids
}

func (m *MessagesRepositoryInMemory) MessagesCount() int {
	m.mutex.RLock()
	count := len(m.store)
	m.mutex.RUnlock()

	return count
}
