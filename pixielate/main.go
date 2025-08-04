package main

import (
	"fmt"
	"os"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func main() {
	inputPath := "iam.jpeg" // Can be PNG, JPG, etc.
	outputPath := "output_pixelated.png"
	pixelSize := 8 // The size of the pixel blocks

	// Check if input file exists
	if _, err := os.Stat(inputPath); os.IsNotExist(err) {
		fmt.Printf("Input file not found: %s\n", inputPath)
		return
	}

	// Build the FFmpeg filter chain
	// The logic is the same as for video: downscale and then upscale
	err := ffmpeg.Input(inputPath).
		// Downscale the image. iw and ih are input width and height
		Filter("scale", ffmpeg.Args{
			fmt.Sprintf("iw/%d", pixelSize),
			fmt.Sprintf("ih/%d", pixelSize),
		}).
		// Upscale it back to the original size using nearest-neighbor for the pixelated effect
		Filter("scale", ffmpeg.Args{
			fmt.Sprintf("iw*%d", pixelSize),
			fmt.Sprintf("ih*%d", pixelSize),
		}, ffmpeg.KwArgs{"flags": "neighbor"}).
		// Define the output file and overwrite if it exists
		Output(outputPath, ffmpeg.KwArgs{"y": ""}).
		OverWriteOutput().
		Run()

	if err != nil {
		fmt.Printf("Error running ffmpeg-go command: %v\n", err)
	} else {
		fmt.Printf("Image successfully pixelated and saved to %s\n", outputPath)
	}
}
