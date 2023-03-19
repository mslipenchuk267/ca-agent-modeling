package main

import (
	"fmt"

	agent "github.com/ca-agent-modeling/src/pkg/agent"
)

func main() {
	my2dAgent := &agent.Agent{Coords: []int{1,2}, Activation: 1, Suggestibility: 0.5}
	my3dAgent := &agent.Agent{Coords: []int{1,2,3}, Activation: 0, Suggestibility: 0.1}
	
	fmt.Println("Printing Agent Info on 2d and 3d coords")
	my2dAgent.ToString()
	my3dAgent.ToString()
}

