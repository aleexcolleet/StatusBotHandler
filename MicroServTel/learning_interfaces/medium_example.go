package main

import "fmt"

/*
	Medium_Example for interfaces learning.
	There's three types of vehicles: [car, truck, motorcycle]
	There's three methods: [Start(), Stop(), Details()]

	I need an interface (vehicle) that outline the common methods
	This way with an array of vehicles I can perform all functions adequately
	with a loop
*/

type Vehicle interface {
	Start()
	Stop()
	Details()
}

type Car struct {
	numOfWheels int
	horsePower  int
	numOfDoors  int
}

type Truck struct {
	numOfWheels  int
	horsePower   int
	numOfDoors   int
	loadCapacity int
}
type Motorcycle struct {
	numOfWheels int
	horsePower  int
}

// CAR METHODS
func (c *Car) Start() {
	fmt.Println("Car started\n")
}
func (c *Car) Stop() {
	fmt.Println("Car stopped\n")
}
func (c *Car) Details() {
	fmt.Printf("-------- Car Details ----------\n")
	fmt.Printf("numOfWheels: %d\n", c.numOfWheels)
	fmt.Printf("horsePower: %d Cv\n", c.horsePower)
	fmt.Printf("numOfDoors: %d\n", c.numOfDoors)
	fmt.Printf("-------------------------------\n")
}

// TRUCK METHODS
func (t *Truck) Start() {
	fmt.Println("Truck started\n")
}
func (t *Truck) Stop() {
	fmt.Println("Truck stopped\n")
}

func (t *Truck) Details() {
	fmt.Printf("-------- Truck Details --------\n")
	fmt.Printf("numOfWheels: %d\n", t.numOfWheels)
	fmt.Printf("horsePower: %d Cv\n", t.horsePower)
	fmt.Printf("numOfDoors: %d\n", t.numOfDoors)
	fmt.Printf("loadCapacity: %d kg\n", t.loadCapacity)
	fmt.Printf("-------------------------------\n")
}

// MOTORCYCLE METHODS
func (m *Motorcycle) Start() {
	fmt.Println("Motorcycle started\n")
}
func (m *Motorcycle) Stop() {
	fmt.Println("Motorcycle stopped\n")
}
func (m *Motorcycle) Details() {
	fmt.Printf("-------- Motorcycle Details ---\n")
	fmt.Printf("numOfWheels: %d\n", m.numOfWheels)
	fmt.Printf("horsePower: %d Cv\n", m.horsePower)
	fmt.Printf("-------------------------------\n")
}

func main() {

	//Generating Instances
	//To maximize the benefits of interfaces we use an array of vehicles (Polymorphism)
	vehicles := []Vehicle{
		&Car{4, 100, 5},
		&Truck{8, 530, 3, 80000},
		&Motorcycle{2, 140},
	}
	for _, v := range vehicles {
		v.Start()
		v.Details()
		v.Stop()
	}
}
