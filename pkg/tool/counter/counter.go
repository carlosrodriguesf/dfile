package counter

import "sync"

type Counter struct {
	total float64
	value int
	mutex sync.Mutex
}

func New(total int) Counter {
	return Counter{
		total: float64(total),
	}
}

func (m *Counter) Increment() (value int, percentage float64) {
	m.mutex.Lock()
	m.value++
	m.mutex.Unlock()
	percentage = (float64(m.value) / m.total) * 100
	return m.value, percentage
}
