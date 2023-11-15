package main

import (
	"fmt"
	"log"
	"time"

	"github.com/WhoAskedxD/anonymize_scans"
	"github.com/suyashkumar/dicom/pkg/tag"
)

//Notes
//UID instances
//Syntax for a UID is RootUID + Unit UID + Timestamp 17 digits long + SUBUID
//example 1.2.392.200036.9163.31.0938.20231019103353867.1.1
//ROOTUID = 1.2.392.200036.9163 | Unit UID = 31 | Timestamp = 20231019103353867 || SUBUID = 1.1

func main() {
	startTime := time.Now()
	folderPath := "/Users/harrymbp/Developer/Projects/PreXion/temp copy"
	outputPath := "/Users/harrymbp/Developer/Projects/PreXion/output"
	//grab parentFolder and scan details related to that parent folder
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
	for parentPath, scanDetails := range dicomFolders {
		fmt.Printf("current key is: %s\nValue: %s\n", parentPath, scanDetails)
		//generate a map that associates each scan with a new output path
		folderInfo, err := anonymize_scans.MakeOutputPath(parentPath, outputPath, counter, scanDetails)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("outputPaths for:%s\nare %s\n", parentPath, folderInfo)
		fmt.Printf("sending folderInfo to MakeDicomFolder")
		err = anonymize_scans.MakeDicomFolders(folderInfo, newDicomAttribute)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("done making Dicoms")
		counter++
		//test code block
		// listOfScans, err := anonymize_scans.GetScanList(scanDetails)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// fmt.Print(listOfScans)
		break
	}
	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)
	fmt.Printf("Total time Elapsed time: %.2f seconds\n", elapsedTime.Seconds())
}
