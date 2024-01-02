// This is an ingredient model
// Super basic.
package recipe

import "fmt"

type Measurement struct {
	quantity    int
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
	return fmt.Sprintf("%-10s%-10s%-10s", i.measurement.quantity, i.measurement.measureType, i.name)
}

func (r Recipe) describe() string {
	ingredientList := ""
	for _, ingredient := range r.ingredients {
		ingredientList += fmt.Sprintf("%s\n", ingredient.toString())
	}
	output := fmt.Sprintf("%-s\n%-s\n\n", r.name, ingredientList)
}

//func (r Recipe) addIngredient(name string, quantity int, measureType string) {
//	r.ingredient.
//}

func (r Recipe) New(name string) Recipe {
	r := Recipe{name: name, description: "", ingredients: []Ingredient{}, tags: []string}
	return r
}
