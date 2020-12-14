package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	UP   = "up"
	DOWN = "down"
)

type Vertex struct {
	// bag type & color
	Key      string
	Vertices map[string]*Vertex
}

func NewVertex(key string) *Vertex {
	return &Vertex{
		Key:      key,
		Vertices: map[string]*Vertex{},
	}
}

type Graph struct {
	Vertices map[string]*Vertex
}

func NewGraph() *Graph {
	return &Graph{
		Vertices: map[string]*Vertex{},
	}
}

func (g *Graph) AddVertex(key string) {
	v := NewVertex(key)
	g.Vertices[key] = v
}

func (g *Graph) AddEdge(k1, k2 string, direction string) {
	v1 := g.Vertices[k1]
	v2 := g.Vertices[k2]

	// create vertex 1 if it doesn't exist
	if v1 == nil {
		g.AddVertex(k1)
		v1 = g.Vertices[k1]
	}
	// create vertex 2 if it doesn't exist
	if v2 == nil {
		g.AddVertex(k2)
		v2 = g.Vertices[k2]
	}

	if direction == DOWN {
		// add directional edge from vertex 1 to vertex 2
		if _, ok := v1.Vertices[v2.Key]; !ok {
			v1.Vertices[v2.Key] = v2
		}
	}

	if direction == UP {
		// add directional edge from vertex 2 to vertex 1
		if _, ok := v2.Vertices[v1.Key]; !ok {
			v2.Vertices[v1.Key] = v1
		}
	}

	// update our vertices
	g.Vertices[v1.Key] = v1
	g.Vertices[v2.Key] = v2
}

func parseBags(file string) []string {
	text, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("Error parsing file: %v\n", err)
	}

	// replace all occurrences of "contain" with a comma for easier data manipulation
	str := strings.ReplaceAll(string(text), " contain", ",")

	bagRules := strings.Split(str, "\n")

	return bagRules
}

func populateBagGraph(g *Graph, bagRules []string, direction string) {
	for _, rule := range bagRules {
		// split to get the bag specs from the rule
		ruleSpecs := strings.Split(rule, ", ")
		v1Name := ""
		v2Name := ""

		for i, spec := range ruleSpecs {
			details := strings.Split(spec, " ")

			// starting bag
			if i == 0 {
				v1Name = fmt.Sprintf("%v-%v", string(details[0]), string(details[1]))
			} else { // following bags
				v2Name = fmt.Sprintf("%v-%v", string(details[1]), string(details[2]))
			}

			if i != 0 {
				// create edges if the rule contains 2 or more bag specs
				if v1Name != "" && v2Name != "" && v2Name != "other-bags." {
					g.AddEdge(v1Name, v2Name, direction)
				} else { // else just add a vertex for the first bag
					g.AddVertex(v1Name)
				}
			}
		}
	}
}

// checks whether an array contains a element
func hasElem(arr *[]string, elemToFind string) bool {
	for _, elem := range *arr {
		if elem == elemToFind {
			return true
		}
	}
	return false
}

// checks whether an array contains a element
func hasElemByVal(arr []string, elemToFind string) bool {
	for _, elem := range arr {
		if elem == elemToFind {
			return true
		}
	}
	return false
}

func traverseUp(g *Graph, vertexKey string, sources *[]string) {
	for key, value := range g.Vertices[vertexKey].Vertices {
		// append key if our sources array doesn't contain it already
		if !hasElem(sources, key) {
			*sources = append(*sources, key)
		}

		// while the current vertex has parent vertexes, recursively call traverseUp
		if len(value.Vertices) > 0 {
			traverseUp(g, key, sources)
		}
	}
}

//func traverseDown (downwardGraph *Graph, upwardGraph *Graph, vertexKey string, sources *[]string) {
//	leaf := BFS(downwardGraph, vertexKey)
//
//	if leaf != "" {
//
//	}
//
//	for key, value := range g.Vertices[vertexKey].Vertices {
//		// append key if our sources array doesn't contain it already
//		if !hasElem(sources, key) {
//			*sources = append(*sources, key)
//		}
//
//		// while the current vertex has parent vertexes, recursively call traverseUp
//		if len(value.Vertices) == 0 {
//			traverseUp(g, key, sources)
//		}
//	}
//}

func BFS(g *Graph, startVertex string) {
	var visited []string
	var queue []string
	prev := make(map[string]string)
	queue = append(queue, startVertex)
	visited = append(visited, startVertex)
	prev[startVertex] = "stop"
	//var leaf string

	for len(queue) > 0 {
		// pop off first element
		curr := queue[0]
		queue = queue[1:]

		for key, vertex := range g.Vertices[curr].Vertices {
			prev[key] = curr

			if len(vertex.Vertices) == 0 {
				fmt.Printf("Found leaf: %v\n", key)
				//leaf = key

				//queue = nil
				//break
			}

			if !hasElemByVal(visited, vertex.Key) {
				queue = append(queue, key)
				visited = append(visited, key)
			}
		}
	}

	//for leaf != "stop" {
	fmt.Printf("%v\n", prev)
	//	leaf = prev[leaf]
	//}

	//return ""
}

func main() {
	//upwardGraph := NewGraph()
	downwardGraph := NewGraph()
	//var uniqueSources []string

	bagRules := parseBags("./bags.txt")

	//populateBagGraph(upwardGraph, bagRules, UP)
	//
	//traverseUp(upwardGraph, "shiny-gold", &uniqueSources)
	//
	//fmt.Printf("Sources %v\n", len(uniqueSources))

	populateBagGraph(downwardGraph, bagRules, DOWN)

	BFS(downwardGraph, "shiny-gold")
}
