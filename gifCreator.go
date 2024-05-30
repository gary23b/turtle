// Tell the compile that this file should only be compiled for Windows, Linux, Mac.
//go:build linux || windows || darwin
// +build linux windows darwin

package turtle

import (
	"fmt"
	"image"
	"log"
	"time"

	"github.com/gary23b/easygif"
)

// Take an image of the turtle program window's contents and save it as a PNG.
func TakeScreenshot(window Window, outputPNGPath string) error {
	screenshot := window.GetCanvas().GetScreenshot()
	return easygif.SaveImageToPNG(screenshot, outputPNGPath)
}

// Take a set of screenshots of the turtle program window's content and return a slice of images.
// The length of the captured video will be delayBetweenScreenshots * frameCount
func TakeScreenshotVideo(
	window Window,
	delayBetweenScreenshots time.Duration,
	frameCount int,
) []image.Image {
	canvas := window.GetCanvas()

	// Collect the images
	frames := make([]image.Image, 0, frameCount)
	nextTime := time.Now()
	for frameIndex := 0; frameIndex < frameCount; frameIndex++ {
		screenShot := canvas.GetScreenshot()
		frames = append(frames, screenShot)

		nextTime = nextTime.Add(delayBetweenScreenshots)
		time.Sleep(time.Until(nextTime))
	}

	return frames
}

// Start this as a go routine to create a GIF of your creation.
// The 256 most common colors will be found and used.
// The length of the captured video will be delayBetweenScreenshots * frameCount
// The length of the gif will be delayBetweenGifFrames * frameCount
func CreateGif(
	window Window,
	delayBetweenScreenshots time.Duration,
	delayBetweenGifFrames time.Duration,
	outputGifFilePath string,
	frameCount int,
) {
	// Collect the images
	fmt.Printf("GIF: %s: Collecting images\n", outputGifFilePath)
	frames := TakeScreenshotVideo(window, delayBetweenScreenshots, frameCount)

	fmt.Printf("GIF: %s: Processing images\n", outputGifFilePath)
	err := easygif.MostCommonColorsWrite(frames, delayBetweenGifFrames, outputGifFilePath)
	if err != nil {
		log.Printf("Error while running easygif.EasyGifWrite(): %v\n", err)
	}

	fmt.Printf("GIF: %s: Done\n", outputGifFilePath)
}

// Start this as a go routine to create a GIF of your creation.
// Dithering will be used to achieve more colors, but images artifacts are introduced. This is best for images with lots of shading.
// The length of the captured video will be delayBetweenScreenshots * frameCount
// The length of the gif will be delayBetweenGifFrames * frameCount
func CreateGifDithered(
	window Window,
	delayBetweenScreenshots time.Duration,
	delayBetweenGifFrames time.Duration,
	outputGifFilePath string,
	frameCount int,
) {
	// Collect the images
	fmt.Printf("GIF: %s: Collecting images\n", outputGifFilePath)
	frames := TakeScreenshotVideo(window, delayBetweenScreenshots, frameCount)

	fmt.Printf("GIF: %s: Processing images\n", outputGifFilePath)
	err := easygif.DitheredWrite(frames, delayBetweenGifFrames, outputGifFilePath)
	if err != nil {
		log.Printf("Error while running easygif.EasyGifWrite(): %v\n", err)
	}

	fmt.Printf("GIF: %s: Done\n", outputGifFilePath)
}

// Start this as a go routine to create a GIF of your creation.
// This uses the nearest color in the Plan9 palette. So it is best for images with blocks of solid colors.
// The length of the captured video will be delayBetweenScreenshots * frameCount
// The length of the gif will be delayBetweenGifFrames * frameCount
func CreateGifNearestPlan9(
	window Window,
	delayBetweenScreenshots time.Duration,
	delayBetweenGifFrames time.Duration,
	outputGifFilePath string,
	frameCount int,
) {
	// Collect the images
	fmt.Printf("GIF: %s: Collecting images\n", outputGifFilePath)
	frames := TakeScreenshotVideo(window, delayBetweenScreenshots, frameCount)

	fmt.Printf("GIF: %s: Processing images\n", outputGifFilePath)
	err := easygif.NearestWrite(frames, delayBetweenGifFrames, outputGifFilePath)
	if err != nil {
		log.Printf("Error while running easygif.EasyGifWrite(): %v\n", err)
	}

	fmt.Printf("GIF: %s: Done\n", outputGifFilePath)
}
