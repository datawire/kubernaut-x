package util

import (
	"math/rand"
)

type SillyNameGenerator struct {
	Separator          string
	WithColor          bool
	WithAdjective      bool
	WithNoun           bool
	RandomSuffix       string
	RandomSuffixLength int

	random *rand.Rand
}

func NewSillyNameGeneratorUsingCustomSeed(seed int64) StringGenerator {
	return &SillyNameGenerator{random: rand.New(rand.NewSource(seed))}
}

func (g *SillyNameGenerator) Generate() string {
	return ""
}
