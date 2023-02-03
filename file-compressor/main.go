package main

import (
	"archive/zip"
	"io"
	"log"
	"os"
)

func main() {
	textFile, err := os.Open("./sample/A1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer textFile.Close()
	// Create the zip file
	zipFile, err := os.Create("./text_file.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer zipFile.Close()
	// Create a new zip writer
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()
	// Add the text file to the zip
	zipTextFile, err := zipWriter.Create("A1.txt")
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.Copy(zipTextFile, textFile)
	if err != nil {
		log.Fatal(err)
	}

}
