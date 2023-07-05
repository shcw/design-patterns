package abstractfactory

import (
	"fmt"
)

type BombedMazeGame struct {
	MazeGame
}

func (g *BombedMazeGame) MakeRoom(roomNumber int) IRoom {
	return NewRoomWithABomb(roomNumber, true)
}

func (g *BombedMazeGame) MakeWall() IWall {
	return NewBombedWall(true)
}

// BombedWall 炸弹墙
type BombedWall struct {
	Wall
	Bomb bool
}

func NewBombedWall(bombed bool) *BombedWall {
	return &BombedWall{
		Wall: *NewWall(),
		Bomb: bombed,
	}
}

func (w *BombedWall) Enter() bool {
	if w.Bomb {
		fmt.Println("碰到炸弹墙")
		return false
	}
	return w.Wall.Enter()
}

// RoomWithABomb 有炸弹的屋子
type RoomWithABomb struct {
	*Room
	Bomb bool
}

func NewRoomWithABomb(roomNumber int, bombed bool) *RoomWithABomb {
	return &RoomWithABomb{
		Room: NewRoom(roomNumber),
		Bomb: bombed,
	}
}

func (r *RoomWithABomb) Enter() bool {
	if r.Bomb {
		fmt.Printf("进入有炸弹的房间%d\n", r.RoomNumber)
		return true
	}
	return r.Room.Enter()
}
