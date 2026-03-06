package constants

import "time"

// WebSocket Actions
const (
	JoinUserAction      = "joinUser"
	JoinRoomAction      = "joinRoom"
	LeaveRoomAction     = "leaveRoom"
	StartTypingAction   = "startTyping"
	StopTypingAction    = "stopTyping"
	ToggleOnlineAction  = "toggleOnline"
	ToggleOfflineAction = "toggleOffline"
)

// Emitted Messages
const (
	NewMessageAction       = "new_message"
	AddToTypingAction      = "addToTyping"
	RemoveFromTypingAction = "removeFromTyping"
	ToggleOnlineEmission   = "toggle_online"
	ToggleOfflineEmission  = "toggle_offline"
)

// Type of notification for workspace platform
const (
	NewMessage    = 4
	LowBattery    = 5
	NewTask       = 6
	TaskUpdate    = 7
	DocumentShare = 8
	TeamInvite    = 9
)

const (
	// WriteWait Max wait time when writing message to peer
	WriteWait = 10 * time.Second
)
