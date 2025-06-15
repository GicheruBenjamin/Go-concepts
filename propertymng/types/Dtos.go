package types
// Data transfer objects
// types/Dtos.go
/*
A property has house/s
A house has  floor/s
A floor has  room/s
A room has  window/s and doors
*/

import (
	"fmt"
)

type Property struct {
	Name string
	Houses []House
}

func (prop *Property) GetHouses() []House {
	return prop.Houses
}

type House struct {
	Floors []Floor
}

func (house *House) GetFloors() []Floor {
	return house.Floors
}

type Floor struct {
	Level int
	Rooms []Room
}

func (floor *Floor) GetRooms() []Room {
	return floor.Rooms
}

type Room struct {
	Name string
	Windows int
	Doors int
}

func (room *Room) GetName() string {
	return room.Name
}

func (room *Room) GetWindows() int {
	return room.Windows
}

func (room *Room) GetDoors() int {
	return room.Doors
}

func Displayproperty(prop Property) {
	fmt.Println("Property Name: " + prop.Name)
	for _, house := range prop.Houses {
		fmt.Println("House Name: " + house.GetFloors()[0].Rooms[0].GetName())
	}
}