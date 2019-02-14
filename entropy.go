package main

import (
	"math/rand"
	"time"
)

type WorldWithEntropy struct {
	state         []rune
	origins       []rune
	indexes       []int
	it            int
	teleologyMode bool
	chaosSteps    uint
	rand          *rand.Rand
}

func NewWorld(origins string) *WorldWithEntropy {
	var now = time.Now().UnixNano()
	var rnd = rand.New(rand.NewSource(now))
	var state = []rune(origins)
	var indexes = make([]int, len(state))

	for i := range state {
		indexes[i] = i
	}
	rnd.Shuffle(len(state), func(i, j int) {
		indexes[i], indexes[j] = indexes[j], indexes[i]
	})

	return &WorldWithEntropy{
		state:         state,
		origins:       []rune(origins),
		indexes:       indexes,
		teleologyMode: false,
		chaosSteps:    0,
		rand:          rnd,
	}
}

func (world *WorldWithEntropy) maxOfEntropy() bool {
	return uint(len(world.state)) == world.chaosSteps
}

func (world *WorldWithEntropy) minOfEntropy() bool {
	return world.chaosSteps == 0
}

func (world *WorldWithEntropy) Step() {
	switch {
	case !world.teleologyMode && world.maxOfEntropy():
		world.teleologyMode = true
	case world.teleologyMode && !world.minOfEntropy():
		world.teleologicStep()
		world.chaosSteps--
	case world.teleologyMode && world.minOfEntropy():
		world.teleologyMode = false
	case !world.teleologyMode && !world.maxOfEntropy():
		world.chaoticStep()
		world.chaosSteps++
	default:
		panic("Broken World: stopping Universe execution")
	}
}

func (world *WorldWithEntropy) chaoticStep() {
	var rand = world.rand
	var alpha = alphabet()
	var mutationIndex = rand.Intn(len(alpha))
	var selectedIndex = world.getNextIndex()
	if world.state[selectedIndex] != '\n' {
		world.state[selectedIndex] = alpha[mutationIndex]
	}
}

func (world *WorldWithEntropy) teleologicStep() {
	var selectedIndex = world.getNextIndex()

	world.state[selectedIndex] = world.origins[selectedIndex]
}

func (world *WorldWithEntropy) getNextIndex() int {
	var rand = world.rand
	var indexes = world.indexes
	var iLen = len(indexes)

	if world.it == iLen {
		rand.Shuffle(iLen, func(i, j int) {
			indexes[i], indexes[j] = indexes[j], indexes[i]
		})
		world.it = 0
	}
	var ind = indexes[world.it]
	world.it++
	return ind
}

func alphabet() []rune {
	return []rune(`abcdefghijklmnopqrstuvwxyz"!@#$%^&*()_+{}?><,./|\\`)
}
