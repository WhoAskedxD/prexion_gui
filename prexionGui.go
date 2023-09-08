package main

import (
	"fmt"
	"log"

	"github.com/WhoAskedxD/anonymize_scans"
)

func main() {

	folderPath := "/Users/harrymbp/Developer/Projects/PreXion/temp"
	//dicomFilePath := "/Users/harrymbp/Developer/Projects/PreXion/dicomFiles/2313920.1194420868.1125777922.144317013718248927734763.2419061/2313920.1194420868.1125777922.206729801520575063.241906.824.11/00001_2313920.1194420868.1125777922.206415228720675263.241906.824.21.dcm"
	results, err := anonymize_scans.GetFilePathsInFolders(folderPath)
	if err != nil {
		log.Println("ERROR:", err)
		fmt.Println("not a valid dicom file.")
		return
	}
	for _, file := range results {
		anonymize_scans.GetDicomFolders(folderPath, file)
	}

}
