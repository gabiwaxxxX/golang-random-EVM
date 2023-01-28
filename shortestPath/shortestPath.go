package shortestPath

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
)

func QueryUniswapLps() {
	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")

    if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }

    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(responseData))
}





//function Dijkstra Algorithm in Golang 
// Find Shortest Distance ie Shortest Path between two nodes in weighted graph
func FindShortestPath(graph map[string]map[string]int, start, end string) ([]string, int) {
	visited := make(map[string]bool)
	dist := make(map[string]int)
	prev := make(map[string]string)
	for k := range graph {
		dist[k] = 999999
	}
	dist[start] = 0
	prev[start] = ""
	for {
		min := 999999
		for k := range dist {
			if !visited[k] && dist[k] < min {
				min = dist[k]
				minKey := k
			}
		}
		if min == 999999 {
			break
		}
		visited[minKey] = true
		for k, v := range graph[minKey] {
			if !visited[k] && v+dist[minKey] < dist[k] {
				dist[k] = v + dist[minKey]
				prev[k] = minKey
			}
		}
	}
	path := make([]string, 0)
	for k := end; k != start; k = prev[k] {
		path = append(path, k)
	}
	path = append(path, start)
	reverse(path)
	return path, dist[end]
}

func reverse(s []string) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

//function Dijkstra Algorithm in Golang 
// Find Shortest Distance ie Shortest Path between two nodes in graph not weighted
func FindShortestPathNotWeighted(graph map[string]map[string]bool, start, end string) ([]string, int) {

