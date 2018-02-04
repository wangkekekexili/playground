package main

import (
	"image"
	"log"

	"gocv.io/x/gocv"
)

func main() {
	webcam, err := gocv.VideoCaptureDevice(0)
	if err != nil {
		log.Fatal(err)
	}
	defer webcam.Close()

	window := gocv.NewWindow("blur-face")
	defer window.Close()

	img := gocv.NewMat()
	defer img.Close()

	classifier := gocv.NewCascadeClassifier()
	defer classifier.Close()
	ok := classifier.Load("haarcascade_frontalface_default.xml")
	if !ok {
		log.Fatal("failed to load classifier xml file")
	}

	for {
		ok = webcam.Read(img)
		if !ok {
			log.Fatal("failed to read image")
		}
		if img.Empty() {
			log.Fatal("image is empty")
		}

		blurface(img, classifier)

		window.IMShow(img)
		window.WaitKey(1)
	}
}

func blurface(img gocv.Mat, classifier gocv.CascadeClassifier) {
	rects := classifier.DetectMultiScale(img)
	if len(rects) == 0 {
		return
	}

	for _, rect := range rects {
		subImg := img.Region(rect)
		defer subImg.Close()
		gocv.GaussianBlur(subImg, subImg, image.Pt(43, 43), 30, 50, 4)
	}
}
