package config

import (
	"fmt"
	"os"
	"strconv"
)

func checkMissingEnv(name, value string) {
	if value == "" {
		panic(fmt.Sprintf("environment variable %s is missing", name))
	}
}

func GetEnvString(envName string) (env string) {
	env = os.Getenv(envName)
	checkMissingEnv(envName, env)
	return
}

func GetEnvInteger(envName string) (env int) {
	envString := os.Getenv(envName)
	checkMissingEnv(envName, envString)
	env, err := strconv.Atoi(envString)
	if err != nil {
		panic(fmt.Sprintf("Environment variable %s is not a integer", envName))
	}
	return
}

func GetEnvBoolean(envName string) bool {
	envString := os.Getenv(envName)
	checkMissingEnv(envName, envString)
	if envString != "true" && envString != "false" {
		panic(fmt.Sprintf("Environment variable %s is not a boolean", envName))
	}
	return envString == "true"
}
