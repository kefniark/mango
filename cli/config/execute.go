package config

type Executer interface {
	Execute(app string) error
}
