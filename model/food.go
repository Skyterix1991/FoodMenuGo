package model

import (
	"encoding/json"
	"fmt"
)

type Food struct {
	Id          int
	Name        string
	Ingredients []string
}

func NewFood(id int, name string, ingredients []string) *Food {
	return &Food{Id: id, Name: name, Ingredients: ingredients}
}

func (b Food) MarshalJSON() []byte {
	e, err := json.Marshal(Food {
		Id:          b.Id,
		Name:        b.Name,
		Ingredients: b.Ingredients,
	})
	if err != nil {
		fmt.Println(err)

		return nil
	}

	return e
}
