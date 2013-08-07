package main

import (
	"testing"
)

func TestWorldPopulationSet(t *testing.T){
	// create a world
	world,err := NewWorld(nil)
	if err != nil{
		t.Fatal("Failed to create World")
	}
	// set the population
	world.SetPopulation(5)
	// get the population
	if world.Population() != 5{
		t.Fatal("population getter didn't return the correct population")
	}
}