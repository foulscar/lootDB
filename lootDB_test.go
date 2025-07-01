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
	fmt.Println(*db)
}
