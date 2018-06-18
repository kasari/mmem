package main

import (
	"fmt"

	"github.com/kasari/mmem"
	"github.com/kasari/mmem/example/data"
	"github.com/kasari/mmem/example/repository"
)

func main() {
	err := mmem.Generate("./repository", []interface{}{
		&data.Card{},
		&data.Unit{},
	})

	if err != nil {
		fmt.Println(err)
	}
}

func example() {
	unitRepo := repository.NewUnitRepository()

	units := unitRepo.Where(func(unit *data.Unit) bool {
		return unit.Attack > 1000
	})

	for _, unit := range units {
		fmt.Printf("%+v\n", unit)
	}
}
