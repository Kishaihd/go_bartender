package main

func main() {
	oldFashioned := NewRecipe("Old Fashioned")
	// oldFashioned := NewRecipe("Old Fashioned")
	oldFashioned.AddIngredient("bourbon", 1.5, "oz")
	oldFashioned.AddIngredient("angostura bitters", 3, "dashes")
	oldFashioned.AddIngredient("simple syrup", 0.5, "oz")
	oldFashioned.AddIngredient("orange peel", 1, "strip")
	oldFashioned.Describe()
}
