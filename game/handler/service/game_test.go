package service

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestGame_GameRandom(t *testing.T) {
	random1(10)
}

func random1(level int) int {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	dice1 := rand.Intn(6) + 1
	dice2 := rand.Intn(6) + 1

	fmt.Println(dice1)
	fmt.Println(dice2)

	base := (dice1 + dice2) * level
	score := 0
	if dice1 != dice2 {
		score = base
	} else {
		switch dice1 {
		case 1:
			score = base * 2
		case 2:
			score = base * 4
		case 3:
			score = base * 6
		case 4:
			score = base * 8
		case 5:
			score = base * 10
		case 6:
			score = base * 12
		default:
			break
		}
	}
	fmt.Println(score)
	return score
}
