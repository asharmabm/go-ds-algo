package main

func main() {
	// Graph
	//        [2]--------[0]
	//        / \
	//	     /   \
	//      /     \
	//    [1]-----[3]

	// Edge List
	graph := [][]int{{0, 2}, {2, 3}, {2, 1}, {1, 3}}

	// Adjacent List
	graph = [][]int{{2}, {2, 3}, {0, 1, 3}, {1, 2}}

	// Adjacent Metrics
	graph = [][]int{
		{0, 0, 1, 0},
		{0, 0, 1, 1},
		{1, 1, 0, 3},
		{0, 1, 1, 0},
	}

	_ = graph

}
