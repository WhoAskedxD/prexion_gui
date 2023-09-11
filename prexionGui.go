package main

import (
	"github.com/WhoAskedxD/anonymize_scans"
)

func main() {

	folderPath := "/Users/harrymbp/Developer/Projects/PreXion/temp"
	searchPath := "/Users/harrymbp/Developer/Projects/PreXion/temp/1.2.392.200036.9163.41.127414021.344261765"
	//dicomFilePath := "/Users/harrymbp/Developer/Projects/PreXion/dicomFiles/2313920.1194420868.1125777922.144317013718248927734763.2419061/2313920.1194420868.1125777922.206729801520575063.241906.824.11/00001_2313920.1194420868.1125777922.206415228720675263.241906.824.21.dcm"
	// results, err := anonymize_scans.GetFilePathsInFolders(folderPath)
	// if err != nil {
	// 	log.Println("ERROR:", err)
	// 	fmt.Println("not a valid dicom file.")
	// 	return
	// }
	testPath := "/Volumes/Harrypc/temp/PX4Scan/20230817065427_1_Ceph/40920587/45724265/00001DCM.dcm"
	anonymize_scans.GetDicomFolders(folderPath, searchPath)
	anonymize_scans.CheckScanType(testPath)
}
