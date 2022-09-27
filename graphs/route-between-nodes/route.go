package main

func IsRouteExists(graph map[int][]int, from, to int) bool {
	queue := NewQueue()
	queue.Enqueue(from)
	visited := map[int]bool{}

	for !queue.IsEmpty() {
		node, _ := queue.Dequeue()
		if _, ok := visited[node]; !ok {
			visited[node] = true
			if neighbours, ok := graph[node]; ok {
				for _, neighbour := range neighbours {
					if neighbour == to {
						return true
					}
					queue.Enqueue(neighbour)
				}
			}
		}
	}
	return false
}
