package session

import (
	"context"
	"sync"
	"time"

	"github.com/google/uuid"
)

type SessionData struct {
	UserID    int64
	CreatedAt time.Time
	ExpiresAt time.Time
}

type SessionManager struct {
	sessions map[string]SessionData
	mutex    sync.RWMutex
}

var Manager = NewSessionManager()

// NewSessionManager creates a new session manager.
func NewSessionManager() *SessionManager {
	manager := &SessionManager{
		sessions: make(map[string]SessionData),
	}

	// Start a goroutine to clean up expired sessions
	go manager.cleanupExpiredSessions()

	return manager
}


