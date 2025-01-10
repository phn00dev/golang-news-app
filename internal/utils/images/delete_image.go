package images

import "os"

func DeleteImage(imagePath string) error {
	if err := os.Remove(imagePath); err != nil {
		return err
	}
	return nil
}
