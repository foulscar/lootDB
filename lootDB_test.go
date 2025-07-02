package lootDB

import (
	"fmt"
	"testing"
)

func TestExample(t *testing.T) {
	db, err := NewLootDBFromDir("example")
	if err != nil {
		t.Error(err)
	}
	if valid, err := db.IsValid(); !valid {
		t.Error(err)
	}
	fmt.Println(*db)
	fmt.Println()
	for itemID, val := range db.ItemDB {
		fmt.Println(string(itemID)+":", *val)
	}
	for _, catEntries := range db.TableDB {
		for _, table := range catEntries {
			fmt.Println("\n" + table.Index)
			fmt.Println(table.Roll())
		}
	}
}
