package maiinsn

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/go-git/go-git/v5"
)

func sfunctisons() {

	// its in other build server voice breaking too

	var cloneDir string = "D:/tmp/CiCD"
	var cloneGitUrl string = "https://github.com/YaminNather/ci-cd-test-subject.git"

	// Clone the given repository to the given directory

	fmt.Println("git clone   " + cloneGitUrl + "at lociation " + cloneDir)
	_, err := git.PlainClone("/tmp/CiCD", false, &git.CloneOptions{
		URL:      cloneGitUrl,
		Progress: os.Stdout,
	})
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	// changing the working directory and working
	fmt.Printf("working on the npm install")
	cmd := exec.Command("npm", "install")
	cmd.Dir = cloneDir // set the working directory
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error : ", err)
		os.Exit(1)
	}
	// Setting the NPM run dev
	fmt.Printf("Runing the Command on the npm run build")
	runCmd := exec.Command("npm", "run", "build")
	runCmd.Dir = cloneDir
	runCmd.Stderr = os.Stderr
	runCmd.Stdout = os.Stdout
	err = runCmd.Run()
	if err != nil {
		fmt.Println("Error running 'npm run dev':", err)
		os.Exit(1)
	}
	// The development server should now be running
	fmt.Println("Next.js app is running!")

}
