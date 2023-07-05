package abstractfactory

// 根据《设计模式：可复用面向对象软件的基础》的迷宫案例编写

type Direction int

const (
	North Direction = iota
	South
	East
	West
)

var DirectionNames = [...]string{"北", "南", "东", "西"}

type Spell struct {
	Key string
}

type MapSite interface {
	Enter() bool
}

type IRoom interface {
	MapSite
	SetSide(Direction, MapSite)
}

type IWall interface {
	MapSite
}

type IDoor interface {
	MapSite
}

type GameFactory interface {
	MakeMaze() *Maze
	MakeWall() IWall
	MakeRoom(n int) IRoom
	MakeDoor(r1 IRoom, r2 IRoom) IDoor
}
