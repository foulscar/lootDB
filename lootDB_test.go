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
    for _, entry := range db.ItemDB {
        fmt.Println(entry.Name+":", entry.CalcWorth)
    }
	for _, catEntries := range db.TableDB {
		for _, table := range catEntries {
			fmt.Println("\n" + table.Index)
			fmt.Println(table.Roll())
		}
	}
}
