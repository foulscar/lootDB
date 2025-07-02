package lootDB

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

type LootDB struct {
	ItemDB  ItemDB
	TableDB TableDB
}

func NewLootDBFromDir(dirPath string) (*LootDB, error) {
	rootEntries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	tableCatDirs := make([]os.DirEntry, 0)

	var itemsCSV os.DirEntry
	for _, entry := range rootEntries {
		if entry.IsDir() {
			tableCatDirs = append(tableCatDirs, entry)
			continue
		}

		if entry.Name() == "items.csv" {
			itemsCSV = entry
		}
	}
	if itemsCSV == nil {
		return nil, errors.New("does not contain items.csv")
	}

	lootDB := LootDB{
		TableDB: make(TableDB),
	}

	file, err := os.Open(dirPath + "/" + itemsCSV.Name())
	if err != nil {
		return nil, err
	}

	itemDB, err := UnmarshalItemDB(file)
	if err != nil {
		return nil, err
	}
	file.Close()

	lootDB.ItemDB = *itemDB

	for _, entry := range tableCatDirs {
		tableCat := TableCat(entry.Name())
		lootDB.TableDB[tableCat] = make(map[ItemID]*Table)

		subEntries, err := os.ReadDir(dirPath + "/" + string(tableCat))
		if err != nil {
			return nil, err
		}

		for _, subEntry := range subEntries {
			if subEntry.IsDir() {
				continue
			}

			if filepath.Ext(subEntry.Name()) != ".json" {
				continue
			}

			fileNameParts := strings.Split(subEntry.Name(), ".json")
			itemID := ItemID(fileNameParts[0])
			if _, ok := lootDB.ItemDB[itemID]; !ok {
				return nil, errors.New(string(itemID) + " does not exist in ItemDB")
			}

			data, err := os.ReadFile(dirPath + "/" + string(tableCat) + "/" + subEntry.Name())
			if err != nil {
				return nil, err
			}

			table := Table{Index: string(tableCat) + "/" + string(itemID) + ".json"}
			err = UnmarshalTable(data, &table)
			if err != nil {
				return nil, err
			}

			lootDB.TableDB[tableCat][itemID] = &table
		}
	}

	lootDB.LinkItemDBWithTables()
	lootDB.CalculateItemsWorth(ItemID("emerald"))

	return &lootDB, nil
}

func (db *LootDB) IsValid() (bool, error) {
	if db.ItemDB == nil || db.TableDB == nil {
		return false, errors.New("fields are nil")
	}

	undefinedErr := errors.New("undefined ItemID")

	for _, tableCat := range db.TableDB {
		for itemID, table := range tableCat {
			if valid, err := table.IsValid(); !valid {
				return false, err
			}
			if _, ok := db.ItemDB[itemID]; !ok {
				return false, undefinedErr
			}
			for _, pool := range table.Pools {
				if valid, err := pool.IsValid(); !valid {
					return false, err
				}
				for _, entry := range pool.Entries {
					if _, ok := db.ItemDB[ItemID(entry.ID)]; !ok {
						return false, undefinedErr
					}
				}
			}
		}
	}

	return true, nil
}

func (db *LootDB) LinkItemDBWithTables() {
	for mainItemID := range db.ItemDB {
		itemEntry := db.ItemDB[mainItemID]
		for tableCat, entries := range db.TableDB {
			for itemID, table := range entries {
				if itemID == mainItemID {
					itemEntry.UsedFor = append(itemEntry.UsedFor, tableCat)
				}
				for _, pool := range table.Pools {
					for _, entry := range pool.Entries {
						if ItemID(entry.ID) == mainItemID {
							itemEntry.FoundIn[tableCat] = append(itemEntry.FoundIn[tableCat], itemID)
						}
					}
				}
			}
		}
	}
}
