package main

import (
	ldb "github.com/foulscar/lootDB"
	"strconv"
	"strings"
)

func handleLootTable(args []string) (string, bool) {
	if len(args) != 3 {
		return "must provide itemID and count", false
	}

	tableCat := ldb.TableCat(args[0])
	if _, ok := mainGame.LootDB.TableDB[tableCat]; !ok {
		return "invalid table", false
	}

	itemID := ldb.ItemID(args[1])
	if _, ok := mainGame.LootDB.ItemDB[itemID]; !ok {
		return "item does not exist", false
	}

	count, err := strconv.Atoi(args[2])
	if err != nil {
		return "must provide count", false
	}

	if mainGame.Inventory[itemID] < count {
		return "you do not have enough", false
	}

	usesTable := false
	for _, usableTableCat := range mainGame.LootDB.ItemDB[itemID].UsedFor {
		if usableTableCat == tableCat {
			usesTable = true
		}
	}
	if !usesTable {
		return "this item cannot be used for this", false
	}

	table := mainGame.LootDB.TableDB[tableCat][itemID]
	if table == nil {
		return "error", false
	}

	finalResult := make(map[ldb.ItemID]int)

	for i := 0; i < count; i++ {
		result := table.Roll()
		for item, count := range result {
			finalResult[item] += count
		}
	}

	mainGame.Inventory[itemID] -= count

	resultStr := strings.Builder{}
	resultStr.WriteString("---Result---\n\n")

	for item, count := range finalResult {
		mainGame.Inventory[item] += count
		resultStr.WriteString("x" + strconv.Itoa(count))
		resultStr.WriteString(" " + mainGame.LootDB.ItemDB[item].Name + "\n")
	}

	resultStr.WriteString("\n------------")

	return resultStr.String(), true
}
