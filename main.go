package main

import (
	"fmt"
	"math"
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
	graph      Graph
	pheromones Graph
	antPaths   [ANTS_AMOUNT][GRAPH_SIZE]int
	bestTour   [GRAPH_SIZE]int
	bestCost   = 1000000000.0
)

type Graph [GRAPH_SIZE][GRAPH_SIZE]float64

func main() {
	initGraph()
	initAnts()

	for i := 0; i < ITERATIONS; i++ {
		calculate()

		// See if there is a better path
		for _, path := range antPaths {
			pathCost := getPathLength(path[:])
			if pathCost < bestCost {
				bestCost = pathCost
				for i := range path {
					bestTour[i] = path[i]
				}
			}
		}
	}

	fmt.Println(bestTour)
	fmt.Println(bestCost)
}

// Initialises a random graph.
func initGraph() {
	for i := 0; i < GRAPH_SIZE; i++ {
		for j := 0; j <= i; j++ {
			edgeweigth := rand.Float64()
			graph[i][j] = edgeweigth
			graph[j][i] = edgeweigth
		}
	}
}

// Puts one ant every node since amount of ants == graph size
func initAnts() {
	for i := 0; i < ANTS_AMOUNT; i++ {
		antPaths[i][0] = i
	}
}

func calculate() {
	for ant := range antPaths {
		for step := 1; step < GRAPH_SIZE; step++ {
			distribution := getProbabilityDistribution(ant, step)
			antPaths[ant][step] = pickIndex(distribution[:])
		}
	}

	updatePheromone()
}

func getProbabilityDistribution(ant int, step int) [GRAPH_SIZE]float64 {
	var numerators [GRAPH_SIZE]float64
	possibilityUniverse := 0.0
	lastNode := antPaths[ant][step-1]

	for node := 0; node < GRAPH_SIZE; node++ {
		if contains(antPaths[ant][:step], node) {
			numerators[node] = 0
		} else {
			numerators[node] = math.Pow(pheromones[lastNode][node], HEURISTIC_IMPORTANCE) * math.Pow(graph[lastNode][node], PHEROMONE_IMPORTANCE)
			possibilityUniverse += numerators[node]
		}
	}

	var probabilities [GRAPH_SIZE]float64
	for node := range numerators {
		probabilities[node] = numerators[node] / possibilityUniverse
	}
	return probabilities
}

func pickIndex(probablities []float64) int {
	cum := 0.0
	randSelection := rand.Float64()

	for i, prob := range probablities {
		cum += prob
		if cum > randSelection {
			return i
		}
	}

	return len(probablities) - 1
}

func updatePheromone() {
	// Decay
	for i, row := range pheromones {
		for j := range row {
			pheromones[i][j] *= 1 - DECAY
		}
	}

	// Delta
	for _, path := range antPaths {
		// calculate tour length
		tourLength := getPathLength(path[:])

		for idx := range path {
			pheromones[path[idx]][path[(idx+1)%len(path)]] += 1 / tourLength
		}
	}
}

func getPathLength(path []int) float64 {
	tourLength := 0.0
	for idx := range path {
		tourLength += graph[path[idx]][path[(idx+1)%len(path)]]
	}

	return tourLength
}

func contains(targetSlice []int, value int) bool {
	for _, element := range targetSlice {
		if element == value {
			return true
		}
	}

	return false
}
