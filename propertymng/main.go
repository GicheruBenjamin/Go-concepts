package main

import (
	"fmt"
	"hsmd/types"
)


func main() {
	fmt.Println("  Property Management system  ")
	fmt.Println("Upcountry Property")
	types.Displayproperty(Upcountryproperty)
	fmt.Println("Holiday Home Property")
	types.Displayproperty(Holidayhomeproperty)
	fmt.Println("Main Property")
	types.Displayproperty(Mainproperty)

}



var Upcountryproperty = types.Property{
	Name: "Upcountry",
	Houses: []types.House{
		{
			Floors: []types.Floor{
				{
					Level: 1,
					Rooms: []types.Room{
						{
							Name: "Kitchen",
							Windows: 2,
							Doors: 2,
						},
						{
							Name: "Living Room",
							Windows: 2,
							Doors: 2,
						},
					},
				},
				{
					Level: 3,
					Rooms: []types.Room{
						{
							// Name: "Kitchen",
							Windows: 2,
							Doors: 4,
						},
						{
							// Name: "Living Room",
							Windows: 2,
							Doors: 4,
						},
					},
				},
			},
		},
		{
			Floors: []types.Floor{
				{
					Level: 2,
					Rooms: []types.Room{
						{
							//Name: "Kitchen",
							Windows: 2,
							Doors: 2,
						},
						{
							//Name: "Living Room",
							Windows: 2,
							Doors: 2,
						},
					},
				},
				{
					Level: 5,
					Rooms: []types.Room{
						{
							//Name: "Bedroom",
							Windows: 2,
							Doors: 6,
						},
						{
							//Name: "Kitchen",
							Windows: 2,
							Doors: 6,
						},
					},
				},
			},
		},
	},
}

var Holidayhomeproperty = types.Property{
	Name: "Holiday Home",
	Houses :[]types.House{
		{
			Floors: []types.Floor{
				{
					Level: 1,
					Rooms: []types.Room{
						{
							Name: "Kitchen",
							Windows: 2,
							Doors: 2,
						},
						{
							Name: "Living Room",
							Windows: 2,
							Doors: 2,
						},
					},
				},
				{
					Level: 3,
					Rooms: []types.Room{
						{
							Name: "Dining Room",
							Windows: 2,
							Doors: 4,
						},
						{
							Name: "Living Room",
							Windows: 2,
							Doors: 4,
						},
					},
				},
			},
		},
	},
}

var Mainproperty = types.Property{
	Houses: []types.House{
		{
			Floors: []types.Floor{
				{
					Level: 1,
					Rooms: []types.Room{
						{
							Name: "Kitchen",
							Windows: 2,
							Doors: 2,
						},
						{
							Name: "Living Room",
							Windows: 2,
							Doors: 2,
						},
					},
				},
				{
					Level: 3,
					Rooms: []types.Room{
						{
							Name: "Bathroom",
							Windows: 2,
							Doors: 4,
						},
						{
							Name: "Living Room",
							Windows: 2,
							Doors: 4,
						},
					},
				},
			},
		},
		{
			Floors: []types.Floor{
				{
					Level: 2,
					Rooms: []types.Room{
						{
							// Name: "Kitchen",
							Windows: 2,
							Doors: 2,
						},
						{
							//Name: "Living Room",
							Windows: 2,
							Doors: 2,
						},
					},
				},
				{
					Level: 5,
					Rooms: []types.Room{
						{
							//Name: "Bedroom",
							Windows: 2,
							Doors: 6,
						},
						{
							//Name: "Kitchen",
							Windows: 2,
							Doors: 6,
						},
					},
				},
			},
		},
	},
}