package main

import (
	"fmt"
	"log"
	"time"

	"github.com/WhoAskedxD/anonymize_scans"
	"github.com/suyashkumar/dicom/pkg/tag"
)

func main() {
	startTime := time.Now()
	folderPath := "/Users/harrymbp/Developer/Projects/PreXion/temp"
	outputPath := "/Users/harrymbp/Developer/Projects/PreXion/output"
	dicomFolders, err := anonymize_scans.GetDicomFolders(folderPath)
	if err != nil {
		log.Fatal(err)
	}
	// listOfScans, err := anonymize_scans.GetScanNames(dicomFolders)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//data to modify the dicom info with. takes in tag type for key and value string. | https://pkg.go.dev/github.com/suyashkumar/dicom@v1.0.7/pkg/tag#pkg-constants
	newDicomAttribute := map[tag.Tag]string{
		tag.PatientName: "new testing function",
	}
	//need to make a function to let users select which Tags they want to modify.

	counter := 0
	for key, value := range dicomFolders {
		fmt.Printf("current key is: %s\nValue: %s\n", key, value)
		folderInfo, err := anonymize_scans.MakeOutputPath(key, outputPath, counter, value)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("outputPaths for:%s\nare %s\n", key, folderInfo)
		fmt.Printf("sending folderInfo to MakeDicomFolder")
		err = anonymize_scans.MakeDicomFolders(folderInfo, newDicomAttribute)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("done making Dicoms")
		counter++
		break
	}
	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)
	fmt.Printf("Total time Elapsed time: %.2f seconds\n", elapsedTime.Seconds())
}
