package lootDB

import (
	"encoding/csv"
	"errors"
	"io"
)

type ItemID string

type ItemDB map[ItemID]*ItemDBEntry

type ItemDBEntry struct {
	Name      string
	CalcWorth float64
	FoundIn   map[TableCat][]ItemID
	UsedFor   []TableCat
}

func UnmarshalItemDB(r io.Reader) (*ItemDB, error) {
	csvReader := csv.NewReader(r)

	headers, err := csvReader.Read()
	if err != nil {
		return nil, err
	}
	if len(headers) < 3 {
		return nil, errors.New("csv is invalid")
	}

	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	itemDB := make(ItemDB)

	for _, record := range records {
		if len(record) < 3 {
			return nil, errors.New("csv is invalid")
		}
        if record[0] == "" {
            continue
        }

		itemDB[ItemID(record[0])] = &ItemDBEntry{
			Name:    record[1],
			FoundIn: make(map[TableCat][]ItemID),
			UsedFor: make([]TableCat, 0),
		}
	}

	return &itemDB, nil
}
