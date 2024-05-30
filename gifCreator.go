package turtle

import (
	"fmt"
	"image"
	"log"
	"time"

	"github.com/gary23b/easygif"
)

func TakeScreenshot(window Window, outputPNGPath string) error {
	screenshot := window.GetCanvas().GetScreenshot()
	return easygif.SaveImageToPNG(screenshot, outputPNGPath)
}

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
