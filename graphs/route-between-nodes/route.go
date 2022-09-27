package main

func IsRouteExists(graph map[int][]int, from, to int) bool {
	queue := make([]int, 0)
	queue = append(queue, from)
	visited := map[int]bool{}

	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]

		if _, ok := visited[node]; !ok {
			visited[node] = true
			if neighbours, ok := graph[node]; ok {
				for _, neighbour := range neighbours {
					if neighbour == to {
						return true
					}
					queue = append(queue, neighbour)
				}
			}
		}
	}
	return false
}
