package helpers

import goenv "github.com/ricdotnet/goenvironmental"

var Envs = goenv.EnvVariables

func PrepareEnv() {
	goenv.ParseEnv()
	Envs = goenv.EnvVariables
}