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
//ROOTUID = 1.2.392.200036.9163 | Unit UID = 31.0938 | Timestamp = 20231019103353867 || SUBUID = 1.1

func main() {
	startTime := time.Now()
	folderPath := "/Users/harrymbp/Developer/Projects/PreXion/temp/"
	outputPath := "/Volumes/Harrypc/temp/Anonymized scans"
	//grab parentFolder and scan details related to that parent folder
	dicomFolders, err := anonymize_scans.GetDicomFolders(folderPath)
	if err != nil {
		log.Fatal(err)
	}
	//data to modify the dicom info with. takes in tag type for key and value string. | https://pkg.go.dev/github.com/suyashkumar/dicom@v1.0.7/pkg/tag#pkg-constants
	// newDicomAttribute := map[tag.Tag]string{
	// 	tag.PatientName: "new testing function",
	// 	tag.PatientID:   "1700079492",
	// }
	//need to make a function to let users select which Tags they want to modify.
	//create a rootUID and UnitUID = 1.2.392.200036.9163.99.9999
	totalscans := len(dicomFolders)
	fmt.Printf("found %d scans\n", totalscans)
	counter := 0
	for parentPath, scanDetails := range dicomFolders {
		fmt.Printf("started working on scan: %s\n", parentPath)
		//generate a map that associates each scan with a new output path
		fmt.Printf("Generating output paths\n")
		folderInfo, err := anonymize_scans.MakeOutputPath(parentPath, outputPath, counter, scanDetails)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("outputPaths are %s\n", folderInfo)
		fmt.Printf("Generating new information for the dicom files\n")
		//take rootid and time stamp and add them together to create a studyUID for the parent folder and add it to the current newPatientInfo
		newPatientInfo, err := anonymize_scans.RandomizePatientInfo(scanDetails)
		if err != nil {
			log.Fatal(err)
		}
		//set the tag info for StudyInstanceuid to a value we will modify this when working with the scans.
		newPatientInfo[tag.StudyInstanceUID] = "1.2.392.200036.9163.99.9999"
		fmt.Printf("Finished generating new dicom info\n")
		logResults, _ := anonymize_scans.LogAnonymizedScan(scanDetails, newPatientInfo)
		fmt.Printf("Generating new dicoms\n")
		err = anonymize_scans.MakeDicomFolders(folderInfo, newPatientInfo)
		if err != nil {
			log.Fatal(err)
		}
		counter++
		//test code block
		fmt.Printf("finished anonymizing Dicoms for %s and ORGINIALPATIENTID is:%s, NEWPATIENTID is:%s\n", logResults["LOCATION"], logResults["ORGINIALPATIENTID"], logResults["NEWPATIENTID"])
		fmt.Printf("finished working on scan %d out of %d\n\n", counter, totalscans)
		// break
	}
	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)
	fmt.Printf("Total time Elapsed time: %.2f seconds\n", elapsedTime.Seconds())
}
