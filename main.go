package main

import (
	"fmt"
	"math/rand"
)

const (
	GRAPH_SIZE  = 5
	ANTS_AMOUNT = GRAPH_SIZE
	ITERATIONS  = 10
)

var (
	graph    [GRAPH_SIZE][GRAPH_SIZE]float32
	antPaths [ANTS_AMOUNT][GRAPH_SIZE]uint
)

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
			edgeweigth := rand.Float32()
			graph[i][j] = edgeweigth
			graph[j][i] = edgeweigth
		}
	}
}

// Puts one ant one every node since amount of ants == graph size
func initAnts() {
	for i := 0; i < ANTS_AMOUNT; i++ {
		antPaths[i][0] = uint(i)
	}
}

func doRandomTours() {
	for i := 0; i < GRAPH_SIZE; i++ {
		for ant := 0; ant < ANTS_AMOUNT; ant++ {
			pickNextCity(ant, i)
		}
	}
}

func pickNextCity(ant int, step int) {
	for true {
		maybecity := uint(rand.Intn(GRAPH_SIZE))
		alreadyVisited := false
		for cityIdx := 0; cityIdx < step; cityIdx++ {
			if antPaths[ant][cityIdx] == maybecity {
				alreadyVisited = true
				break
			}
		}

		if alreadyVisited {
			continue
		}

		antPaths[ant][step] = maybecity
		break
	}
}
