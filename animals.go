package main

/* Inheritance in GO example

func xxtest() {

	var animals [5]Animal
	animals[0] = AnimalFactory("dog", "Barnie")
	animals[1] = AnimalFactory("dog", "Jackie")
	animals[2] = AnimalFactory("cat", "Mimi")
	animals[3] = AnimalFactory("dog", "Saddie")
	animals[4] = AnimalFactory("dog", "Saddie")

	for _, a := range animals {
		fmt.Println(a.Speak())
	}
}

func AnimalFactory(animal string, name string) Animal {
	var out Animal
	switch animal {
	case "dog":
		{
			dog := Dog{}
			dog.name = name
			out = dog
		}
	case "cat":
		cat := Cat{}
		cat.name = name
		out = cat
	}
	return out
}

type Animal interface {
	Speak() string
}

type DefaultAnimal struct {
	name string
}

func (a DefaultAnimal) Speak() string {
	return a.name + " says: ummmmm"
}

type Dog struct {
	DefaultAnimal
}


type Cat struct {
	DefaultAnimal
}

func (c Cat) Speak() string {
	return c.name + " says Meow!"
}

type Llama struct {
	DefaultAnimal
}

func (l Llama) Speak() string {
	return "?????"
}

type JavaProgrammer struct {
	DefaultAnimal
}

func (j JavaProgrammer) Speak() string {
	return "Design patterns!"
}
***/
