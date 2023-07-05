package abstractfactory

import "fmt"

type Room struct {
	Sides      [4]MapSite
	RoomNumber int
}

func NewRoom(roomNumber int) *Room {
	return &Room{
		RoomNumber: roomNumber,
	}
}

func (r *Room) Enter() bool {
	fmt.Printf("进入房间%d\n", r.RoomNumber)
	return true
}

func (r *Room) SetSide(k Direction, v MapSite) {
	r.Sides[k] = v
}

type Wall struct{}

func NewWall() *Wall {
	return &Wall{}
}

func (w *Wall) Enter() bool {
	fmt.Println("碰到墙")
	return false
}

type Door struct {
	Room1, Room2 IRoom
	IsOpen       bool
}

func NewDoor(room1, room2 IRoom) *Door {
	return &Door{
		Room1: room1,
		Room2: room2,
	}
}

func (d *Door) Enter() bool {
	fmt.Println("碰到门")
	return true
}

func (d *Door) OtherSideFrom(room IRoom) IRoom {
	if room == d.Room1 {
		return d.Room2
	}
	return d.Room1
}

type Maze struct {
	Rooms []IRoom
}

func NewMaze() *Maze {
	return &Maze{
		Rooms: make([]IRoom, 0),
	}
}

func (m *Maze) AddRoom(room IRoom) {
	m.Rooms = append(m.Rooms, room)
}

type MazeGame struct{}

func (m MazeGame) MakeMaze() *Maze {
	return NewMaze()
}

func (m MazeGame) MakeWall() IWall {
	return NewWall()
}

func (m MazeGame) MakeRoom(n int) IRoom {
	return NewRoom(n)
}

func (m MazeGame) MakeDoor(r1 IRoom, r2 IRoom) IDoor {
	return NewDoor(r1, r2)
}
