package main

import (
	"fmt"
	imageprocessing "goroutines_pipeline/image_processing"
	"image"
	"os"
	"strings"
)

type Job struct {
	InputPath string
	Image     image.Image
	OutPath   string
}

func loadImage(paths []string) <-chan Job {
	out := make(chan Job)
	go func() {
		// For each input path create a job and add it to
		// the out channel
		for _, p := range paths {
			job := Job{InputPath: p,
				OutPath: strings.Replace(p, "images/", "images/output/", 1)}
			//added error handling
			img, err := imageprocessing.ReadImage(p)
			if err != nil {
				fmt.Println("error reading image", err)
				continue
			}
			job.Image = img
			out <- job
		}

		close(out)
	}()
	return out
}

func resize(input <-chan Job) <-chan Job {
	out := make(chan Job)
	go func() {
		// For each input job, create a new job after resize and add it to
		// the out channel
		for job := range input { // Read from the channel
			job.Image = imageprocessing.Resize(job.Image)
			out <- job
		}
		close(out)
	}()
	return out
}

func convertToGrayscale(input <-chan Job) <-chan Job {
	out := make(chan Job)
	go func() {
		for job := range input { // Read from the channel
			job.Image = imageprocessing.Grayscale(job.Image)
			out <- job
		}
		close(out)
	}()
	return out
}

func saveImage(input <-chan Job) <-chan bool {
	out := make(chan bool)
	go func() {
		for job := range input { // Read from the channel
			err := imageprocessing.WriteImage(job.OutPath, job.Image)
			if err != nil {
				fmt.Println("error writing image", err)
				out <- false
				continue
			}
			out <- true
		}
		close(out)
	}()
	return out
}

func main() {
	imagePaths := os.Args[1:]
	if len(imagePaths) == 0 {
		fmt.Println("Usage: program_name image_path1 image_path2 ...")
		return
	}

	channel1 := loadImage(imagePaths)
	channel2 := resize(channel1)
	channel3 := convertToGrayscale(channel2)
	writeResults := saveImage(channel3)

	for success := range writeResults {
		if success {
			fmt.Println("Success!")
		} else {
			fmt.Println("Failed!")
		}
	}
	fmt.Println("Pipeline complete!")
}
