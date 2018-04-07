package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Bot struct {
	V1, V2                  int
	LowDst, HighDst         int
	LowDstType, HighDstType DstType
	JobDone                 bool
}

type DstType bool

const (
	BotDst    DstType = false
	OutputDst DstType = true
)

func (b *Bot) Take(n int) {
	if b.V1 != 0 {
		b.V2 = n
	} else {
		b.V1 = n
	}
}

func (b *Bot) GiveAwayLow() int {
	if b.V2 > b.V1 {
		v := b.V1
		b.V1 = 0
		return v
	} else {
		v := b.V2
		b.V2 = 0
		return v
	}
}

func (b *Bot) GiveAwayHigh() int {
	if b.V2 < b.V1 {
		v := b.V1
		b.V1 = 0
		return v
	} else {
		v := b.V2
		b.V2 = 0
		return v
	}
}

func main() {
	input := loadInput()

	bots := make(map[int]Bot)

	initialState(input, bots)

	todo, done := len(bots), 0

	outputs := make(map[int]int)

	for done != todo {
		for botID, _ := range bots {
			bot := bots[botID]
			if bot.JobDone || bot.V1 == 0 || bot.V2 == 0 {
				continue
			}

			if (bot.V1 == 17 && bot.V2 == 61) || (bot.V1 == 61 && bot.V2 == 17) {
				fmt.Println("Part one:", botID)
			}

			low, high := bot.GiveAwayLow(), bot.GiveAwayHigh()
			if bot.LowDstType == BotDst {
				lowBot := bots[bot.LowDst]
				lowBot.Take(low)
				bots[bot.LowDst] = lowBot
			} else {
				outputs[bot.LowDst] = low
			}
			if bot.HighDstType == BotDst {
				highBot := bots[bot.HighDst]
				highBot.Take(high)
				bots[bot.HighDst] = highBot
			} else {
				outputs[bot.HighDst] = high
			}

			bot.JobDone = true
			done++

			bots[botID] = bot
		}
	}

	fmt.Println("Part two:", outputs[0]*outputs[1]*outputs[2])
}

func initialState(input []string, bots map[int]Bot) {
	for _, line := range input {
		words := strings.Split(line, " ")
		if words[0] == "value" {
			value, _ := strconv.Atoi(words[1])
			botID, _ := strconv.Atoi(words[5])

			bot := bots[botID]
			bot.Take(value)
			bots[botID] = bot
		} else {
			botID, _ := strconv.Atoi(words[1])
			bot := bots[botID]

			if words[5] == "bot" {
				bot.LowDstType = BotDst
			} else {
				bot.LowDstType = OutputDst
			}

			if words[10] == "bot" {
				bot.HighDstType = BotDst
			} else {
				bot.HighDstType = OutputDst
			}

			bot.LowDst, _ = strconv.Atoi(words[6])
			bot.HighDst, _ = strconv.Atoi(words[11])

			bots[botID] = bot
		}
	}
}

func loadInput() []string {
	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(string(bytes), "\n")
}
