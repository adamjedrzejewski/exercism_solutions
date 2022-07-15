package lasagna

// PreparationTime calculates lasagna preparation time based on number of layers and average layer preparation Time.
func PreparationTime(layers []string, averageTime int) (time int) {
	if averageTime == 0 {
		averageTime = 2
	}

	time = len(layers) * averageTime
	return
}

// Quantities calculates required amount of noodles and sauce.
func Quantities(layers []string) (noodles int, sauce float64) {
	noodles = 0
	sauce = 0.0
	for _, layer := range layers {
		switch layer {
		case "noodles":
			noodles += 50
		case "sauce":
			sauce += 0.2
		}
	}
	return
}

// AddSecretIngredient replaces last element in myList with last element from friendsList.
func AddSecretIngredient(friendsList []string, myList []string) {
	friendsListLen := len(friendsList)
	myListLen := len(myList)
	if friendsListLen == 0 || myListLen == 0 {
		return
	}

	myList[myListLen-1] = friendsList[friendsListLen-1]
}

// ScaleRecipe calculates required values for given number portions based on the provided values for 2 portions.
func ScaleRecipe(inputList []float64, portions int) (values []float64) {
	values = make([]float64, len(inputList))
	copy(values, inputList)

	for i := range values {
		values[i] = float64(portions) * values[i] / 2
	}
	return
}
