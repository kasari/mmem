package repository

import (
	"fmt"

	"github.com/kasari/mmem/example/data"
)

type UnitRepository struct {
	records    []*data.Unit
	recordByID map[uint64]*data.Unit
}

func NewUnitRepository() *UnitRepository {
	return &UnitRepository{}
}

func (repo *UnitRepository) All() []*data.Unit {
	return repo.records
}

func (repo *UnitRepository) Where(predicate func(unit *data.Unit) bool) []*data.Unit {
	result := []*data.Unit{}
	for _, record := range repo.records {
		if predicate(record) {
			result = append(result, record)
		}
	}
	return result
}

func (repo *UnitRepository) FindByID(ID uint64) *data.Unit {
	return repo.recordByID[ID]
}

func (repo *UnitRepository) Add(unit *data.Unit) error {
	repo.records = append(repo.records, unit)

	if _, ok := repo.recordByID[unit.ID]; ok {
		return fmt.Errorf("can't add record because duplicate ID(%v)", unit.ID)
	}
	repo.recordByID[unit.ID] = unit

	return nil
}
