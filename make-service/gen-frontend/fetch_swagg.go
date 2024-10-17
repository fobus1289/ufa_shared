package genfrontend

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

var spec OpenAPISpec

func ReadOpenAPISpec(filePath string) error {

	file, err := os.Open(filePath)
	{
		if err != nil {
			return fmt.Errorf("error opening file: %w", err)
		}
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	{
		if err != nil {
			return fmt.Errorf("error reading file: %w", err)
		}
	}

	{
		if err = json.Unmarshal(data, &spec); err != nil {
			return fmt.Errorf("error unmarshaling JSON: %w", err)
		}
	}

	return nil
}
