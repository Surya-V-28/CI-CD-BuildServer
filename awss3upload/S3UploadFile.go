package awss3upload

import (

	"archive/zip"
	"bufio"
	"fmt"
	"os"
	"is"
	"bytes"
	"path/filepath"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	

)

func zipFolder(folderPath string) ([] byte, error) {

	var buf bytes.Buffer
	zipWriter = zip.NewWriter(&buf)

	err:=filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err!=nil {
			return err
		}




		// it the info of the file is directory 
		if info.IsDir() {
			return nil;
		}

		// its basically takes one by one file and add it to the zip file 


		// getting the relative path of the file for the zip floder 
		rePath, err := filepath.Rel(folderPath,path)

		if err !=nil {
			return  fmt.Errorf("failed to caluclate the relative path %v", err)
		}

		//  create a new entry in the zip for the current file

		zipFileWriter ,err:=   zipWriter.Create(relPath)

		if err != nil {
			return fmt.Errorf(" falied to make the erntry for the current file in the  zip file  %v", err)
		}
   
		


	})






	return [], err
}

func starterS3upload() {

	srcFolderPath:=""
	s3BucketName:="appstoragebuild"
	zipFilePath:=""
	applicationNameinS3:=""

   zippedData , err := zipFolder(srcFolderPath)
   if err !=nil {
	  fmt.Printf("error in creating the zipfolder %v", err)
	  return 
   }




}
