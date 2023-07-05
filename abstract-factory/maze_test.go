package abstractfactory

import "fmt"

func CreateMaze(factory GameFactory) *Maze {
	maze := factory.MakeMaze()
	room1 := factory.MakeRoom(1)
	room2 := factory.MakeRoom(2)
	door := factory.MakeDoor(room1, room2)

	maze.AddRoom(room1)
	maze.AddRoom(room2)

	room1.SetSide(North, factory.MakeWall())
	room1.SetSide(East, door)
	room1.SetSide(South, factory.MakeWall())
	room1.SetSide(West, factory.MakeWall())

	room2.SetSide(North, factory.MakeWall())
	room2.SetSide(East, factory.MakeWall())
	room2.SetSide(South, factory.MakeWall())
	room2.SetSide(West, door)

	return maze
}

func Example_mazeGameFactory() {
	maze := CreateMaze(&MazeGame{})
	fmt.Println("普通迷宫：")
	for _, room := range maze.Rooms {
		room := room.(*Room)
		room.Enter()
		for direction, side := range room.Sides {
			if side == nil {
				fmt.Printf("房间%d的%s方向是墙\n", room.RoomNumber, DirectionNames[direction])
			} else {
				fmt.Printf("房间%d的方向：%s\t", room.RoomNumber, DirectionNames[direction])
				side.Enter()
			}
		}
	}
	// Output:
	// 普通迷宫：
	// 进入房间1
	// 房间1的方向：北	碰到墙
	// 房间1的方向：南	碰到墙
	// 房间1的方向：东	碰到门
	// 房间1的方向：西	碰到墙
	// 进入房间2
	// 房间2的方向：北	碰到墙
	// 房间2的方向：南	碰到墙
	// 房间2的方向：东	碰到墙
	// 房间2的方向：西	碰到门
}

func Example_enchantedMazeGameFactory() {
	enchantedMaze := CreateMaze(&EnchantedMazeGame{})
	fmt.Println("有咒语的迷宫：")
	for _, room := range enchantedMaze.Rooms {
		room := room.(*EnchantedRoom)
		room.Enter()
		for direction, side := range room.Sides {
			if side == nil {
				fmt.Printf("房间%d的%s方向是墙\n", room.RoomNumber, DirectionNames[direction])
			} else {
				fmt.Printf("房间%d的方向：%s\t", room.RoomNumber, DirectionNames[direction])
				side.Enter()
			}
		}
	}
	// Output:
	// 有咒语的迷宫：
	// 进入房间1,拿起咒语卷轴
	// 房间1的方向：北	碰到墙
	// 房间1的方向：南	碰到墙
	// 房间1的方向：东	碰到需要咒语的门
	// 房间1的方向：西	碰到墙
	// 进入房间2,拿起咒语卷轴
	// 房间2的方向：北	碰到墙
	// 房间2的方向：南	碰到墙
	// 房间2的方向：东	碰到墙
	// 房间2的方向：西	碰到需要咒语的门
}

func Example_bombedMazeGameFactory() {
	bombedMaze := CreateMaze(&BombedMazeGame{})
	fmt.Println("有炸弹的迷宫：")
	for _, room := range bombedMaze.Rooms {
		room := room.(*RoomWithABomb)
		room.Enter()
		for direction, side := range room.Sides {
			if side == nil {
				fmt.Printf("房间%d的%s方向是墙\n", room.RoomNumber, DirectionNames[direction])
			} else {
				fmt.Printf("房间%d的方向：%s\t", room.RoomNumber, DirectionNames[direction])
				side.Enter()
			}
		}
	}
	// Output:
	// 有炸弹的迷宫：
	// 进入有炸弹的房间1
	// 房间1的方向：北	碰到炸弹墙
	// 房间1的方向：南	碰到炸弹墙
	// 房间1的方向：东	碰到门
	// 房间1的方向：西	碰到炸弹墙
	// 进入有炸弹的房间2
	// 房间2的方向：北	碰到炸弹墙
	// 房间2的方向：南	碰到炸弹墙
	// 房间2的方向：东	碰到炸弹墙
	// 房间2的方向：西	碰到门
}
