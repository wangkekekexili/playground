package main

import (
	"image/color"
	"log"

	"gocv.io/x/gocv"
)

func main() {
	webcam, err := gocv.VideoCaptureDevice(0)
	if err != nil {
		log.Fatal(err)
	}
	defer webcam.Close()

	window := gocv.NewWindow("smile")
	defer window.Close()

	img := gocv.NewMat()
	defer img.Close()

	classifier := gocv.NewCascadeClassifier()
	defer classifier.Close()
	ok := classifier.Load("haarcascade_eye.xml")
	if !ok {
		log.Fatal("failed to load classifier file")
	}

	for {
		ok = webcam.Read(img)
		if !ok {
			log.Fatal("failed to read from webcam")
		}
		if img.Empty() {
			log.Fatal("imag is empty")
		}

		blue := color.RGBA{B: 255, A: 255}
		rects := classifier.DetectMultiScale(img)
		for _, rect := range rects {
			gocv.Rectangle(img, rect, blue, 1)
		}

		window.IMShow(img)
		window.WaitKey(1)
	}
}
