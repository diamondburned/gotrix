package event

import (
	"github.com/diamondburned/gotrix/matrix"
)

var _ StateEvent = &RoomTombstoneEvent{}

// RoomTombstoneEvent is an event where the current room has been upgraded and a new room should be used instead.
type RoomTombstoneEvent struct {
	StateEventInfo  `json:"-"`
	Message         string        `json:"body,omitempty"`
	ReplacementRoom matrix.RoomID `json:"replacement_room,omitempty"`
}
