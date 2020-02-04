package dto

import "fmt"

// Session user session.
type Session struct {
	UserID string `json:"userId,omitempty"`
	Key    string `json:"key,omitempty"`
}

func (s Session) String() string {
	return fmt.Sprintf("Session(userId=%s)", s.UserID)
}

// SessionStatistics statistics for the numbers of handled sessions
type SessionStatistics struct {
	Count uint64 `json:"count"`
}
