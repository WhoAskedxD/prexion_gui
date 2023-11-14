module github.com/WhoAskedxD/prexion_gui

go 1.20

replace github.com/WhoAskedxD/anonymize_scans => ../anonymize_scans

require (
	github.com/WhoAskedxD/anonymize_scans v0.0.0-20230907215908-69fc3c9d73e3
	github.com/suyashkumar/dicom v1.0.6
)

require golang.org/x/text v0.13.0 // indirect
