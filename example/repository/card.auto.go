package repository

import (
	"fmt"

	"github.com/kasari/mmem/example/data"
)

type CardRepository struct {
	records        []*data.Card
	recordByID     map[uint64]*data.Card
	recordByUnitID map[uint64]*data.Card
}

func NewCardRepository() *CardRepository {
	return &CardRepository{}
}

func (repo *CardRepository) All() []*data.Card {
	return repo.records
}

func (repo *CardRepository) Where(predicate func(card *data.Card) bool) []*data.Card {
	result := []*data.Card{}
	for _, record := range repo.records {
		if predicate(record) {
			result = append(result, record)
		}
	}
	return result
}

func (repo *CardRepository) FindByID(ID uint64) *data.Card {
	return repo.recordByID[ID]
}

func (repo *CardRepository) FindByUnitID(UnitID uint64) *data.Card {
	return repo.recordByUnitID[UnitID]
}

func (repo *CardRepository) Add(card *data.Card) error {
	repo.records = append(repo.records, card)

	if _, ok := repo.recordByID[card.ID]; ok {
		return fmt.Errorf("can't add record because duplicate ID(%v)", card.ID)
	}
	repo.recordByID[card.ID] = card

	if _, ok := repo.recordByUnitID[card.UnitID]; ok {
		return fmt.Errorf("can't add record because duplicate UnitID(%v)", card.UnitID)
	}
	repo.recordByUnitID[card.UnitID] = card

	return nil
}
