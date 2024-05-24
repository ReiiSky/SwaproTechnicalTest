package main

import (
	"github.com/ReiiSky/SwaproTechnical/sources/infrastructures/environments"
)

type Runable interface {
	Run()
}

func main() {
	env := environments.MustNewFileEnv("config/.env")

	var API Runable = FiberAPI(env)
	API.Run()
}
