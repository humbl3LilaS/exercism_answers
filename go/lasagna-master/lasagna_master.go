package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, time int) int {
	if time == 0 {
		return len(layers) * 2
	}
	return len(layers) * time
}

// TODO: define the 'Quantities()' function

func Quantities(layers []string) (int, float64) {
	noddleQty := 0
	sauceQty := 0.0
	for _, layer := range layers {
		if layer == "noodles" {
			noddleQty += 50
			continue
		}

		if layer == "sauce" {
			sauceQty += 0.2
			continue
		}
	}
	return noddleQty, sauceQty
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friList []string, myList []string) {
	secretIng := friList[len(friList)-1]
	idx := len(myList) - 1
	myList = append(myList[0:idx], secretIng)
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portion int) []float64 {
	result := []float64{}
	for _, quan := range quantities {

		result = append(result, quan*(float64(portion)/2.0))
	}
	return result
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
