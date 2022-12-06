package threadmanager

type Manager struct {
	track chan bool
}

func NewManager(limit int) *Manager {
	return &Manager{
		track: make(chan bool, limit),
	}
}

func (m *Manager) Lock() {
	m.track <- true
}

func (m *Manager) Release() {
	<-m.track
}

func (m *Manager) Close() {
	close(m.track)
}
