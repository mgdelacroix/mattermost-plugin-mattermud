package mud

import (
	"fmt"
)

// Direction denotes a direction to move
type Direction int

const (
	// North denotes something to the north
	North Direction = iota
	// South denotes something to the south
	South
	// West denotes something to the west
	West
	// East denotes something to the east
	East
	// Up denotes something up, like going upstairs or climbing a ladder
	Up
	// Down denotes something down, like falling through a hole or going down some stairs
	Down
)

// Room stores the information of each room in the game
type Room struct {
	// ID is the unique identifier for this room
	ID string
	// AreaID is the unique identifier for the area
	AreaID string
	// Name shown to the user
	Name string
	// ShortDescription shown to the user when reaching a room
	ShortDescription string
	// LongDescription shown to the user with the command look
	LongDescription string
	// Mobs lists all the mobs present in the room
	Mobs MobList
	// Player lists all players in the room
	Players []*Player
	// Neighbours contains all the neighbour rooms to this one
	Neighbours map[Direction]*RoomDoor
}

// RoomDoor stores information about the transition between a room and the next
type RoomDoor struct {
	// isHidden denotes whether the transition is considered hidden to plain sight
	isHidden bool
	// isInvisible denotes whether the transition is magically invisible
	isInvisible bool
	// isLocked denotes whether the transition has a locked door
	isLocked bool
	// key denotes which key is needed to unlock the door. Empty string would mean the door can be unlocked without any key.
	key string
	// room denotes the room at the other side of the transition
	room *Room
}

// CanMove shows whether there is an open and visible transition from this room in the direction d
func (r *Room) CanMove(d Direction, canSeeHidden, canSeeInvisible bool) bool {
	door, ok := r.Neighbours[d]
	if !ok {
		return false
	}

	if door.isLocked {
		return false
	}

	if door.isHidden && !canSeeHidden {
		return false
	}

	if door.isInvisible && !canSeeInvisible {
		return false
	}

	return true
}

// GetNeighbourRoom returns the room in direction d
func (r *Room) GetNeighbourRoom(d Direction) *Room {
	return r.Neighbours[d].room
}

// CanSeeDoor return whether a locked door can be seen in direction d
func (r *Room) CanSeeDoor(d Direction, canSeeHidden, canSeeInvisible bool) bool {
	door, ok := r.Neighbours[d]
	if !ok {
		return false
	}

	if door.isHidden && !canSeeHidden {
		return false
	}

	if door.isInvisible && !canSeeInvisible {
		return false
	}

	if !door.isLocked {
		return false
	}

	return true
}

func (r *Room) String() string {
	return fmt.Sprintf("%s\n\n%s", r.Name, r.ShortDescription)
}