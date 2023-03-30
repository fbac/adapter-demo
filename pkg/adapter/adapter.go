package adapter

type CompatibilityAdapter interface {
	Discover(endpoint string) error
	Run()
}
