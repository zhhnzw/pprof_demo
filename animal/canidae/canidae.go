package canidae

import "pprof_demo/animal"

type Canidae interface {
	animal.Animal
	Run()
	Howl()
}
