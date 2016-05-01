package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	defaultDir := "/tmp/{repository_name}"

	// Parse command line arguments
	namePtr := flag.String("name", "", "Part of the user's name (case sensitive).")
	dirPtr := flag.String("dir", defaultDir, "Directory to clone repository in.")
	keepPtr := flag.Bool("keep", false, "Set to true to keep the local repository after script ends.")

	flag.Parse()

	// Parse required positional argument
	if len(flag.Args()) == 0 {
		flag.Usage()
		os.Exit(1)
	}
	repositoryUrl := flag.Args()[0]

	// Construct directory name
	urlSplits := strings.Split(repositoryUrl, "/")
	repositoryName := urlSplits[len(urlSplits)-1]

	if *dirPtr == defaultDir {
		*dirPtr = fmt.Sprintf("/tmp/%s", repositoryName)
	}

	gitCommand := "git"

	// If the directory does not yet exist, create it
	_, err := os.Stat(*dirPtr)
	if err != nil {
		err = os.Mkdir(*dirPtr, 0700)

		if err != nil {
			log.Fatal(err)
		}

		// Clone the repository
		gitCloneArgs := []string{"clone", repositoryUrl, *dirPtr}

		cmd := exec.Command(gitCommand, gitCloneArgs...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
	}

	// Remove directory and contents on exit
	if *keepPtr == false {
		defer os.RemoveAll(*dirPtr)
	}

	// Extract names and emails from git log
	gitLogArgs := []string{"-C", *dirPtr, "log", "--all", "--pretty=format:%an <%ae>\n"}

	// Add optional name flag
	if *namePtr != "" {
		gitLogArgs = append(gitLogArgs, fmt.Sprintf("--author=%s", *namePtr))
	}

	// Sort the results to be unique
	sortCmd := exec.Command("sort", "-u")
	gitlogCmd := exec.Command(gitCommand, gitLogArgs...)

	// Construct a pipe from git log to sort
	sortCmd.Stdout = os.Stdout
	sortCmd.Stdin, err = gitlogCmd.StdoutPipe()

	// Execute the commands
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
