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
		lootDB.TableDB[tableCat] = make(map[ItemID]Table)

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

			table := Table{}
			err = UnmarshalTable(data, &table)
			if err != nil {
				return nil, err
			}

			lootDB.TableDB[tableCat][itemID] = table
		}
	}

	return &lootDB, nil
}
