package main

import (
	"math/rand"
	"time"
)

type ActiveCondition func(x int, y int, config *Config) bool

var generators = map[int]ActiveCondition{
	0: activeCondition1,
	1: activeCondition2,
	2: activeCondition3,
}

func activeCondition1(x int, y int, config *Config) bool {
	return x > 0 && x < config.Width-1 && y%3 == 0 && y > 3
}
func activeCondition2(x int, y int, config *Config) bool {
	return x > 3 && x < config.Width-3 && y%5 == 0 && y > 3
}
func activeCondition3(x int, y int, config *Config) bool {
	return x < 5 && x < config.Width-2 && y%2 != 0 && y > 3
}

func GetGenerators(config *Config) []ActiveCondition {
	var gens []ActiveCondition
	for i := 0; i < config.Width; i++ {
		for j := 0; j < config.Height; j++ {
			gens = append(gens, toCondition(j, i, config))
		}
	}
	return gens
}

func toCondition(x int, y int, config *Config) ActiveCondition {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(11) + 1
	return func(x, y int, config *Config) bool {
		return x > 0 && x < config.Width && y%r == 0 && y > 0
	}
}
