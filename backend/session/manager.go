package session

import (
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

// CreateSession creates a new session for a user.
func (sm *SessionManager) CreateSession(userID int64, duration time.Duration) string {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()

	sessionID := uuid.New().String()

	now := time.Now()
	sessionData := SessionData{
		UserID:    userID,
		CreatedAt: now,
		ExpiresAt: now.Add(duration),
	}

	sm.sessions[sessionID] = sessionData
	return sessionID
}

// GetUserID retrieves the user ID associated with a session.
func (sm *SessionManager) GetUserID(sessionID string) (int64, bool) {
	sm.mutex.RLock()
	defer sm.mutex.RUnlock()

	sessionData, exists := sm.sessions[sessionID]
	if !exists {
		return 0, false
	}

	// Remove session if session has expired
	if time.Now().After(sessionData.ExpiresAt) {
		go sm.RemoveSession(sessionID)
		return 0, false
	}

	return sessionData.UserID, true
}

// RemoveSession removes a session.
func (sm *SessionManager) RemoveSession(sessionID string) {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()

	delete(sm.sessions, sessionID)
}
