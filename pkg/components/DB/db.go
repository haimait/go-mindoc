package db

type Instance struct {
	DataSource string `yaml:"DataSource"`
}

const name = "DB"

var instance *Instance

func (i *Instance) Run() error {
	instance = i
	return nil
}

func (i *Instance) GetName() string {
	return name
}

func Get() *Instance {
	return instance
}
