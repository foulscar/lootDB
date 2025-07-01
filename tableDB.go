package lootDB

import (
	"encoding/json"
	"errors"
	"strconv"
)

type TableCat string

type TableDB map[TableCat]map[ItemID]Table

type Table struct {
	Pools []Pool `json:"pools"`
}

type Pool struct {
	Rolls   int         `json:"rolls"`
	Chance  float64     `json:"chance"`
	Entries []PoolEntry `json:"entries"`
}

type PoolEntry struct {
	Type   string `json:"type"`
	ID     string `json:"id"`
	Weight int    `json:"weight"`
}

var ValidPoolEntryTypes = []string{
	"item",
	"loottable",
}

func MarshalTable(table *Table) ([]byte, error) {
	isValid, err := table.IsValid()
	if !isValid {
		return nil, errors.New("table is invalid. " + err.Error())
	}

	return json.Marshal(table)
}

func UnmarshalTable(data []byte, table *Table) error {
	err := json.Unmarshal(data, &table)
	if err != nil {
		return err
	}

	isValid, err := table.IsValid()
	if !isValid {
		return errors.New("table is invalid. " + err.Error())
	}

	return nil
}

func (t *Table) IsValid() (bool, error) {
	if t.Pools == nil {
		return false, errors.New("does not contain a 'pools' array")
	}

	for i, pool := range t.Pools {
		poolName := "pools[" + strconv.Itoa(i) + "]"
		valid, err := pool.IsValid()
		if !valid {
			return false, errors.New(poolName + " is invalid. " + err.Error())
		}
	}

	return true, nil
}

func (p Pool) IsValid() (bool, error) {
	if p.Rolls <= 0 {
		return false, errors.New("contains an invalid 'rolls' field")
	}
	if p.Chance <= 0 || p.Chance > 1 {
		return false, errors.New("contains an invalid 'chance' field")
	}
	if p.Entries == nil {
		return false, errors.New("does not contain an 'entries' field")
	}

	for i, entry := range p.Entries {
		entryName := "entries[" + strconv.Itoa(i) + "]"
		entryIsValid, err := entry.IsValid()
		if !entryIsValid {
			return false, errors.New(entryName + " is invalid. " + err.Error())
		}
	}

	return true, nil
}

func (e PoolEntry) IsValid() (bool, error) {
	typeIsValid := false
	for _, validType := range ValidPoolEntryTypes {
		if e.Type == validType {
			typeIsValid = true
			break
		}
	}
	if !typeIsValid {
		return false, errors.New("does not contain a valid 'type' field")
	}

	if e.ID == "" {
		return false, errors.New("does not contain an 'id' field or is empty")
	}

	if e.Weight <= 0 {
		return false, errors.New("does not contain a valid 'weight' field")
	}

	return true, nil
}
