# dataPipelineProject

## Project Details

This project explores a data pipeline application for processing images in Go that utilizes concurrency with goroutines. The original pipeline and code repository was created by Amrit Singh, and the repository can be found here https://github.com/code-heim/go_21_goroutines_pipeline.

The original repository was cloned using the git clone command with the Github repo URL.


The go build command was used to build and the program was run. The output from that original build was saved in the init_build.txt for documenting purposes. Proper error checking was added to the image file input and output sections in the image_processing.go file. The original code included a panic that would stop the execution of the program. It is best practice to use explicit error handling by returning errors from a function, allowing the caller of the function to handle the error. The load and save image functions in the main.go file were also updated to reflect these changes for error handling. Four new images were added to the program. All four images were generated by Google Gemini, (Bard, Google AI. (2024)) using various prompts that included "generate pictures of bigfoot, sasquatch or gorillia".

Unit testing was added in the `main_test.go` file. The code in the program focuses on each component of the pipeline, read, write, grayscale, resize individually and the pipeline itself. When run the tests all pass. For documentation purposes, the output generated for the tests can be found in the testing.txt file with added context. Benchmarking was also included for each stage of the pipeline and then for the pipeline as a whole. According to the benchmark output, also found in the testing.txt file, loadIamge took 2 seconds, resizeImage took 4.4 seconds, ConvertToGrayscale took 4.5 seconds, and SaveImage took 4.6 seconds to run. Interestingly, when `go test bench=.` was run, to run all the benchmarks sequentially, my system crashed each time. Due to this issue the benchmark iterations were reduced as seen with the b.N/10 in each benchmark function. The final component, which is the entirety of the pipeline all at once, ran in 40 seconds.

In addition to adding error handling, code modifications were conducted in the main.go file. The original code provides images in the main function as input, however this section of the program was hardcoded as a slice of images. To avoid hard coding, the main function was modified, and an additional function was added titled runPipeline. The purpose of this addition was to incorporate command-line arguments and pass them to runPipeline to allow for a program that is more reusable with any image and to avoid hard coding.

### Running the Program
In the terminal, navigate to the directory that contains go_routines_pipeline.exe

**INPUT:**

 ./program.exe image_path/image

Where image_path is the path where the image is contained, and image is the name of the image file.

For this program it should look like this:

`./goroutines_pipeline.exe images/1.jpg images/2.jpg images/3.jpg images/4.jpg`

**OUTPUT:**

Success!

Success!

Success!

Success!

Pipeline complete!

If the command line arguments are incorrect the user will see this message:

Usage: program_name image_path1 image_path2 ...