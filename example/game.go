package main

import (
	"bufio"
	"fmt"
	ldb "github.com/foulscar/lootDB"
	"os"
	"strings"
)

var mainGame *Game

type Game struct {
	LootDB    *ldb.LootDB
	Inventory map[ldb.ItemID]int
}

func initGame() error {
	db, err := ldb.NewLootDBFromDir("lootDB")
	if err != nil {
		return err
	}
	if valid, err := db.IsValid(); !valid {
		return err
	}

	mainGame = &Game{
		LootDB:    db,
		Inventory: make(map[ldb.ItemID]int),
	}

	mainGame.Inventory[emerald] = 50

	return nil
}

func play() {
	sb := strings.Builder{}
	sb.WriteString("Available commands: uses inv sell ")

	for cmd := range cmdHandlers {
		switch cmd {
		case "uses", "inv", "sell":
			continue
		}
		sb.WriteString(cmd)
		sb.WriteString(" ")
	}

	availableCMDs := sb.String()

	lastMsg := availableCMDs

	invalid := availableCMDs + "\n\n" + invalidCMD

	reader := bufio.NewReader(os.Stdin)

	for {
		cleanInv()
		clearScreen()
		fmt.Println(lastMsg)

		fmt.Print("\n\n> ")
		input, _ := reader.ReadString('\n')
		args := strings.Fields(input)

		handler, ok := cmdHandlers[args[0]]
		if !ok {
			lastMsg = invalid
			continue
		}

		result, ok := handler(args)
		if !ok && result == "" {
			lastMsg = invalid
			continue
		}

		lastMsg = availableCMDs + "\n\n" + result
	}
}
