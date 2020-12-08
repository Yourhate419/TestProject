package main

import (
	"os"
)

func setenv() {
	os.Setenv("FILENAMECSV", "data.csv")
	os.Setenv("FILENAMEJSON", "output.json")
}