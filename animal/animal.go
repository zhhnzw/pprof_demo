package animal

import (
	"pprof_demo/animal/canidae/dog"
	"pprof_demo/animal/canidae/wolf"
	"pprof_demo/animal/felidae/cat"
	"pprof_demo/animal/felidae/tiger"
	"pprof_demo/animal/muridae/mouse"
)

var (
	AllAnimals = []Animal{
		&dog.Dog{},
		&wolf.Wolf{},

		&cat.Cat{},
		&tiger.Tiger{},

		&mouse.Mouse{},
	}
)

type Animal interface {
	Name() string
	Live()

	Eat()
	Drink()
	Shit()
	Pee()
}
