package lootDB

import (
	"math/rand"
)

func (t *Table) Roll() map[ItemID]int {
	result := make(map[ItemID]int)

	for _, pool := range t.Pools {
		for currRoll := 1; currRoll <= pool.Rolls; currRoll++ {
			if rand.Float64() > pool.Chance {
				continue
			}

			helperNum := 0.0
			entryChosen := rand.Float64()
			for _, entry := range pool.Entries {
				helperNum += entry.CalcChance
				if entryChosen > helperNum {
					continue
				}
				if entry.Type != "item" {
					break
				}

				rangeDif := entry.CountMax - entry.CountMin
				if rangeDif < 1 {
					result[ItemID(entry.ID)]++
					break
				}

				count := rand.Intn(rangeDif) + entry.CountMin
				if count > 0 {
					result[ItemID(entry.ID)] += count
				}
				break
			}
		}
	}

	return result
}
