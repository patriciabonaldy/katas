package main

import (
	"fmt"
	"os"
	"strconv"
)

var (
	items = []Item{
		&item{"+5 Dexterity Vest", 10, 20},
		&Brie{item{name: "Aged Brie", sellIn: 2, quality: 0}},
		&item{"Elixir of the Mongoose", 5, 7},
		&Sulfuras{item{"Sulfuras, Hand of Ragnaros", 0, 80}},
		&Sulfuras{item{"Sulfuras, Hand of Ragnaros", -1, 80}},
		&BackStage{item{"Backstage passes to a TAFKAL80ETC concert", 15, 20}},
		&BackStage{item{"Backstage passes to a TAFKAL80ETC concert", 10, 49}},
		&BackStage{item{"Backstage passes to a TAFKAL80ETC concert", 5, 49}},
		&item{"Conjured Mana Cake", 3, 6}, // <-- :O
	}
	err error
)

func main() {
	days := 2
	if len(os.Args) > 1 {
		days, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		days++
	}

	for day := 0; day < days; day++ {
		fmt.Printf("-------- day %d --------\n", day)
		fmt.Println("name, sellIn, quality")
		for i := 0; i < len(items); i++ {
			fmt.Println(items[i])
		}
		UpdateQuality(items)
	}
}
