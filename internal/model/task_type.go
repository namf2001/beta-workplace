package model

// TaskPriority represents the priority level of a task
type TaskPriority string

const (
	TaskPriorityHighest TaskPriority = "highest"
	TaskPriorityHigh    TaskPriority = "high"
	TaskPriorityMedium  TaskPriority = "medium"
	TaskPriorityLow     TaskPriority = "low"
	TaskPriorityLowest  TaskPriority = "lowest"
)

// String converts to string value
func (p TaskPriority) String() string {
	return string(p)
}

// IsValid checks if TaskPriority is valid
func (p TaskPriority) IsValid() bool {
	return p == TaskPriorityHighest ||
		p == TaskPriorityHigh ||
		p == TaskPriorityMedium ||
		p == TaskPriorityLow ||
		p == TaskPriorityLowest
}

// TaskLinkType represents the relationship type between two linked tasks
type TaskLinkType string

const (
	TaskLinkTypeBlocks      TaskLinkType = "blocks"
	TaskLinkTypeIsBlockedBy TaskLinkType = "is_blocked_by"
	TaskLinkTypeDuplicates  TaskLinkType = "duplicates"
	TaskLinkTypeRelatesTo   TaskLinkType = "relates_to"
)

// String converts to string value
func (t TaskLinkType) String() string {
	return string(t)
}

// IsValid checks if TaskLinkType is valid
func (t TaskLinkType) IsValid() bool {
	return t == TaskLinkTypeBlocks ||
		t == TaskLinkTypeIsBlockedBy ||
		t == TaskLinkTypeDuplicates ||
		t == TaskLinkTypeRelatesTo
}

// ProjectAccess represents the visibility of a project
type ProjectAccess string

const (
	ProjectAccessPublic  ProjectAccess = "public"
	ProjectAccessPrivate ProjectAccess = "private"
)

// String converts to string value
func (a ProjectAccess) String() string {
	return string(a)
}

// IsValid checks if ProjectAccess is valid
func (a ProjectAccess) IsValid() bool {
	return a == ProjectAccessPublic || a == ProjectAccessPrivate
}

// ChannelType represents the type of a chat channel
type ChannelType string

const (
	ChannelTypeGlobal  ChannelType = "global"
	ChannelTypeDM      ChannelType = "dm"
	ChannelTypeGroup   ChannelType = "group"
	ChannelTypeProject ChannelType = "project"
)

// String converts to string value
func (t ChannelType) String() string {
	return string(t)
}

// IsValid checks if ChannelType is valid
func (t ChannelType) IsValid() bool {
	return t == ChannelTypeGlobal ||
		t == ChannelTypeDM ||
		t == ChannelTypeGroup ||
		t == ChannelTypeProject
}
