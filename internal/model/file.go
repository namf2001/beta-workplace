package model

import "time"

// File represents an uploaded file stored in object storage (e.g. S3)
type File struct {
	ID         int64     `json:"id,omitempty"          db:"id"`
	UploadedBy int64     `json:"uploaded_by,omitempty" db:"uploaded_by"`
	Bucket     string    `json:"bucket,omitempty"      db:"bucket"`   // storage bucket name
	FileKey    string    `json:"file_key,omitempty"    db:"file_key"` // object key / path inside bucket
	FileURL    string    `json:"file_url,omitempty"    db:"file_url"` // public access URL
	FileName   string    `json:"file_name,omitempty"   db:"file_name"`
	FileSize   int64     `json:"file_size,omitempty"   db:"file_size"` // bytes
	MimeType   string    `json:"mime_type,omitempty"   db:"mime_type"`
	CreatedAt  time.Time `json:"created_at,omitempty"  db:"created_at"`
}

// Prepare auto-sets created_at before insert
func (f *File) Prepare() {
	if f.CreatedAt.IsZero() {
		f.CreatedAt = time.Now()
	}
}
