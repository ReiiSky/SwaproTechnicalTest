package environments

import (
	"fmt"

	"os"

	"github.com/joho/godotenv"
)

func MustNewFileEnv(filepath string) Env {
	readedBytes, err := os.ReadFile(filepath)

	if err != nil {
		panic(fmt.Errorf("error read file %s", filepath))
	}

	readedFileContent := string(readedBytes)
	envMap, err := godotenv.Unmarshal(readedFileContent)

	if err != nil {
		panic(fmt.Errorf("error parse conf key: %s", err))
	}

	return Env{
		envMap,
	}
}
