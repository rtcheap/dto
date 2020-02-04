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
	Started uint64 `json:"started"`
	Ended   uint64 `json:"ended"`
}

// InProgress callcualated the number of in progress sessions.
func (s SessionStatistics) InProgress() uint64 {
	return s.Started - s.Ended
}

func (s SessionStatistics) String() string {
	return fmt.Sprintf("SessionStatistics(started=%d, ended=%d)", s.Started, s.Ended)
}
