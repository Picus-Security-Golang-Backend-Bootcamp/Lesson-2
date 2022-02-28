package main

import "fmt"

func main() {

	c := Ship{
		Vehicle: Vehicle{
			Name:  "Vapur",
			Color: "Kırmızı",
		},
		EngineType: "Otomatik",
	}

	c.Print()

	c.ChangeColor("Mavi")
	c.Print()
	//c.Vehicle.ChangeColor("string")
}

//struct embedding

// Composition HAS-A
// Inheritance IS-A
type Car struct {
	Vehicle
	GearboxType string
}

func (c Car) Print() {
	fmt.Printf("Name : %s, Color : %s, GearboxType : %s\n", c.Name, c.Color, c.GearboxType)
}

type Ship struct {
	Vehicle
	EngineType string
}

func (c Ship) Print() {
	fmt.Printf("Name : %s, Color : %s, EngineType : %s\n", c.Name, c.Color, c.EngineType)
}

type Vehicle struct {
	Name  string
	Color string
}

func (c *Vehicle) ChangeColor(color string) {
	c.Color = color
}
