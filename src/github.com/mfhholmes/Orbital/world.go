package main

import()

type World_Builder interface{
	Population() int
	SetPopulation(int) error
}

type world_values struct{
	population int
}

type World struct{
	wv *world_values
}

func NewWorld (wv *world_values) (*World,error){
	result := new(World)
	// might want to test the world values for consistency at some point
	if wv!=nil{
		result.wv = wv
	}	else {
		result.wv = new(world_values)
	}
	return result, nil
}
func (world World) Population() int {
	return world.wv.population
}
func (world *World) SetPopulation(newPop int) error {
	world.wv.population = newPop
	return nil
}