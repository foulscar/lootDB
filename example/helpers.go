package main

import (
	"fmt"
	ldb "github.com/foulscar/lootDB"
)

func getItemEntry(item ldb.ItemID) *ldb.ItemDBEntry {
	return mainGame.LootDB.ItemDB[item]
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func cleanInv() {
	for item, count := range mainGame.Inventory {
		if count > 0 {
			continue
		}
		delete(mainGame.Inventory, item)
	}
}
