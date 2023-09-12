package main

import (
	"fmt"
	"log"
	"time"

	"github.com/WhoAskedxD/anonymize_scans"
)

func main() {
	startTime := time.Now()
	folderPath := "/Users/harrymbp/Developer/Projects/PreXion/temp"
	dicomFolders, err := anonymize_scans.GetDicomFolders(folderPath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("dicom Folders are ", dicomFolders)
	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)
	fmt.Printf("Elapsed time: %.2f seconds\n", elapsedTime.Seconds())
}
