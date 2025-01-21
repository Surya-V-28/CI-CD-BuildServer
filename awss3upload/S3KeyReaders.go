package awss3upload

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readAWSCredentials(filepath string) (string, string, string, error) {
	// openning the AWS Credentials Files
	file, err := os.Open(filepath)
	if err != nil {
		return "", "", "", fmt.Errorf("error in opening the file %v", err)
	}

	defer file.Close() // closing the file after reading the contents in it

	// intiatize the variables to pass in the S3
	var accessKeyID, secretAccessKey, region string
	scanner := bufio.NewScanner(file)

	// reading the line by line

	for scanner.Scan() {

		line := strings.TrimSpace(scanner.Text())
		// check in the each line if the sepecifc elements exists or not
		if strings.HasPrefix(line, "access_key_id=") {
			accessKeyID = strings.TrimPrefix(line, "access_key_id=")
		} else if strings.HasPrefix(line, "secret_access_key=") {
			secretAccessKey = strings.TrimPrefix(line, "secret_access_key=")
		} else if strings.HasPrefix(line, "region=") {
			region = strings.TrimPrefix(line, "region=")
		}

	}

	return "", "", "", nil

}
