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
	//create a new application
	app := app.New()
	mainWindow := app.NewWindow("PreXion Internal Tools V.0.0.1") //create a mainwindow for the application
	mainContainer := container.NewWithoutLayout()
	//login objects
	username := widget.NewEntry()
	username.SetPlaceHolder("Username")

	password := widget.NewEntry()
	password.Password = true
	password.SetPlaceHolder("password")
	// message := widget.NewLabel("Wrong Password")
	// message.Hide() //hide the message by default if password is correcet then show it
	loginForm := container.New(layout.NewFormLayout(), widget.NewLabel("login"), username, widget.NewLabel("password"), password)

	//create views and their tabs
	anonymizeView := anonymizeContent()
	scriptsView := scriptContent()
	toolsView := toolsContent()
	//create main tab
	mainTabsContainer := container.NewAppTabs(
		container.NewTabItem("Anonymize", anonymizeView),
		container.NewTabItem("Scripts", scriptsView),
		container.NewTabItem("Tools", toolsView),
	)

	// testLabel := widget.NewLabel("TestLabel") //test label widget
	loginButton := widget.NewButton("Login", func() {
		loginFunction(username.Text, password.Text, mainWindow, mainContainer, mainTabsContainer)
	})
	mainWindow.SetMaster()
	mainWindow.Resize(fyne.NewSize(400, 200))
	mainContainer = container.New(layout.NewVBoxLayout(), loginForm, loginButton)
	mainWindow.SetContent(mainContainer)
	mainWindow.Canvas().Focus(username)
	username.OnSubmitted = func(p string) {
		loginFunction(username.Text, p, mainWindow, mainContainer, mainTabsContainer)
	}
	password.OnSubmitted = func(p string) {
		loginFunction(username.Text, p, mainWindow, mainContainer, mainTabsContainer)
	}
	mainWindow.ShowAndRun()
}

// takes in username and password as well as the window and tabs we want, if login is correct change the mainwindows content to the default page.
func loginFunction(username, password string, mainWindow fyne.Window, mainContainer *fyne.Container, mainTabsContainer *container.AppTabs) {
	length := len(mainContainer.Objects)
	log.Printf("username is:%s password is:%s length is %d", username, password, length)
	if username == "admin" && password == "admin" {
		mainWindow.SetContent(mainTabsContainer)
	} else if length <= 2 {
		mainContainer.Add(widget.NewLabel("Wrong Password"))
	}
}

// generates the Content or canvas for the the anonymize view.
func anonymizeContent() *container.Split {
	var anonymizeView *container.Split
	content := container.New(layout.NewVBoxLayout(), widget.NewLabel("Input Path"), widget.NewEntry(), widget.NewLabel("new button should be here"))
	anonymizeTabsContainer := container.NewVBox(widget.NewButton("Scan info", func() {
		anonymizeScansView(0, content)
	}), widget.NewButton("Anonymize Scans", func() {
		anonymizeScansView(1, content)
	}))
	anonymizeView = container.NewHSplit(anonymizeTabsContainer, content)
	anonymizeView.Offset = 0.2 //offsets the split view left side is smaller.
	return anonymizeView
}
func anonymizeScansView(tab int, content *fyne.Container) {
	switch tab {
	case 1: //second tab option to anonymize scans using default settings.
		content.RemoveAll()
		inputPath := widget.NewEntry()
		anonymizeButton := widget.NewButton("Anonymize!", func() {
			log.Println("inputpath is:", inputPath.Text)
		})
		content.Add(container.New(layout.NewVBoxLayout(), widget.NewLabel("Input Path"), inputPath, anonymizeButton))
	default: //first tab option grabs the scan info from the input path provided.
		content.RemoveAll()
		content.Add(container.New(layout.NewFormLayout(), widget.NewLabel("Input Path"), widget.NewEntry(), widget.NewLabel("Outputpath"), widget.NewEntry()))
	}

}

// generates the Content or canvas for the the script view.
func scriptContent() *container.Split {
	scriptTabs := container.NewVBox(widget.NewButton("script tab 1", nil), widget.NewButton("script tab 2", nil))
	scriptsView := container.NewHSplit(scriptTabs, widget.NewLabel("Script view!"))
	scriptsView.Offset = 0.2
	return scriptsView
}

// generates the Content or canvas for the the tools view.
func toolsContent() *container.Split {
	toolsTabs := container.NewVBox(widget.NewButton("tools tab 1", nil), widget.NewButton("tool tabs 2", nil))
	toolsView := container.NewHSplit(toolsTabs, widget.NewLabel("Tools view!"))
	toolsView.Offset = 0.2 //offsets the split view left side is smaller.
	return toolsView
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
