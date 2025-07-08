package main

var cmdHandlers = map[string]func([]string) (string, bool){
	"uses": handleCheckUses,
	"inv":  handleCheckInv,
	"sell": handleSellItem,
}

func initCMDHandlers() {
	for tableCat := range mainGame.LootDB.TableDB {
		cmdHandlers[string(tableCat)] = handleLootTable
	}
}
