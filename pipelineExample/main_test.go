package main

import (
	imageprocessing "goroutines_pipeline/image_processing"
	"testing"
)

func loadImageForTesting(t *testing.T) []string {
	return []string{"images/1.jpg", "images/2.jpg", "images/3.jpg"}
}

func TestImageProcessingFunctions(t *testing.T) {
	tests := []struct {
		name     string
		function func() error
	}{
		{"ReadImage", func() error {
			_, err := imageprocessing.ReadImage("images/1.jpg")
			return err
		}},
		{"WriteImage", func() error {
			img, err := imageprocessing.ReadImage("images/1.jpg")
			if err != nil {
				return err
			}
			return imageprocessing.WriteImage("images/1_copy.jpg", img)
		}},
		{"Grayscale", func() error {
			img, err := imageprocessing.ReadImage("images/1.jpg")
			if err != nil {
				return err
			}
			grayImg := imageprocessing.Grayscale(img)
			return imageprocessing.WriteImage("images/1_gray.jpg", grayImg)
		}},
		{"Resize", func() error {
			img, err := imageprocessing.ReadImage("images/1.jpg")
			if err != nil {
				return err
			}
			resizedImg := imageprocessing.Resize(img)
			return imageprocessing.WriteImage("images/1_resized.jpg", resizedImg)
		}},
		{"Pipeline", func() error {
			images := loadImageForTesting(t)
			pipeline := convertToGrayscale(resize(loadImage(images)))
			for job := range pipeline {
				err := imageprocessing.WriteImage(job.OutPath, job.Image)
				if err != nil {
					return err
				}
			}
			return nil
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.function(); err != nil {
				t.Errorf("Test %s failed: %v", tt.name, err)
			}
		})
	}
}

func BenchmarkPipeline(b *testing.B) {
	imagePaths := []string{"images/1.jpg", "images/2.jpg", "images/3.jpg", "images/4.jpg"}

	for i := 0; i < b.N/10; i++ { // Reduce the number of iterations
		//create Pipeline
		channel1 := loadImage(imagePaths)
		channel2 := resize(channel1)
		channel3 := convertToGrayscale(channel2)
		_ = saveImage(channel3)
	}
}
