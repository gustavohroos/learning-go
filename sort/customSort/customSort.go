package main

import (
	"fmt"
	"sort"
)

type brand struct {
	name string
}

type car struct {
	name        string
	brand       brand
	power       int
	consumption float64
}

type sortCarByPower []car

func (s sortCarByPower) Len() int           { return len(s) }
func (s sortCarByPower) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s sortCarByPower) Less(i, j int) bool { return s[i].power < s[j].power }

type sortCarByConsumption []car

func (s sortCarByConsumption) Len() int           { return len(s) }
func (s sortCarByConsumption) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s sortCarByConsumption) Less(i, j int) bool { return s[i].consumption < s[j].consumption }

type sortCarByName []car

func (s sortCarByName) Len() int           { return len(s) }
func (s sortCarByName) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s sortCarByName) Less(i, j int) bool { return s[i].name < s[j].name }

type sortCarByBrand []car

func (s sortCarByBrand) Len() int           { return len(s) }
func (s sortCarByBrand) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s sortCarByBrand) Less(i, j int) bool { return s[i].brand.name < s[j].brand.name }

func main() {
	brands := []brand{
		{"Toyota"},
		{"Ford"},
		{"Chevrolet"},
		{"Nissan"},
		{"Volkswagen"},
		{"Honda"},
		{"Hyundai"},
	}

	cars := []car{
		{"Corolla", brands[0], 140, 14.5},
		{"Fiesta", brands[1], 120, 12.5},
		{"Camaro", brands[2], 300, 20.5},
		{"Sentra", brands[3], 150, 15.5},
		{"Gol", brands[4], 100, 10.5},
		{"Civic", brands[5], 130, 13.5},
		{"HB20", brands[6], 110, 11.5},
	}

	sort.Sort(sortCarByPower(cars))
	fmt.Println("Cars sorted by power:")
	for _, c := range cars {
		println(c.name, c.power)
	}

	sort.Sort(sortCarByConsumption(cars))
	fmt.Println("Cars sorted by consumption:")
	for _, c := range cars {
		println(c.name, c.consumption)
	}

	sort.Sort(sortCarByName(cars))
	fmt.Println("Cars sorted by name:")
	for _, c := range cars {
		println(c.name)
	}

	sort.Sort(sortCarByBrand(cars))
	fmt.Println("Cars sorted by brand:")
	for _, c := range cars {
		println(c.name, c.brand.name)
	}

}
