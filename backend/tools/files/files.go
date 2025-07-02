package files

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

const MD_FOLDER string = "blogs_md_folder"

func InitalizeFolder() error {
	if _, err := os.Stat(MD_FOLDER); os.IsNotExist(err) {
		// Doesn't exist — try to create
		if err := os.Mkdir(MD_FOLDER, 0750); err != nil {
			return fmt.Errorf("failed to create folder: %w", err)
		}
	}
	return nil // Either aleady exists or created
}

func checkForFolder() error {
	if _, err := os.Stat(MD_FOLDER); os.IsNotExist(err) {
		return err
	}
	return nil
}

func UploadFile(uuid uuid.UUID, data []byte) error {
	if err := checkForFolder(); err != nil {
		return err
	}

	filename := filepath.Join(MD_FOLDER, uuid.String()+".md")
	if err := os.WriteFile(filename, data, 0750); err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}
	return nil
}
func ReadFile(uuid uuid.UUID) ([]byte, error) {
	if err := checkForFolder(); err != nil {
		return nil, err
	}
	filename := filepath.Join(MD_FOLDER, uuid.String()+".md")
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}
	return data, nil
}
