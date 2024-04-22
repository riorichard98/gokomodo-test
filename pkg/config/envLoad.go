package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func LoadEnv(fileName string) {
	err := loadEnvFromFile(fileName)
	if err != nil {
		panic(fmt.Sprintf("cannot load env variables with erro: %v", err.Error()))
	}
}

func loadEnvFromFile(filename string) error {
	// Open the .env file
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Split each line into key-value pairs
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue // Skip invalid lines
		}
		// Set the environment variable
		os.Setenv(parts[0], parts[1])
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
