package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func main() {

	// Parse command line arguments
	namePtr := flag.String("name", "", "Part of the user's name (case sensitive).")

	flag.Parse()

	// Parse required positional argument
	if len(flag.Args()) == 0 {
		flag.Usage()
		os.Exit(1)
	}
	repositoryUrl := flag.Args()[0]

	// Create a temporary directory to clone the repository in
	tempDirPath, err := ioutil.TempDir("", "repository")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Created temporary directory: %s\n", tempDirPath)
	defer os.RemoveAll(tempDirPath)

	// Clone the repository
	gitCommand := "git"
	gitCloneArgs := []string{"clone", repositoryUrl, tempDirPath}

	cmd := exec.Command(gitCommand, gitCloneArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	// Extract emails from git log
	gitLogArgs := []string{"-C", tempDirPath, "log", "--all", "--pretty=format:%an <%ae>\n"}

	if *namePtr != "" {
		gitLogArgs = append(gitLogArgs, fmt.Sprintf("--author=%s", *namePtr))
	}

	gitlogCmd := exec.Command(gitCommand, gitLogArgs...)
	sortCmd := exec.Command("sort", "-u")

	sortCmd.Stdout = os.Stdout
	sortCmd.Stdin, err = gitlogCmd.StdoutPipe()

	if err != nil {
		log.Fatal(err)
	}
	err = sortCmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	err = gitlogCmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	err = sortCmd.Wait()
	if err != nil {
		log.Fatal(err)
	}
}
