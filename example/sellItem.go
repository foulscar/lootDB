package main

import (
	ldb "github.com/foulscar/lootDB"
	"math"
	"strconv"
)

func handleSellItem(args []string) (string, bool) {
	if len(args) != 3 {
		return "", false
	}

	itemID := ldb.ItemID(args[1])

	if itemID == emerald {
		return "cannot sell emerald", false
	}

	amountToSell, err := strconv.Atoi(args[2])
	if err != nil {
		return "", false
	}

	itemEntry := getItemEntry(itemID)
	if itemEntry == nil {
		return "item does not exist", false
	}

	if itemEntry.CalcWorth <= 0 {
		return "item is not worth anything", false
	}

	minNeeded := int(math.Ceil(1.0 / itemEntry.CalcWorth))
	currCount := mainGame.Inventory[itemID]

	if currCount < amountToSell {
		return "you do not have enough", false
	}

	if amountToSell < minNeeded {
		return "you need atleast x" + strconv.Itoa(minNeeded) + " to sell", false
	}

	mainGame.Inventory[itemID] -= amountToSell

	numEmeralds := int(float64(amountToSell) * itemEntry.CalcWorth)
	mainGame.Inventory[emerald] += numEmeralds

	return "sold x" + args[2] + " " + itemEntry.Name + " for x" + strconv.Itoa(numEmeralds) + " Emerald", true
}
