package main

import (
	"fmt"
	"math/rand"
)

const (
	GRAPH_SIZE           = 5
	ANTS_AMOUNT          = GRAPH_SIZE
	ITERATIONS           = 100
	HEURISTIC_IMPORTANCE = 0.9
	PHEROMONE_IMPORTANCE = 1.5
	DECAY                = 0.9
)

var (
	cities Graph
)

type Graph [GRAPH_SIZE][GRAPH_SIZE]float64

func main() {
	initGraph()
	printGraph()
	tour := findTour(cities)
	fmt.Println(tour)
	fmt.Println(bestCost)
}

// Initialises a random graph.
func initGraph() {
	// the graph itself
	for i := 0; i < GRAPH_SIZE; i++ {
		for j := 0; j <= i; j++ {
			if i == j {
				cities[i][j] = 0
			} else {
				edgeweigth := rand.Float64()
				cities[i][j] = edgeweigth
				cities[j][i] = edgeweigth
			}
		}
	}
	// the pheromones
	for i := 0; i < GRAPH_SIZE; i++ {
		for j := 0; j < GRAPH_SIZE; j++ {
			pheromones[i][j] = 1
		}
	}
}

// DEBUGTOOLS ------------------------------------
func printGraph() {
	for i := 0; i < GRAPH_SIZE; i++ {
		fmt.Println(cities[i])
	}
}
