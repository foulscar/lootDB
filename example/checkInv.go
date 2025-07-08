package main

import (
	"strconv"
	"strings"
)

func handleCheckInv(args []string) (string, bool) {
	if len(args) != 1 {
		return "", false
	}

	sb := strings.Builder{}
	sb.WriteString("---Inventory---\n\n")
	sb.WriteString("Emerald (emerald): x")
	sb.WriteString(strconv.Itoa(mainGame.Inventory[emerald]))
	sb.WriteString("\n")

	for item, count := range mainGame.Inventory {
		if item == emerald {
			continue
		}
		sb.WriteString(mainGame.LootDB.ItemDB[item].Name)
		sb.WriteString(" (")
		sb.WriteString(string(item))
		sb.WriteString("): x")
		sb.WriteString(strconv.Itoa(count))
		sb.WriteString("\n")
	}

	sb.WriteString("\n---------------")

	return sb.String(), true
}
