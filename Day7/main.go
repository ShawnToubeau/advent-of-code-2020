package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
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

func populateBagGraph(g *Graph, bagRules []string, direction string, containmentMap map[string]int) {
	for _, rule := range bagRules {
		// split to get the bag specs from the rule
		ruleSpecs := strings.Split(rule, ", ")
		v1Name := ""
		v2Name := ""
		childAmount := 0

		for i, spec := range ruleSpecs {
			details := strings.Split(spec, " ")

			// starting bag
			if i == 0 {
				v1Name = fmt.Sprintf("%v-%v", string(details[0]), string(details[1]))
			} else { // following bags
				v2Name = fmt.Sprintf("%v-%v", string(details[1]), string(details[2]))
				childAmount, _ = strconv.Atoi(details[0])
			}

			if i != 0 {
				// create edges if the rule contains 2 or more bag specs
				if v1Name != "" && v2Name != "" && v2Name != "other-bags." {
					g.AddEdge(v1Name, v2Name, direction)

					mapKey := fmt.Sprintf("%v-%v", v1Name, v2Name)
					containmentMap[mapKey] = childAmount
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

// Part 1
func findUniqueParents(g *Graph, vertexKey string, sources *[]string) {
	for key, value := range g.Vertices[vertexKey].Vertices {
		// append key if our sources array doesn't contain it already
		if !hasElem(sources, key) {
			*sources = append(*sources, key)
		}

		// while the current vertex has parent vertexes, recursively call findUniqueParents
		if len(value.Vertices) > 0 {
			findUniqueParents(g, key, sources)
		}
	}
}

// Part 2
func sumEnclosedChildren(g *Graph, startVertex string, quantityMap map[string]int) {
	var stack []string
	// append first vertex
	stack = append(stack, startVertex)
	// starting bag quantifier
	startQuantifier := 1
	// for every level of children, we'll need to know how many parent bags there are so we can
	// accurately keep track of the quantities
	var quantifiers []int
	quantifiers = append(quantifiers, startQuantifier)
	// our result
	res := 0

	// loop while our stack isn't empty
	for len(stack) > 0 {
		// index of the top element
		n := len(stack) - 1
		// pop off top element
		curr := stack[n]
		stack = stack[:n]
		// pop off the top bag quantifier
		currQuantifiers := quantifiers[n]
		quantifiers = quantifiers[:n]
		// push every child of the current element onto the stack
		for _, vertex := range g.Vertices[curr].Vertices {
			// quantity map accessor, `parent-bag`-`child-bag` = # of children within the parent
			mapKey := fmt.Sprintf("%v-%v", curr, vertex.Key)
			// using the known quantity and the quantifier we know exactly how many multiples of the children there are
			res = res + currQuantifiers*quantityMap[mapKey]

			stack = append(stack, vertex.Key)
			// record this iterations children quantities for future bag children
			quantifiers = append(quantifiers, currQuantifiers*quantityMap[mapKey])
		}
	}

	fmt.Printf("Num bags %v\n", res)
}

func main() {
	upwardGraph := NewGraph()
	downwardGraph := NewGraph()
	containmentMap := make(map[string]int)
	var uniqueSources []string

	bagRules := parseBags("./bags.txt")

	populateBagGraph(upwardGraph, bagRules, UP, make(map[string]int))

	findUniqueParents(upwardGraph, "shiny-gold", &uniqueSources)

	fmt.Printf("Sources %v\n", len(uniqueSources))

	populateBagGraph(downwardGraph, bagRules, DOWN, containmentMap)

	sumEnclosedChildren(downwardGraph, "shiny-gold", containmentMap)
}
