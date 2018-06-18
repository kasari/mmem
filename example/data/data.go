package data

type Card struct {
	ID     uint64 `mmem:"unique"`
	UnitID uint64 `mmem:"unique"`
	Name   string
}

type Unit struct {
	ID     uint64 `mmem:"unique"`
	Name   string
	Attack uint32
}
