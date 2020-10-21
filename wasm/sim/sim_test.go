package sim

import (
	"fmt"
	"math"
	"reflect"
	"strings"
	"testing"

	"github.com/quasilyte/gophers-and-dragons/game"
)

func TestCalculateHealed(t *testing.T) {
	const maxHP = 15

	tests := []struct {
		roll    int
		current int
		want    int
	}{
		{roll: 0, current: 10, want: 0},
		{roll: 1, current: 10, want: 1},
		{roll: 5, current: 10, want: 5},
		{roll: 6, current: 10, want: 5},
		{roll: 100, current: 10, want: 5},
		{roll: 5, current: 11, want: 4},
		{roll: 100, current: 1, want: 14},
		{roll: 1, current: maxHP, want: 0},
	}

	for _, test := range tests {
		have := calculateHealed(test.roll, test.current, maxHP)
		if have != test.want {
			t.Errorf("roll=%d current=%d max=%d:\nhave: %d\nwant: %d",
				test.roll, test.current, maxHP, have, test.want)
		}
	}
}

func TestRun(t *testing.T) {
	tests := []struct {
		seed int64
		fn   func(game.State) game.CardType
	}{
		{1, func(state game.State) game.CardType { return game.CardRetreat }},
		{2, func(state game.State) game.CardType { return game.CardAttack }},
	}

	for _, test := range tests {
		config := &Config{
			AvatarHP: 40,
			AvatarMP: 20,
			Rounds:   10,
			Seed:     test.seed,
		}
		firstResult := Run(config, test.fn)
		secondResult := Run(config, test.fn)

		if !reflect.DeepEqual(firstResult, secondResult) {
			builder := strings.Builder{}
			builder.WriteString(fmt.Sprintf("seed=%d different results\n", test.seed))
			firstLen, secondLen := len(firstResult), len(secondResult)
			for i := 0; i < int(math.Min(float64(firstLen), float64(secondLen))); i++ {
				builder.WriteString(fmt.Sprintf("%v\t\t%v\n", firstResult[i], secondResult[i]))
			}
			t.Error(builder.String())
		}
	}
}
