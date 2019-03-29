package main

type FabricCut struct {
	id            int64
	startPosition *Dimension
	size          *Dimension
}

func newFabricCut() *FabricCut {
	instance := new(FabricCut)
	instance.startPosition = new(Dimension)
	instance.size = new(Dimension)

	return instance
}
