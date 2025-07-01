package lootDB

import (
	"encoding/csv"
	"errors"
	"io"
)

type ItemID string

type ItemDB map[ItemID]ItemDBEntry

type ItemDBEntry struct {
	Name    string
	FoundIn map[TableCat][]*Table
	UsedFor map[TableCat][]*Table
}

func UnmarshalItemDB(r io.Reader) (*ItemDB, error) {
	csvReader := csv.NewReader(r)

	headers, err := csvReader.Read()
	if err != nil {
		return nil, err
	}
	if len(headers) != 2 {
		return nil, errors.New("csv is invalid")
	}

	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	itemDB := make(ItemDB)

	for _, record := range records {
		if len(record) != 2 {
			return nil, errors.New("csv is invalid")
		}

		itemDB[ItemID(record[0])] = ItemDBEntry{
			Name:    record[1],
			FoundIn: make(map[TableCat][]*Table),
			UsedFor: make(map[TableCat][]*Table),
		}
	}

	return &itemDB, nil
}
