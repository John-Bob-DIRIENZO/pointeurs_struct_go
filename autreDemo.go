package main

import "fmt"

type Employee struct {
	ID      int
	Name    string
	Address string
}

func demoOne() {

	// C'est une référence
	andy := &Employee{}

	andy.Name = "Andy"

	// Du coup, les deux éléments pointent vers la même allocation mémoire
	brad := andy

	// Donc modifier l'un ou l'autre modifie la même donnée
	brad.Name = "Brad"

	fmt.Println(andy.Name)
	// Output : "Brad"
}

func demoTwo() {

	// Cette fois, ce n'est pas une référence, mais bien une nouvelle entité
	andy := Employee{}

	andy.Name = "Andy"

	// Donc ici, je ne fais pas une copie du pointeur (et donc de l'adresse mémoire)
	// mais une copie de l'entité, j'ai donc maintenant, deux entités différentes
	brad := andy

	// Donc modifier l'un ne modifie pas l'autre
	brad.Name = "Brad"

	fmt.Println(andy.Name)
	// Output : "Andy"
}
