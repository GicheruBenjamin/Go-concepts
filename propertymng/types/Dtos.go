package types

/*
A property has house/s
A house has  floor/s
A floor has  room/s
A room has  window/s and doors
*/

type Property struct {
	Name string
	Houses []House
}

type House struct {
	Floors []Floor
}

type Floor struct {
	Level int
	Rooms []Room
}

type Room struct {
	Name string
	Windows int
	Doors int
}
