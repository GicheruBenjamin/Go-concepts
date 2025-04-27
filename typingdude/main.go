package main

import (
	"fmt"
)

 

type Food struct {
	Name string
}

type Cook struct {
	Name       string
	FoodCooked []Food
}

func main() {
	sam := Cook{
		Name: "sam",
		FoodCooked: []Food{
			{Name: "Pizza"},
			{Name: "Pasta"},
			{Name: "Salad"},
		},
	}

	var saminfo string = sam.Name + " has cooked " +
		sam.FoodCooked[0].Name + ", " +
		sam.FoodCooked[1].Name + ", " +
		sam.FoodCooked[2].Name

	fmt.Println(saminfo)
}
