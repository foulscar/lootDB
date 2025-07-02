package lootDB

func (db *LootDB) CalculateItemsWorth(rootItem ItemID) {
	db.ItemDB[rootItem].CalcWorth = 1

	iterateTable := func(rootWorth float64, table *Table) []ItemID {
		finalAverageItemCount := make(map[ItemID]float64)

		for _, pool := range table.Pools {
			averageRolls := float64(pool.Rolls) * pool.Chance
			averageItemCount := make(map[ItemID]float64)

			for _, entry := range pool.Entries {
				rangeMean := float64(entry.CountMin+entry.CountMax) / 2.0
				averageCount := rangeMean * entry.CalcChance

				averageItemCount[ItemID(entry.ID)] += averageCount * averageRolls
			}

			for itemID, averageCount := range averageItemCount {
				finalAverageItemCount[itemID] += averageCount
			}
		}

		changedItems := make([]ItemID, 0)

		for itemID, averageCount := range finalAverageItemCount {
			newWorth := rootWorth / averageCount

			currWorth := db.ItemDB[itemID].CalcWorth
			if currWorth != 0.0 && currWorth < newWorth {
				continue
			}
			db.ItemDB[itemID].CalcWorth = newWorth
			changedItems = append(changedItems, itemID)
		}

		return changedItems
	}

	itemsToIterate := []ItemID{rootItem}
	for i := 0; i < len(itemsToIterate); i++ {
		currItem := itemsToIterate[i]
		rootWorth := db.ItemDB[currItem].CalcWorth
		for _, tableCat := range db.ItemDB[currItem].UsedFor {
			table := db.TableDB[tableCat][currItem]
			changedItems := iterateTable(rootWorth, table)
			itemsToIterate = append(itemsToIterate, changedItems...)
		}
	}
}
