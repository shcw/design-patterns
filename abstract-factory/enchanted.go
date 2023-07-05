package abstractfactory

import "fmt"

type EnchantedMazeGame struct {
	MazeGame
}

func (g *EnchantedMazeGame) MakeRoom(roomNumber int) IRoom {
	return NewEnchantedRoom(roomNumber, &Spell{Key: "456"})
}

func (g *EnchantedMazeGame) MakeDoor(room1, room2 IRoom) IDoor {
	return NewDoorNeedingSpell(room1, room2)
}

// EnchantedRoom 魔法屋子
type EnchantedRoom struct {
	Room
	Spell *Spell
}

func NewEnchantedRoom(roomNumber int, spell *Spell) *EnchantedRoom {
	return &EnchantedRoom{
		Room:  *NewRoom(roomNumber),
		Spell: spell,
	}
}

func (r *EnchantedRoom) Enter() bool {
	if r.HasSpell() {
		fmt.Printf("进入房间%d,拿起咒语卷轴\n", r.RoomNumber)
		return true
	}
	return r.Room.Enter()
}

func (r *EnchantedRoom) HasSpell() bool {
	return r.Spell != nil
}

func (r *EnchantedRoom) PickUpSpell() *Spell {
	return r.Spell
}

type DoorNeedingSpell struct {
	Door
	Spell *Spell
}

func NewDoorNeedingSpell(room1, room2 IRoom) *DoorNeedingSpell {
	return &DoorNeedingSpell{
		Door:  *NewDoor(room1, room2),
		Spell: &Spell{Key: "123"},
	}
}

func (d *DoorNeedingSpell) Enter() bool {
	if d.HasSpell() {
		fmt.Println("碰到需要咒语的门")
		return true
	}
	return d.Door.Enter()
}

func (d *DoorNeedingSpell) HasSpell() bool {
	return d.Spell != nil
}

func (d *DoorNeedingSpell) CastSpell() {
	fmt.Printf("使用咒语%s开启门\n", d.Spell.Key)
}
