package helpers

import (
	"image"
	"image/jpeg"
	"os"
)

func readTiffAndCreateThumbnail(tiffPath string) (*image.Image, error) {
	// Read the TIFF file.
	tiffFile, err := os.Open(tiffPath)
	if err != nil {
		return nil, err
	}

	defer tiffFile.Close()

	tiffImage, _, err := image.Decode(tiffFile)
	if err != nil {
		return nil, err
	}

	// Create a thumbnail from the TIFF image.
	thumbnail, err := thumbnail.Thumbnail(tiffImage, 100, 100, image.Lanczos3)
	if err != nil {
		return nil, err
	}

	return thumbnail, nil
}

func saveThumbnail(thumbnail *image.Image, path string) error {
	// Save the thumbnail to a file.
	thumbnailFile, err := os.Create(path)
	if err != nil {
		return err
	}
	defer thumbnailFile.Close()

	options := &jpeg.Options{Quality: 75}
	err = jpeg.Encode(thumbnailFile, *thumbnail, options)
	return err
}
