package model

import "time"

// Friendship represents a friend relationship between two users
type Friendship struct {
	ID          int64            `json:"id,omitempty"           db:"id"`
	RequesterID int64            `json:"requester_id,omitempty" db:"requester_id"`
	ReceiverID  int64            `json:"receiver_id,omitempty"  db:"receiver_id"`
	Status      FriendshipStatus `json:"status,omitempty"       db:"status"`
	CreatedAt   time.Time        `json:"created_at,omitempty"   db:"created_at"`
	UpdatedAt   time.Time        `json:"updated_at,omitempty"   db:"updated_at"`
}

// Prepare auto-sets timestamps before insert/update
func (f *Friendship) Prepare() {
	if f.CreatedAt.IsZero() {
		f.CreatedAt = time.Now()
	}
	f.UpdatedAt = time.Now()
}
