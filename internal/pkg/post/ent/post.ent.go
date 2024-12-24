package ent

import (
	"app/internal/pkg/record/ent"
	"gorm.io/gorm"
)

// Post represents the main entity, such as a physical location or asset.
type Post struct {
	gorm.Model
	Title       string        `gorm:"not null"` // Title is required, e.g., "Master Door"
	Description *string       // Optional description, e.g., "office main entrance"
	Phrase      string        // Optional secret phrase, like an admin password
	Location    *PostLocation `gorm:"foreignKey:PostID;constraint:OnDelete:CASCADE"` // One-to-One optional relation
	Records     []ent.Record  `gorm:"foreignKey:PostID;constraint:OnDelete:CASCADE"` // One Post has many Records
}

// PostLocation provides additional location-related metadata for a Post.
type PostLocation struct {
	gorm.Model
	PostID    uint    `gorm:"not null;uniqueIndex"` // Foreign key to Post, one-to-one relation
	Comment   *string // Optional comment, e.g., "Door code is 0808#"
	Country   string  // Country code, e.g., KZ, US
	City      string  // City name, e.g., Almaty
	Address   string  // Main address, e.g., Some st. 21
	Latitude  float64 // Latitude coordinate (optional)
	Longitude float64 // Longitude coordinate (optional)
}
