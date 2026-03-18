package model

import "time"

// Channel represents a chat channel (global, DM, group, or project-linked)
type Channel struct {
	ID          int64       `json:"id,omitempty"           db:"id"`
	WorkplaceID int64       `json:"workplace_id,omitempty" db:"workplace_id"` // 0 = DM channel
	ProjectID   int64       `json:"project_id,omitempty"   db:"project_id"`   // 0 = not a project channel
	Name        string      `json:"name,omitempty"         db:"name"`
	Type        ChannelType `json:"type,omitempty"         db:"type"`
	CreatedBy   int64       `json:"created_by,omitempty"   db:"created_by"`
	CreatedAt   time.Time   `json:"created_at,omitempty"   db:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at,omitempty"   db:"updated_at"`
}

// Prepare auto-sets timestamps before insert/update
func (c *Channel) Prepare() {
	if c.CreatedAt.IsZero() {
		c.CreatedAt = time.Now()
	}
	c.UpdatedAt = time.Now()
}

// ChannelMember represents a user's membership in a channel
type ChannelMember struct {
	ID         int64     `json:"id,omitempty"           db:"id"`
	ChannelID  int64     `json:"channel_id,omitempty"   db:"channel_id"`
	UserID     int64     `json:"user_id,omitempty"      db:"user_id"`
	LastReadAt time.Time `json:"last_read_at,omitempty" db:"last_read_at"`
	JoinedAt   time.Time `json:"joined_at,omitempty"    db:"joined_at"`
}

// Message represents a chat message in a channel
type Message struct {
	ID        int64     `json:"id,omitempty"         db:"id"`
	ChannelID int64     `json:"channel_id,omitempty" db:"channel_id"`
	SenderID  int64     `json:"sender_id,omitempty"  db:"sender_id"`
	ParentID  int64     `json:"parent_id,omitempty"  db:"parent_id"` // 0 = top-level message
	Content   string    `json:"content,omitempty"    db:"content"`
	IsEdited  bool      `json:"is_edited,omitempty"  db:"is_edited"`
	IsDeleted bool      `json:"is_deleted,omitempty" db:"is_deleted"`
	CreatedAt time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

// Prepare auto-sets timestamps before insert/update
func (m *Message) Prepare() {
	if m.CreatedAt.IsZero() {
		m.CreatedAt = time.Now()
	}
	m.UpdatedAt = time.Now()
}

// MessageAttachment represents a file attached to a message (references the files table)
type MessageAttachment struct {
	MessageID int64     `json:"message_id,omitempty" db:"message_id"`
	FileID    int64     `json:"file_id,omitempty"    db:"file_id"`
	CreatedAt time.Time `json:"created_at,omitempty" db:"created_at"`
}

// MessageReaction represents an emoji reaction to a message
type MessageReaction struct {
	ID        int64     `json:"id,omitempty"         db:"id"`
	MessageID int64     `json:"message_id,omitempty" db:"message_id"`
	UserID    int64     `json:"user_id,omitempty"    db:"user_id"`
	Emoji     string    `json:"emoji,omitempty"      db:"emoji"`
	CreatedAt time.Time `json:"created_at,omitempty" db:"created_at"`
}
