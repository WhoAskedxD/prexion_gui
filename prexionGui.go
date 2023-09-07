package main

import (
	"fmt"
	"log"

	"github.com/WhoAskedxD/anonymize_scans"
)

func main() {

	//folderPath := "/Users/harrymbp/Developer/Projects/PreXion/temp"
	dicomFilePath := "/Users/harrymbp/Developer/Projects/PreXion/temp/1.2.392.200036.9163.41.127414021.344261765/1.2.392.200036.9163.41.127414021.344261765.10632.1/00000_1.2.392.200036.9163.41.127414021.344261765.10632.2.dcm"
	results, err := anonymize_scans.DicomInfoGrabber(dicomFilePath)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}
	fmt.Println(results)

}
