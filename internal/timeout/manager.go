
package timeout

import "time"

type Manager struct {
	Default time.Duration
}

func NewManager() *Manager {

	return &Manager{
		Default: 120 * time.Millisecond,
	}
}

func (m *Manager) Timeout(score float64) time.Duration {

	if score > 0.8 {
		return 150 * time.Millisecond
	}

	if score > 0.5 {
		return 120 * time.Millisecond
	}

	return 80 * time.Millisecond
}
