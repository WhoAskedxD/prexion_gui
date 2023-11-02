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
	dicomReference, err := anonymize_scans.GetScanNames(dicomFolders)
	for key, value := range dicomReference {
		fmt.Printf("parent folder is %s, while scan is %s\n", key, value)
	}
	if err != nil {
		log.Fatal(err)
	}
	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)
	fmt.Printf("Total time Elapsed time: %.2f seconds\n", elapsedTime.Seconds())
}
