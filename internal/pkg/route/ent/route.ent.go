package ent

import "gorm.io/gorm"

// Route represents a collection of milestones to be visited.
type Route struct {
	gorm.Model
	Title      string           `gorm:"not null"`                                       // Route title, e.g., "Evening check"
	Milestones []RouteMilestone `gorm:"foreignKey:RouteID;constraint:OnDelete:CASCADE"` // One Route has many milestones
}

// RouteMilestone represents a specific Post within a Route.
type RouteMilestone struct {
	gorm.Model
	RouteID  uint `gorm:"not null;index"` // Foreign key to Route
	PostID   uint `gorm:"not null"`       // Foreign key to Post
	Priority uint `gorm:"not null"`       // Visit priority/order, starting from 1
}
