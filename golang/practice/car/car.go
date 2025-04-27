package car

import "fmt"

type Vehicle struct {
	Make  string
	Model string
	Year  int
}

type Car struct {
	Vehicle
	Seats int
}

func DisplayInfoPtr(c *Car) {
	fmt.Printf("Car(make=%s, model=%s, year=%d, seats=%d)\n", c.Make, c.Model, c.Year, c.Seats)
}

func (c *Car) DisplayInfoVal() {
	fmt.Printf("Car(make=%s, model=%s, year=%d, seats=%d)\n", c.Make, c.Model, c.Year, c.Seats)
}
