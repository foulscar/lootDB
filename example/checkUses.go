package main

import (
	ldb "github.com/foulscar/lootDB"
	"strings"
)

func handleCheckUses(args []string) (string, bool) {
	if len(args) != 2 {
		return "", false
	}

	entry := getItemEntry(ldb.ItemID(args[1]))
	if entry == nil {
		return "item does not exist", false
	}

	result := strings.Builder{}
	result.WriteString("Uses: ")

	for _, tableCat := range entry.UsedFor {
		result.WriteString(string(tableCat) + " ")
	}

	return result.String(), true
}
