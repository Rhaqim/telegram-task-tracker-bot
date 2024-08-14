package repo

import "time"

type TrackingRequest struct {
	UserID       int64
	ChatID       int64
	TrackingInfo string
	Timestamp    time.Time
}
