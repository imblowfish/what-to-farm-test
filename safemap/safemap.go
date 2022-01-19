package safemap

import "sync"

type SafeMap struct {
	mu     sync.Mutex
	prices map[string]string
}

func New() *SafeMap {
	return &SafeMap{
		prices: make(map[string]string),
	}
}

func (m *SafeMap) Set(symbol string, newPrice string) {
	m.mu.Lock()
	m.prices[symbol] = newPrice
	m.mu.Unlock()
}

func (m *SafeMap) Value(symbol string) string {
	m.mu.Lock()
	defer m.mu.Unlock()
	price, ok := m.prices[symbol]
	if !ok {
		return ""
	}
	return price
}

func (m *SafeMap) Get() map[string]string {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.prices
}
