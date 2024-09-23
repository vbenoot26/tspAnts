package main

// Represents the pheromones + paths an ant took after an iteratation.
type state struct {
	pheromones Graph
	antPaths   [ANTS_AMOUNT][GRAPH_SIZE]int
}

var iterationStates []state

func saveState(pheromones Graph, antPaths [ANTS_AMOUNT][GRAPH_SIZE]int) {
	iterationStates = append(iterationStates, state{
		pheromones, antPaths,
	})
}
