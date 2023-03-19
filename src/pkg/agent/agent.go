package agent

import (
	"fmt"
	"strings"
)

type Agent struct {
	Coords []int
}

func (agent Agent) ToString() {
	strCoords := strings.Trim(fmt.Sprint(agent.Coords), "[]")
	agentStr := "Agent Coords: " + strCoords
	fmt.Println(agentStr)
}

