package main

import (
	"fmt"
	"reflect"
)

// Test
// je déclare une structure, c'est l'équivalent d'une interface en TypeScript
type Test struct {
	id    int
	words string
}

// Modifiable
// et ça, c'est une interface, l'équivalent... d'une interface en PHP :p
type Modifiable interface {
	changeWords(words string)
	multiplyId(multiplier int) int
}

// Je demande à ce qu'on me renvoie un pointeur, comme ça je peux
// directement modifier la structure qu'on m'envoie, c'est vraiment
// le délire orienté objet
func (t *Test) changeWords(words string) {
	t.words = words
	return
}

// Ici je ne vais pas modifier la valeur qui est contenue dans ma structure,
// mais juste créer une copie et renvoyer une valeur qui sera corrélée avec
// les valeurs de ma struct... c'est plus le délire "static" mais pas tout à fait
func (t Test) multiplyId(multiplier int) int {
	t.id = t.id * multiplier
	return t.id
}

// Je crée une fonction qui demande un élément implémentant une interface comme argument
func interfaced(truc Modifiable) int {
	return truc.multiplyId(5)
}

func main() {
	// Je peux déclarer un emplacement mémoire et après l'affecter
	// ou utiliser directement l'opérateur ":=" pour inférer son type
	var t1 Test
	t1 = Test{
		id:    1,
		words: "coucou maman",
	}

	// Une api de reflect <3 !
	fmt.Println(t1, &t1, reflect.TypeOf(t1), reflect.TypeOf(&t1))
	// Output : {1 coucou maman} &{1 coucou maman} main.Test *main.Test

	t2 := &Test{
		id:    2,
		words: "Coucou référence",
	}

	// &t2 va me donner l'adresse mémoire du pointeur, un pointeur vers un pointeur
	fmt.Println(t2, &t2, reflect.TypeOf(t2), reflect.TypeOf(&t2))
	// Output : &{2 Coucou référence} 0xc0000aa020 *main.Test **main.Test

	t1.changeWords("francis")
	t2.changeWords("toujours une référence ?")

	val := t1.multiplyId(7)

	fmt.Println(t1, t2, val)
	// Output : {1 francis} &{2 toujours une référence ?} 7

	// Je peux faire ça car, comme j'ai implémenté toutes les méthodes,
	// t2 implémente de facto l'interface "Modifiable", je n'ai pas besoin de le
	// spécifier comme je le ferais en PHP avec "implement"
	fmt.Println(interfaced(t2))
	// Output : 10
}
