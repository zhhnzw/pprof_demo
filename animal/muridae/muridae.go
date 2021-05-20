package muridae

import "pprof_demo/animal"

type Muridae interface {
	animal.Animal
	Hole()
	Steal()
}
