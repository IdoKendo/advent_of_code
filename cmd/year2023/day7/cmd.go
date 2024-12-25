package day7

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day7",
	Short: "day7",
	Long:  "day7",
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd.Parent().Name(), cmd.Name())

	},
}

func execute(parent, command string) {
	content, err := os.ReadFile(fmt.Sprintf("cmd/year%s/%s/input1.txt", parent, command))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	result := Part1(string(content))
	fmt.Println("Part 1 result: ", result)

	content, err = os.ReadFile(fmt.Sprintf("cmd/year%s/%s/input2.txt", parent, command))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	result = Part2(string(content))
	fmt.Println("Part 2 result: ", result)
}

type Power int

const (
	Five      Power = 6
	Four      Power = 5
	FullHouse Power = 4
	Three     Power = 3
	TwoPair   Power = 2
	Pair      Power = 1
	High      Power = 0
)

var Card = map[byte]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
}

var CardWithJoker = map[byte]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
	'J': 1,
}

type Hand struct {
	Raw   string
	Bid   int
	Power Power
}

func calculatePower(raw string, jokerSymbol string) Power {
	cards := strings.Split(raw, "")
	hand := make(map[string]int, len(cards))
	for _, card := range cards {
		hand[card]++
	}
	jokerCount := hand[jokerSymbol]
	if jokerCount == 5 {
		return Five
	}
	maxPower := High
	for symbol, count := range hand {
		if symbol == jokerSymbol {
			continue
		}
		switch count {
		case 5:
			maxPower = Five
		case 4:
			if jokerCount >= 1 {
				maxPower = Five
			} else if maxPower < Four {
				maxPower = Four
			}
		case 3:
			if jokerCount >= 2 {
				maxPower = Five
			} else if jokerCount == 1 && maxPower < Four {
				maxPower = Four
			} else if maxPower == Pair && maxPower < FullHouse {
				maxPower = FullHouse
			} else if maxPower < Three {
				maxPower = Three
			}
		case 2:
			if jokerCount >= 3 {
				maxPower = Five
			} else if jokerCount == 2 && maxPower < Four {
				maxPower = Four
			} else if maxPower == Three && maxPower < FullHouse {
				maxPower = FullHouse
			} else if jokerCount == 1 && maxPower < Three {
				maxPower = Three
			} else if maxPower == Pair && maxPower < TwoPair {
				maxPower = TwoPair
			} else if maxPower < Pair {
				maxPower = Pair
			}
		case 1:
			if jokerCount >= 4 {
				maxPower = Five
			} else if jokerCount == 3 && maxPower < Four {
				maxPower = Four
			} else if jokerCount == 2 && maxPower < Three {
				maxPower = Three
			} else if jokerCount == 1 && maxPower < Pair {
				maxPower = Pair
			}
		}
	}
	return maxPower
}

func parseHands(input, jokerSign string) []Hand {
	lines := strings.Split(input, "\n")
	rawHands := lines[:len(lines)-1]
	hands := make([]Hand, len(rawHands))
	for i, h := range rawHands {
		parts := strings.Split(h, " ")
		bid, _ := strconv.Atoi(parts[1])
		hands[i] = Hand{
			Raw:   parts[0],
			Bid:   bid,
			Power: calculatePower(parts[0], jokerSign),
		}
	}
	return hands
}

func Part1(input string) int {
	hands := parseHands(input, "THERE_IS_NO_JOKER")
	sort.Slice(hands, func(i, j int) bool {
		if hands[i].Power == hands[j].Power {
			for x := range 5 {
				if hands[i].Raw[x] == hands[j].Raw[x] {
					continue
				}
				return Card[hands[i].Raw[x]] < Card[hands[j].Raw[x]]
			}
		}
		return hands[i].Power < hands[j].Power
	})
	result := 0
	for i, hand := range hands {
		result += hand.Bid * (i + 1)
	}
	return result
}

func Part2(input string) int {
	hands := parseHands(input, "J")
	sort.Slice(hands, func(i, j int) bool {
		if hands[i].Power == hands[j].Power {
			for x := range 5 {
				if hands[i].Raw[x] == hands[j].Raw[x] {
					continue
				}
				return CardWithJoker[hands[i].Raw[x]] < CardWithJoker[hands[j].Raw[x]]
			}
		}
		return hands[i].Power < hands[j].Power
	})
	result := 0
	for i, hand := range hands {
		result += hand.Bid * (i + 1)
	}
	return result
}
