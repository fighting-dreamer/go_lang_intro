package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
)

func addToZip(fileName string, zw *zip.Writer) error {
	//Open the file to archive to zip
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	//Create adds a field to zip file using given name
	//Create return io.Writer to which file content should be written
	wr, err := zw.Create(fileName)
	if err != nil {
		return err
	}
	//Write content to zip file
	if _, err := io.Copy(wr, file); err != nil {
		return err
	}
	return nil
}

func archiveFiles(files []string, archive string) error {
	flags := os.O_CREATE | os.O_WRONLY | os.O_TRUNC
	//Open the zip file
	file, err := os.OpenFile(archive, flags, 0644)
	if err != nil {
		return err
	}
	defer file.Close() // closing the zip file
	//Create a zip file writer
	zw := zip.NewWriter(file)
	defer zw.Close()
	//Iterate through files and zip them
	for _, filename := range files {
		// Write file to zip file
		if err := addToZip(filename, zw); err != nil {
			return err
		}
	}
	return nil
}

func readFromArchive(archive string) error {
	// Open the zip file specified by the given name and return a REadCloser
	rc, err := zip.OpenReader(archive)
	if err != nil {
		return err
	}
	defer rc.Close()
	//Iterate through each file
	for _, file := range rc.File {
		frc, err := file.Open()
		if err != nil {
			return err
		}
		defer frc.Close()
		fmt.Fprintf(os.Stdout, "Contents of File : %s\n", file.Name)
		// write teh content to Stdout
		copied, err := io.Copy(os.Stdout, frc)
		if err != nil {
			return err
		}
		if uint64(copied) != file.UncompressedSize64 {
			return fmt.Errorf("Length of the File Content doesn`t match with the file %s", file.Name)
		}
		fmt.Println() // to add a empty line
	}
	return nil
}

func main() {
	//Name a zip File
	archive := "source.zip"
	filesToBeArchived := []string{"main.txt", "readme.txt"}
	//Archive files
	err := archiveFiles(filesToBeArchived, archive)
	if err != nil {
		log.Fatalf("Eror While Writing to Zip file : %s\n", err)
	}
	//Read from Zip File
	err = readFromArchive(archive)
	if err != nil {
		log.Fatalf("Error While reading from Zip file : %s\n", err.Error())
	}
	// Side note, to uncompress via command line, u may need to change the (suffix) file  extension from "zip" to "gz" as is source.gz from source.zip
}
