package agent

import (
	"fmt"
	"strings"
)

type Agent struct {
	Coords []int
	Activation int
	Money int
	Happiness int
	CommonSense float32
	BookSmarts float32
	Suggestibility float32
}

func (agent Agent) ToString() {
	strCoords := strings.Trim(fmt.Sprint(agent.Coords), "[]")
	fmt.Printf("Agent Info\n\tCoords: %s\n\tActivation: %d\n\tMoney: %d\n\tCommon-Sense: %f\n\tBook-Smarts: %f\n\tSuggestibility: %f\n", strCoords, agent.Activation, agent.Money, agent.CommonSense, agent.BookSmarts, agent.Suggestibility)
}

