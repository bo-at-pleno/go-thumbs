package helpers

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/color"
	"image/png"
	"os"

	"golang.org/x/image/draw"
)

// ThumbnailOptions holds the configuration for generating a thumbnail.
type ThumbnailOptions struct {
	Interpolation                    InterpolationMethod
	Width, Height                    int
	LowerPercentile, UpperPercentile float64
}

// InterpolationMethod specifies the type of interpolation to use for resizing.
type InterpolationMethod int

const (
	// NearestNeighbor interpolation method.
	NearestNeighbor InterpolationMethod = iota
	// Bilinear interpolation method.
	Bilinear
	// Bicubic interpolation method.
	Bicubic
)

// Thumbnail generates a thumbnail of the input image based on the provided options.
func Thumbnail(img *image.Image, options ThumbnailOptions) *image.Image {
	// Step 1: Resize the image using the specified interpolation method.
	resizedImg := resizeImage(img, options.Width, options.Height, options.Interpolation)
	// Step 2: Autocontrast the image based on the specified clipping percentiles.
	autoContrastedImg := autoContrast(&resizedImg, options.LowerPercentile, options.UpperPercentile)
	return &autoContrastedImg
}

// resizeImage resizes the input image using the specified interpolation method.
func resizeImage(img *image.Image, width, height int, interpolation InterpolationMethod) image.Image {
	// Create a new destination image with the specified width and height.
	dst := image.NewRGBA(image.Rect(0, 0, width, height))

	if interpolation == NearestNeighbor {
		draw.NearestNeighbor.Scale(dst, dst.Bounds(), *img, (*img).Bounds(), draw.Over, nil)
	} else if interpolation == Bilinear {
		draw.BiLinear.Scale(dst, dst.Bounds(), *img, (*img).Bounds(), draw.Over, nil)
	} else {
		draw.CatmullRom.Scale(dst, dst.Bounds(), *img, (*img).Bounds(), draw.Over, nil)
	}
	return dst
}

// autoContrast applies autocontrast to the input image based on the specified clipping percentiles.
func autoContrast(img *image.Image, lowerPercentile, upperPercentile float64) image.Image {
	// Calculate the histogram of the image.
	histogram := make(map[uint8]int)
	totalPixels := (*img).Bounds().Dx() * (*img).Bounds().Dy()
	bounds := (*img).Bounds()

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			pixel := color.GrayModel.Convert((*img).At(x, y)).(color.Gray)
			histogram[pixel.Y]++
		}
	}

	// Calculate the cumulative distribution function (CDF).
	var cdf [256]int
	sum := 0
	for i := 0; i < 256; i++ {
		sum += histogram[uint8(i)]
		cdf[i] = sum
	}

	// Calculate the lower and upper thresholds based on percentiles.
	lowerThreshold := int(lowerPercentile / 100 * float64(totalPixels))
	upperThreshold := int(upperPercentile / 100 * float64(totalPixels))

	// Calculate the intensity range based on the thresholds.
	var minIntensity, maxIntensity uint8
	for i := 0; i < 256; i++ {
		if cdf[i] >= lowerThreshold {
			minIntensity = uint8(i)
			break
		}
	}

	for i := 255; i >= 0; i-- {
		if cdf[i] <= upperThreshold {
			maxIntensity = uint8(i)
			break
		}
	}

	// Perform the linear contrast stretch.
	stretchedImg := image.NewGray(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			pixel := color.GrayModel.Convert((*img).At(x, y)).(color.Gray)
			intensity := pixel.Y
			newIntensity := uint8(float64(intensity-minIntensity) / float64(maxIntensity-minIntensity) * 255)
			stretchedImg.SetGray(x, y, color.Gray{Y: newIntensity})
		}
	}

	return stretchedImg
}

func ReadTiff(tiffPath string) (*image.Image, error) {
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

	return &tiffImage, nil
}

func ImageToBase64(img image.Image) (string, error) {
	// Convert the image to a byte slice.
	buffer := new(bytes.Buffer)
	err := png.Encode(buffer, img)
	if err != nil {
		return "", err
	}

	// Encode the byte slice as a base64 string.
	encodedString := base64.StdEncoding.EncodeToString(buffer.Bytes())
	return encodedString, nil
}
