package data

import "fmt"

type Fruit struct {
	Name        string
	Description string
	Price       float64
}

func generateFruit(fruit Fruit) []string {

	froot := []string{}

	froot = append(froot, fruit.Name)
	froot = append(froot, fruit.Description)
	froot = append(froot, fmt.Sprintf("%2f", fruit.Price))

	return froot
}

func FruitList() [][]string {

	fruits := []Fruit{
		{
			Name:        "Apple",
			Description: "Red and juicy",
			Price:       2.00,
		},
		{
			Name:        "Orange",
			Description: "Orange and juicy",
			Price:       3.00,
		},
		{
			Name:        "Mango",
			Description: "Purple and Yellow",
			Price:       3.5,
		},
	}

	var froots [][]string

	for _, fruit := range fruits {
		froots = append(froots, generateFruit(fruit))
	}

	return froots
}
