package gotrix

import (
	"fmt"

	"github.com/diamondburned/gotrix/matrix"
)

// MarkRoomAsDM fetches the DM room list, appends the provided room and reuploads the list.
// It is the caller's duty to make sure only one instance is called at once.
func (c *Client) MarkRoomAsDM(remoteID matrix.UserID, roomID matrix.RoomID) error {
	directEvent, err := c.DMRooms()
	if err != nil {
		return fmt.Errorf("error while marking room as DM: %w", err)
	}

	directEvent.Rooms[remoteID] = append(directEvent.Rooms[remoteID], roomID)
	err = c.DMRoomsSet(directEvent)
	if err != nil {
		return fmt.Errorf("error while marking room as DM: %w", err)
	}
	return nil
}
