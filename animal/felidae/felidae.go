package felidae

import "pprof_demo/animal"

type Felidae interface {
	animal.Animal
	Climb()
	Sneak()
}
