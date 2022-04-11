package event

import (
	"github.com/diamondburned/gotrix/matrix"
)

var _ Event = &TagEvent{}

// TagEvent represents an event that informs the client of the tags on a room.
type TagEvent struct {
	EventInfo `json:"-"`

	Tags map[matrix.TagName]matrix.Tag `json:"tags"`
}
