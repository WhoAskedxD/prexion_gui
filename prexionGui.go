package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/WhoAskedxD/anonymize_scans"
	"github.com/suyashkumar/dicom/pkg/tag"
)

//Notes

func main() {
	startTime := time.Now()
	// folderPath := "/Users/harrymbp/Developer/Projects/PreXion/temp/"
	// outputPath := "/Volumes/Harrypc/temp/Anonymized scans"
	// enableLogging := true
	mainGuiWindow()
	// AnonymizeAllScans(folderPath, outputPath, enableLogging)
	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)
	fmt.Printf("Total time Elapsed time: %.2f seconds\n", elapsedTime.Seconds())
}

func mainGuiWindow() {
	app := app.New()
	mainWindow := app.NewWindow("PreXion Internal Tools V.0.0.1")

	testLabel := widget.NewLabel("Testing label")
	testLabel2 := widget.NewLabel("Testing label2")
	testLabel3 := widget.NewLabel("Testing label3")
	testLabel4 := widget.NewLabel("Testing label4")
	anonymizeLabel := widget.NewLabel("Anonymize")
	toolsLabel := widget.NewLabel("Tools")
	scriptsLabel := widget.NewLabel("Scripts")
	var leftborder *fyne.Container
	//not working as intended
	firstButton := widget.NewButton("Open new window", func() {
		// secondWindow := app.NewWindow("Second")
		// secondWindow.SetContent(widget.NewLabel("second window label"))
		// secondWindow.Show()
		leftborder = container.New(layout.NewVBoxLayout(), testLabel4, scriptsLabel)
		log.Println("button clicked...")

	})
	topBorder := container.New(layout.NewGridLayout(3), anonymizeLabel, toolsLabel, firstButton)
	leftborder = container.New(layout.NewVBoxLayout(), testLabel2, testLabel3)

	// border := container.New(layout.NewBorderLayout(anonymizeLabel, anonymizeLabel, toolsLabel, scriptsLabel))
	content := container.NewBorder(topBorder, nil, leftborder, nil, testLabel)
	mainWindow.SetMaster()
	mainWindow.SetContent(content)
	mainWindow.ShowAndRun()
}

// UID instances
// Syntax for a UID is RootUID + Unit UID + Timestamp 17 digits long + SUBUID
// example 1.2.392.200036.9163.31.0938.20231019103353867.1.1
// ROOTUID = 1.2.392.200036.9163 | Unit UID = 31.0938 | Timestamp = 20231019103353867 || SUBUID = 1.1
func AnonymizeAllScans(inputFolderPath, outputFolderPath string, enableLogging bool) {
	fmt.Printf("------- Start of AnonymizeAllScans Script ---------\n")
	// takes in the inputFolderPath and returns a map with the folderpath:scanDetails | scanDetails are map[scaninfo like fov|name|manufacture]details like the path or fov size and name of the patient
	dicomFolders, err := anonymize_scans.GetDicomFolders(inputFolderPath, enableLogging)
	if err != nil {
		log.Fatal(err)
	}
	//data to modify the dicom info with. takes in tag type for key and value string. | https://pkg.go.dev/github.com/suyashkumar/dicom@v1.0.7/pkg/tag#pkg-constants
	// newDicomAttribute := map[tag.Tag]string{
	// 	tag.PatientName: "new testing function",
	// 	tag.PatientID:   "1700079492",
	// }
	//create a rootUID and UnitUID = 1.2.392.200036.9163.99.9999
	totalscans := len(dicomFolders)
	fmt.Printf("found %d scans\n", totalscans)
	//create a log file to keep track of modified scans
	logFileName := filepath.Join(outputFolderPath, "ModifiedScans.txt")
	logFile, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error occured creating a log file for the modified scans.")
	}
	defer logFile.Close()
	logger := log.New(logFile, "", log.LstdFlags)
	counter := 0
	for parentPath, scanDetails := range dicomFolders {
		fmt.Printf("working on scan: %s\n", parentPath)
		//generate a map that associates each scan with a new output path
		fmt.Printf("Generating output paths\n")
		folderInfo, err := anonymize_scans.MakeOutputPath(parentPath, outputFolderPath, counter, scanDetails, enableLogging)
		if err != nil {
			log.Fatal(err)
		}
		//grab the parent folder and assign it to newScanParentFolder
		var newScanParentFolder string
		for _, parentFolder := range folderInfo {
			newScanParentFolder = filepath.Dir(parentFolder)
			break
		}
		fmt.Printf("outputPaths are \n%s\n", folderInfo)
		fmt.Printf("Generating new information for the dicom files\n")
		newPatientInfo, err := anonymize_scans.RandomizePatientInfo(scanDetails, enableLogging)
		if err != nil {
			log.Fatal(err)
		}
		//set the tag info for StudyInstanceuid to a value we will modify this when working with the scans.
		newPatientInfo[tag.StudyInstanceUID] = "1.2.392.200036.9163.99.9999"
		fmt.Printf("Finished generating new dicom info\n")
		logResults, _ := anonymize_scans.LogAnonymizedScan(scanDetails, newPatientInfo, enableLogging)
		fmt.Printf("Generating dicom files with the new patientInfo\n")
		err = anonymize_scans.MakeStudyFolder(folderInfo, newPatientInfo, enableLogging)
		if err != nil {
			log.Fatal(err)
		}
		counter++
		fmt.Printf("finished anonymizing Dicoms for %s the new scan is located at %s\nORGINIALPATIENTID is:%s, NEWPATIENTID is:%s\n\n", logResults["LOCATION"], newScanParentFolder, logResults["ORGINIALPATIENTID"], logResults["NEWPATIENTID"])
		logger.Printf("finished anonymizing Dicoms for %s the new scan is located at %s\nORGINIALPATIENTID is:%s, NEWPATIENTID is:%s\n\n", logResults["LOCATION"], newScanParentFolder, logResults["ORGINIALPATIENTID"], logResults["NEWPATIENTID"])
		fmt.Printf("finished working on scan %d out of %d\n\n", counter, totalscans)
	}
	fmt.Printf("------- End of AnonymizeAllScans Script ---------\n\n")
}
