package turtleutil

import (
	"bytes"
	"fmt"
	"image"
	"os"
)

// Load an image file and covert it to an image.Image.
// For decode to work on file type, it must be registered by including the codec specific package.
func LoadSpriteFile(path string) (image.Image, error) {
	spriteFileData, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("Failed to read sprite file: %s, %w", path, err)
	}
	img, _, err := image.Decode(bytes.NewReader(spriteFileData))
	if err != nil {
		return nil, fmt.Errorf("Failed to decode image data: %s, %w", path, err)
	}
	return img, nil
}
