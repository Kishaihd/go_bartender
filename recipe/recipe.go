package recipe

// This is an ingredient model
// Super basic.

import (
	"fmt"
	"strconv"
)

type Measurement struct {
	quantity    float64
	measureType string
}

type Ingredient struct {
	name        string
	measurement Measurement
	// image <image type>
}

type Recipe struct {
	name        string
	description string
	ingredients []Ingredient
	tags        []string
}

func (i Ingredient) toString() string {
	outputString := ""
	// qFormat := ""
	// if i.measurement.quantity
	outputString = fmt.Sprintf("%-10s%-10s%-10s", strconv.FormatFloat(i.measurement.quantity, 'f', -1, 64), i.measurement.measureType, i.name)
	return outputString
}

func (r Recipe) Describe() string {
	ingredientList := ""
	for _, ingredient := range r.ingredients {
		ingredientList += fmt.Sprintf("%s\n", ingredient.toString())
	}
	output := fmt.Sprintf("%-s\n%-s\n\n", r.name, ingredientList)
	fmt.Println(output)
	return output
}

func (r *Recipe) AddIngredient(newName string, newQuantity float64, newMeasureType string) {
	m := Measurement{quantity: newQuantity, measureType: newMeasureType}
	newIngredient := Ingredient{name: newName, measurement: m}
	r.ingredients = append(r.ingredients, newIngredient)
}

func NewRecipe(name string) *Recipe {
	newRecipe := Recipe{name: name, description: "", ingredients: []Ingredient{}, tags: []string{}}
	return &newRecipe
}
