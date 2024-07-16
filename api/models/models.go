package models

import (
	"time"
)

// User represents a user in the system
type User struct {
	ID               uint `gorm:"primaryKey"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        *time.Time `gorm:"index"`
	Username         string
	Email            string
	Password         string
	Role             string           // roles like 'registered_user', 'moderator', 'admin'
	Listings         []Listing        `gorm:"foreignKey:CreatedByID"`
	Replies          []Reply          `gorm:"foreignKey:CreatedByID"`
	MessagesSent     []PrivateMessage `gorm:"foreignKey:SenderID"`
	MessagesReceived []PrivateMessage `gorm:"foreignKey:ReceiverID"`
}

// Category represents a category in the system
// @Description Represents a category
type Category struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
	Name      string
	IsPublic  bool      // field to distinguish between public and private categories
	Listings  []Listing `gorm:"foreignKey:CategoryID"`
}

// Listing represents a listing in the system
// @Description Represents a listing
type Listing struct {
	ID          uint `gorm:"primaryKey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time `gorm:"index"`
	Title       string
	Content     string    `gorm:"size:5000"` // Limit content to 5000 characters
	Pictures    []Picture `gorm:"foreignKey:ListingID"`
	CategoryID  uint
	IsPublic    bool
	IsHidden    bool    // field to handle temporary hiding by moderators
	Replies     []Reply `gorm:"foreignKey:ListingID"`
	CreatedByID uint
	CreatedBy   User `gorm:"foreignKey:CreatedByID"`
}

// Reply represents a reply to a listing
// @Description Represents a reply to a listing
type Reply struct {
	ID          uint `gorm:"primaryKey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time `gorm:"index"`
	Content     string     `gorm:"size:255"` // Limit reply content to 255 characters
	ListingID   uint
	CreatedByID uint
	CreatedBy   User `gorm:"foreignKey:CreatedByID"`
}

// PrivateMessage represents a private message between users
// @Description Represents a private message between users
type PrivateMessage struct {
	ID         uint `gorm:"primaryKey"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time `gorm:"index"`
	Content    string
	SenderID   uint
	Sender     User `gorm:"foreignKey:SenderID"`
	ReceiverID uint
	Receiver   User `gorm:"foreignKey:ReceiverID"`
}

// Picture represents a picture in a listing
// @Description Represents a picture in a listing
type Picture struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
	URL       string
	ListingID uint
	Listing   Listing `gorm:"foreignKey:ListingID"`
}
