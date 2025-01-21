package awss3upload

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func zipFolder(folderPath string) ([]byte, error) {

	var buf bytes.Buffer
	zipWriter := zip.NewWriter(&buf)

	// walking the each files and copying it in the zip floder

	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// it the info of the file is directory
		if info.IsDir() {
			return nil
		}

		// its basically takes one by one file and add it to the zip file

		// getting the relative path of the file for the zip floder
		relPath, err := filepath.Rel(folderPath, path)

		if err != nil {
			return fmt.Errorf("failed to caluclate the relative path %v", err)
		}

		//  create a new entry in the zip for the current file

		zipFileWriter, err := zipWriter.Create(relPath)

		if err != nil {
			return fmt.Errorf(" falied to make the erntry for the current file in the  zip file  %v", err)
		}

		//open the file to be added in the Zip
		file, err := os.Open(path)
		if err != nil {
			return fmt.Errorf("Failed to open the file %v", err)
		}

		defer file.Close() // ensure the file is closed aftering the contents in it

		// copy the content of the file to the zip floder
		_, err = io.Copy(zipFileWriter, file)
		if err != nil {
			return fmt.Errorf("Falied to copy the contents in the zip floder %v", err)
		}

		return nil
	})

	// closing the zip opened file after adding the content in it

	err = zipWriter.Close()
	if err != nil {
		return nil, fmt.Errorf("Falied to close the Zip Floder after adding the content %v", err)
	}

	return buf.Bytes(), nil
}
