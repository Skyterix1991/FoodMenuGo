package controller

import (
	"../model"
	"../model/repository"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func deleteFoodById(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	params := mux.Vars(request)

	foodId, _ := strconv.ParseInt(params["foodId"], 10, 64)

	if !repository.CheckIfFoodExists(int(foodId)) {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	response.WriteHeader(http.StatusOK)

	repository.DeleteFoodFromMenu(int(foodId))
}

func createFood(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusCreated)

	decoder := json.NewDecoder(request.Body)

	var partialFood model.Food

	_ = decoder.Decode(&partialFood)

	food := model.NewFood(partialFood.Id, partialFood.Name, partialFood.Ingredients)
	repository.AddFoodToMenu(food)
}

func getFoodById(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	params := mux.Vars(request)

	foodId, _ := strconv.ParseInt(params["foodId"], 10, 64)

	if !repository.CheckIfFoodExists(int(foodId)) {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	response.WriteHeader(http.StatusOK)

	var food *model.Food

	food = repository.GetFoodById(int(foodId))

	jsonResponse := food.MarshalJSON()

	response.Write(jsonResponse)
}

func getAllFood(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)

	allFood := repository.GetAllFoodInMenu()

	jsonResponse, _ := json.Marshal(allFood)

	response.Write(jsonResponse)
}

func updateFoodById(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	params := mux.Vars(request)

	foodId, _ := strconv.ParseInt(params["foodId"], 10, 64)

	if !repository.CheckIfFoodExists(int(foodId)) {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	response.WriteHeader(http.StatusOK)

	decoder := json.NewDecoder(request.Body)

	var partialFood model.Food

	_ = decoder.Decode(&partialFood)

	food := model.NewFood(partialFood.Id, partialFood.Name, partialFood.Ingredients)
	repository.UpdateFood(int(foodId), food)
}

func replaceFoodById(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	params := mux.Vars(request)

	foodId, _ := strconv.ParseInt(params["foodId"], 10, 64)

	if !repository.CheckIfFoodExists(int(foodId)) {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	response.WriteHeader(http.StatusOK)

	decoder := json.NewDecoder(request.Body)

	var partialFood model.Food

	_ = decoder.Decode(&partialFood)

	food := model.NewFood(partialFood.Id, partialFood.Name, partialFood.Ingredients)
	repository.ReplaceFood(int(foodId), food)
}

func InitializeRestController() {
	r := mux.NewRouter()

	r.HandleFunc("/food/", getAllFood).Methods(http.MethodGet)
	r.HandleFunc("/food/", createFood).Methods(http.MethodPost)
	r.HandleFunc("/food/{foodId}/", deleteFoodById).Methods(http.MethodDelete)
	r.HandleFunc("/food/{foodId}/", getFoodById).Methods(http.MethodGet)
	r.HandleFunc("/food/{foodId}/", updateFoodById).Methods(http.MethodPatch)
	r.HandleFunc("/food/{foodId}/", replaceFoodById).Methods(http.MethodPut)

	log.Fatal(http.ListenAndServe(":8080", r))
}
