package maplearning

import "fmt"

func Main() {
	mapLearn()
}

func mapLearn() {
	hash := map[int]string{
		1: "Gurudas",
		2: "Naruto",
		3: "Luffy",
	}

	hash[1] = "abc"
	hash[2] = "abc"
	fmt.Println(hash)

	_, isExist := hash[12]

	if isExist {
		fmt.Println("FOUNDED")
	} else {
		fmt.Println("NOT FOUND")
	}

	for key, val := range hash {
		fmt.Println("KEY : ", key, " Value : ", val)
	}
	v, ok := hash[348]
	fmt.Println("The value:", v, "Present?", ok)
}
