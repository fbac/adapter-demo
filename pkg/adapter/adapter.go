package adapter

type CompatibilityAdapter interface {
	Run()
}

type Adapter struct {
	// Data, tests, etc
}

func (a Adapter) Run(rw CompatibilityAdapter) {
	rw.Run()
}
