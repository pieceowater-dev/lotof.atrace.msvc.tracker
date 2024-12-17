package ent

import (
	"gorm.io/gorm"
	"time"
)

// RecordMethod defines the method used to create a Record.
type RecordMethod int

const (
	MethodUnknown    RecordMethod = iota // Default unknown method
	MethodPostPhrase                     // Created using Post's secret phrase
	MethodQR                             // Created using dynamic QR code
	MethodQRStatic                       // Created using static QR code
)

// Record represents a log entry related to a Post.
type Record struct {
	gorm.Model
	PostID    uint         `gorm:"not null;index"` // Foreign key to Post
	UserID    string       `gorm:"not null"`       // User UUID from an external service
	Timestamp time.Time    `gorm:"autoCreateTime"` // Timestamp in UTC
	Method    RecordMethod `gorm:"default:0"`      // Record creation method
}
