package config

type Executer interface {
	Name() string
	Execute(app string) error
}
