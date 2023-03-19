package agent

import (
	"fmt"
	"strings"
)

type Agent struct {
	Coords []int
	Activation int
	Suggestibility float32
}

func (agent Agent) ToString() {
	strCoords := strings.Trim(fmt.Sprint(agent.Coords), "[]")
	fmt.Printf("Agent Info\n\tCoords: %s\n\tActivation: %d\n\tSuggestibility: %f\n", strCoords, agent.Activation, agent.Suggestibility)
}

