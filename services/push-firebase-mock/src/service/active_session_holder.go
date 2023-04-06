package service

import "sync"

type ActiveSessionHolder struct {
	connections []string
	mx          sync.RWMutex
}

func (a *ActiveSessionHolder) AddSession(sessionId string) {
	a.mx.Lock()
	defer a.mx.Unlock()
	a.connections = append(a.connections, sessionId)
}

func (a *ActiveSessionHolder) RemoveSession(sessionId string) {
	a.mx.Lock()
	defer a.mx.Unlock()
	a.connections = remove(a.connections, sessionId)
}

func (a *ActiveSessionHolder) GetActiveSessions() []string {
	a.mx.RLock()
	defer a.mx.RUnlock()
	out := make([]string, len(a.connections))
	copy(out, a.connections)
	return out
}

func remove(l []string, item string) []string {
	for i, val := range l {
		if val == item {
			return append(l[:i], l[i+1:]...)
		}
	}
	return l
}
