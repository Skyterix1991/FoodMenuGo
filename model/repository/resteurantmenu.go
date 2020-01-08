package repository

import (
	"../../model"
)

var foodInMenu = make(map[int]*model.Food)

func AddFoodToMenu(food *model.Food) {
	foodInMenu[food.Id] = food
}

func GetFoodById(id int) *model.Food {
	var food = foodInMenu[id]

	return food
}

func UpdateFood(foodId int, partialFood *model.Food) {
	if !CheckIfFoodExists(foodId) {
		return
	}

	if partialFood.Name != "" {
		foodInMenu[foodId].Name = partialFood.Name
	}

	if partialFood.Ingredients != nil {
		foodInMenu[foodId].Ingredients = partialFood.Ingredients
	}
}

func ReplaceFood(foodId int, food *model.Food) {
	if !CheckIfFoodExists(foodId) {
		return
	}

	foodInMenu[foodId].Name = food.Name
	foodInMenu[foodId].Ingredients = food.Ingredients
}

func CheckIfFoodExists(foodId int) bool {
	return GetFoodById(foodId) != nil
}

func DeleteFoodFromMenu(foodId int) {
	if !CheckIfFoodExists(foodId) {
		return
	}

	foodInMenu[foodId] = nil
}

func GetAllFoodInMenu() []*model.Food {
	var allFood []*model.Food

	for _, value := range foodInMenu {
		if value == nil {
			continue
		}

		allFood = append(allFood, value)
	}

	return allFood
}