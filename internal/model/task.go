package model

import (
	"encoding/json"
	"time"
)

// Task represents a work item on the Kanban board (similar to Jira issue)
type Task struct {
	ID          int64        `json:"id,omitempty"          db:"id"`
	ProjectID   int64        `json:"project_id,omitempty"  db:"project_id"`
	StatusID    int64        `json:"status_id,omitempty"   db:"status_id"`
	ParentID    int64        `json:"parent_id,omitempty"   db:"parent_id"` // 0 = root task
	Title       string       `json:"title,omitempty"       db:"title"`
	Description string       `json:"description,omitempty" db:"description"`
	Priority    TaskPriority `json:"priority,omitempty"    db:"priority"`
	Position    float64      `json:"position,omitempty"    db:"position"`
	DueDate     time.Time    `json:"due_date,omitempty"    db:"due_date"`
	StartDate   time.Time    `json:"start_date,omitempty"  db:"start_date"`
	Estimate    int          `json:"estimate,omitempty"    db:"estimate"`    // hours; 0 = not estimated
	CreatedBy   int64        `json:"created_by,omitempty"  db:"created_by"`  // who pressed "Create"
	ReporterID  int64        `json:"reporter_id,omitempty" db:"reporter_id"` // who requested the work
	CompletedAt time.Time    `json:"completed_at,omitempty" db:"completed_at"`
	CreatedAt   time.Time    `json:"created_at,omitempty"  db:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at,omitempty"  db:"updated_at"`
}

// Prepare auto-sets timestamps before insert/update
func (t *Task) Prepare() {
	if t.CreatedAt.IsZero() {
		t.CreatedAt = time.Now()
	}
	t.UpdatedAt = time.Now()
}

// TaskAssignee represents the assignment of a user to a task
type TaskAssignee struct {
	TaskID     int64     `json:"task_id,omitempty"     db:"task_id"`
	UserID     int64     `json:"user_id,omitempty"     db:"user_id"`     // assignee (project_member — validated at app layer)
	AssignedBy int64     `json:"assigned_by,omitempty" db:"assigned_by"` // who performed the assign action
	AssignedAt time.Time `json:"assigned_at,omitempty" db:"assigned_at"`
}

// Label represents a color-coded tag for tasks within a project
type Label struct {
	ID        int64     `json:"id,omitempty"         db:"id"`
	ProjectID int64     `json:"project_id,omitempty" db:"project_id"`
	Name      string    `json:"name,omitempty"       db:"name"`
	Color     string    `json:"color,omitempty"      db:"color"`
	CreatedAt time.Time `json:"created_at,omitempty" db:"created_at"`
}

// TaskLabel represents the many-to-many relation between tasks and labels
type TaskLabel struct {
	TaskID  int64 `json:"task_id,omitempty"  db:"task_id"`
	LabelID int64 `json:"label_id,omitempty" db:"label_id"`
}

// TaskComment represents a comment (or thread reply) on a task
type TaskComment struct {
	ID        int64     `json:"id,omitempty"         db:"id"`
	TaskID    int64     `json:"task_id,omitempty"    db:"task_id"`
	AuthorID  int64     `json:"author_id,omitempty"  db:"author_id"`
	ParentID  int64     `json:"parent_id,omitempty"  db:"parent_id"` // 0 = top-level comment
	Content   string    `json:"content,omitempty"    db:"content"`
	IsEdited  bool      `json:"is_edited,omitempty"  db:"is_edited"`
	CreatedAt time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

// Prepare auto-sets timestamps before insert/update
func (c *TaskComment) Prepare() {
	if c.CreatedAt.IsZero() {
		c.CreatedAt = time.Now()
	}
	c.UpdatedAt = time.Now()
}

// TaskAttachment represents a file attached to a task (references the files table)
type TaskAttachment struct {
	TaskID    int64     `json:"task_id,omitempty"    db:"task_id"`
	FileID    int64     `json:"file_id,omitempty"    db:"file_id"`
	CreatedAt time.Time `json:"created_at,omitempty" db:"created_at"`
}

// TaskLink represents a directed relationship between two tasks
type TaskLink struct {
	ID        int64        `json:"id,omitempty"         db:"id"`
	SourceID  int64        `json:"source_id,omitempty"  db:"source_id"`
	TargetID  int64        `json:"target_id,omitempty"  db:"target_id"`
	LinkType  TaskLinkType `json:"link_type,omitempty"  db:"link_type"`
	CreatedBy int64        `json:"created_by,omitempty" db:"created_by"`
	CreatedAt time.Time    `json:"created_at,omitempty" db:"created_at"`
}

// TaskWatcher represents a user watching a task for notifications
type TaskWatcher struct {
	TaskID    int64     `json:"task_id,omitempty"    db:"task_id"`
	UserID    int64     `json:"user_id,omitempty"    db:"user_id"`
	CreatedAt time.Time `json:"created_at,omitempty" db:"created_at"`
}

// TaskActivityLog records every change made to a task (audit trail)
type TaskActivityLog struct {
	ID        int64           `json:"id,omitempty"         db:"id"`
	TaskID    int64           `json:"task_id,omitempty"    db:"task_id"`
	ActorID   int64           `json:"actor_id,omitempty"   db:"actor_id"`
	Action    string          `json:"action,omitempty"     db:"action"`
	OldValue  json.RawMessage `json:"old_value,omitempty"  db:"old_value"`
	NewValue  json.RawMessage `json:"new_value,omitempty"  db:"new_value"`
	CreatedAt time.Time       `json:"created_at,omitempty" db:"created_at"`
}
