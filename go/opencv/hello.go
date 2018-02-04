package main

import (
	"log"

	"gocv.io/x/gocv"
)

func main() {
	webcam, err := gocv.VideoCaptureDevice(0)
	if err != nil {
		log.Fatal(err)
	}
	defer webcam.Close()

	window := gocv.NewWindow("hell0")
	defer window.Close()

	img := gocv.NewMat()
	defer img.Close()

	for {
		webcam.Read(img)
		window.IMShow(img)
		window.WaitKey(1)
	}
}
