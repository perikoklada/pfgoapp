package httpsrv

import (
	"goapp/internal/pkg/watcher"
)

func (s *Server) addWatcher(w *watcher.Watcher) {
	s.watchersLock.Lock()
	defer s.watchersLock.Unlock()
	s.watchers[w.GetWatcherId()] = w
}

func (s *Server) removeWatcher(w *watcher.Watcher) {
	s.watchersLock.Lock()
	defer s.watchersLock.Unlock()
	// Print satistics before removing watcher.
	for i := range s.sessionStats {
		if s.sessionStats[i].id == w.GetWatcherId() {
			s.sessionStats[i].print()

			// Copy last element to the one we are removing and reslice
			s.sessionStats[i] = s.sessionStats[len(s.sessionStats)-1]
			s.sessionStats = s.sessionStats[0 : len(s.sessionStats)-1]

			// Do not continue the iteration. When trying to read the last element
			// of the original slice we will attempt to read out of range.
			break
		}
	}

	// Remove watcher.
	delete(s.watchers, w.GetWatcherId())
}

func (s *Server) notifyWatchers(str string) {
	s.watchersLock.RLock()
	defer s.watchersLock.RUnlock()

	// Send message to all watchers and increment stats.
	for id := range s.watchers {
		s.watchers[id].Send(str)
		s.incStats(id)
	}
}
