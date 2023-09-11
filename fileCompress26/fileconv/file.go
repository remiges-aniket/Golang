package fileconv

import (
	"archive/zip"
	"fmt"
	"log"
	"os"

	"redmine.merce.co/git/go-learning-project/anikets/constants"
)

// File2Zip : This function will add the given file to zip on given path where
// 'path' is the actual file path wothout file name, 'fileToAdd' is the actual file name,
// 'zipFileName' is the output zip file name & return false on failed operation
func File2Zip(path string, fileToAdd string, zipFileName string) bool {
	// Read the given file and return data in []byte
	contents, err := os.ReadFile(path + fileToAdd)
	if err != nil {
		log.Fatal(constants.ERROR_26_2, err)
		return false
	}
	// Create the zip file
	zipReader, err := os.Create(path + zipFileName)
	if err != nil {
		log.Fatal(constants.ERROR_26_3, err)
		return false
	}
	// Create a new gzip writer
	gzWriter := zip.NewWriter(zipReader)

	reader, err := gzWriter.Create(fileToAdd)
	if err != nil {
		log.Fatal(constants.ERROR_26_4, err)
		return false
	}
	// Writing bytes into zip file
	reader.Write(contents)
	gzWriter.Close()
	return true
}

// FileSize : this will return the file size if exists else -1
func FileSize(path string) int64 {
	// Get the size of the compressed file.
	compressedFileSize, err := os.Stat(path)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	return compressedFileSize.Size()
}