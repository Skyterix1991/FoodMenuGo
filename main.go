package main

import (
	"./controller"
	"./model"
	"./model/repository"
)

func main() {
	initializeTestingValues()
	controller.InitializeRestController()
}

func initializeTestingValues() {
	repository.AddFoodToMenu(model.NewFood(0, "Salty mountain", []string{"Salt", "More salt"}))
	repository.AddFoodToMenu(model.NewFood(1, "Cupcake", []string{"Butter", "Sugar"}))
	repository.AddFoodToMenu(model.NewFood(2, "Hotdog", []string{"Sausage", "Roll"}))
	repository.AddFoodToMenu(model.NewFood(3, "Cheeseburger", []string{"Roll", "Meat", "Cheese"}))

}
