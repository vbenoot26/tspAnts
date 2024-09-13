package main

import (
	"fmt"
	"math"
	"math/rand"
)

const (
	GRAPH_SIZE  = 5
	ANTS_AMOUNT = GRAPH_SIZE
	ITERATIONS  = 10
	ALPHA       = 0.9
	BETA        = 1.5
)

var (
	graph      Graph
	pheromones Graph
	antPaths   [ANTS_AMOUNT][GRAPH_SIZE]uint
)

type Graph [GRAPH_SIZE][GRAPH_SIZE]float64

func main() {
	initGraph()
	initAnts()

	doRandomTours()

	for _, path := range antPaths {
		fmt.Println(path)
	}
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
		antPaths[i][0] = uint(i)
	}
}

func doRandomTours() {
	for i := 1; i < GRAPH_SIZE; i++ {
		for ant := 0; ant < ANTS_AMOUNT; ant++ {
			pickNextNode(ant, i)
		}
	}
}

func calculate() {
	for ant := range antPaths {
		for step := 1; step < GRAPH_SIZE; step++ {
			distribution := getProbabilityDistribution(ant, step)
			antPaths[ant][step] = uint(pickIndex(distribution[:]))
		}
	}

	updatePheromone()
}

func pickNextNode(ant int, step int) {
	for true {
		maybenode := uint(rand.Intn(GRAPH_SIZE))
		alreadyVisited := false
		for nodeIdx := 0; nodeIdx < step; nodeIdx++ {
			if antPaths[ant][nodeIdx] == maybenode {
				alreadyVisited = true
				break
			}
		}

		if alreadyVisited {
			continue
		}

		antPaths[ant][step] = maybenode
		break
	}
}

func getProbabilityDistribution(ant int, step int) [GRAPH_SIZE]float64 {
	var numerators [GRAPH_SIZE]float64
	possibilityUniverse := 0.0
	lastNode := antPaths[ant][step-1]

	for node := 0; node < GRAPH_SIZE; node++ {
		if contains(antPaths[ant][:step], uint(node)) {
			numerators[node] = 0
		} else {
			numerators[node] = math.Pow(pheromones[lastNode][node], ALPHA) * math.Pow(graph[lastNode][node], BETA)
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
	// TODO: impelement
}

func contains(targetSlice []uint, value uint) bool {
	for _, element := range targetSlice {
		if element == value {
			return true
		}
	}

	return false
}
